package database

import (
	"fmt"
	"github.com/go-xorm/xorm"
	"time"
)

type VersionState int

const (
	VERSION_STATE_IDLE VersionState = iota // Historic reason to make it starts at 0.
	VERSION_STATE_TESTING
	VERSION_STATE_ONLINE
)

type InfoType int

const (
	INFO_TYPE_ADDED InfoType = iota
)

// 项目
type Project struct {
	Id      int64
	Name    string `xorm:"UNIQUE NOT NULL"`
	Desc    string
	Content string `xorm:"-" json:"-"`
	Created int64  `xorm:"created"`
	Updated int64  `xorm:"updated"`
}

// 项目内容
type ProjectContent struct {
	Id      int64
	Pid     int64
	Content string
}

// project list
func GetProjects(page, pageSize int) ([]*Project, error) {
	projs := make([]*Project, 0, pageSize)
	return projs, x.Limit(pageSize, (page-1)*pageSize).Asc("id").Find(&projs)
}

// project and content add alway
func CreateProject(p *Project) (err error) {
	sess := x.NewSession()
	defer sess.Close()
	if err = sess.Begin(); err != nil {
		return err
	}

	if _, err = sess.Insert(p); err != nil {
		return err
	}
	c := &ProjectContent{
		Pid:     p.Id,
		Content: "",
	}
	if _, err = sess.Insert(c); err != nil {
		return err
	}

	return sess.Commit()
}

// GetProjectByID returns the project object by given ID if exists.
func GetProjectByID(id int64) (*Project, error) {
	return getProjectByID(x, id)
}

// project content q
func (p *Project) GetProjectContent() (err error) {
	p.Content, err = GetProjectContentByID(p.Id)
	return err
}

// project content edit
func (p *Project) EditProjectContent(content string) (err error) {
	return EditProjectContentByID(p.Id, content)
}

func getProjectByID(e *xorm.Engine, id int64) (*Project, error) {
	p := new(Project)
	has, err := e.Id(id).Get(p)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, fmt.Errorf("project does not exist [project_id: %d, name: %s]", id, "")
	}
	return p, nil
}

func GetProjectContentByID(id int64) (string, error) {
	return getProjectContentByID(x, id)
}

func getProjectContentByID(e *xorm.Engine, projectID int64) (string, error) {
	p := new(ProjectContent)
	has, err := e.Where("pid=?", projectID).Get(p)
	if err != nil {
		return "", err
	} else if !has {
		return "", fmt.Errorf("project does not exist [project_id: %d, name: %s]", projectID, "")
	}
	return p.Content, nil
}

func EditProjectContentByID(id int64, content string) error {
	return editProjectContentByID(id, content)
}

func editProjectContentByID(projectID int64, content string) (err error) {
	sess := x.NewSession()
	defer sess.Close()
	if err = sess.Begin(); err != nil {
		return err
	}
	if _, err = sess.Exec("UPDATE `project_content` SET content = ? WHERE pid = ?", content, projectID); err != nil {
		return err
	}

	if _, err = sess.Exec("UPDATE `project` SET updated = ? WHERE id = ?", time.Now().Unix(), projectID); err != nil {
		return err
	}

	return sess.Commit()
}
