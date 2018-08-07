package server

import (
	"github.com/freelifer/gohelper/server/repository/database"
)

type ProjectServerImpl struct {
	Server
}

func (repo *ProjectServerImpl) List() *ServerResponse {
	projects, err := database.GetProjects(1, 10)
	if err != nil {
		return NewRespForParamError1(repo.Mark, "List "+err.Error())
	}

	return NewRespForSuccess1(H{
		"projects": projects,
	})
}

func (repo *ProjectServerImpl) GetProject(id int64) *ServerResponse {
	p, err := database.GetProjectByID(id)
	if err != nil {
		return NewRespForResourceError(repo.Mark, "GetProject "+err.Error())
	}
	e := p.GetProjectContent()
	if e != nil {
		return NewRespForResourceError(repo.Mark, "GetProject "+err.Error())
	}

	return NewRespForSuccess1(H{
		"content": p.Content,
	})
}

func (repo *ProjectServerImpl) AddProject(name string) *ServerResponse {
	if name == "" {
		return NewRespForParamError1(repo.Mark, "AddProject name is empty")
	}
	p := &database.Project{
		Name: name,
	}
	err := database.CreateProject(p)
	if err != nil {
		return NewRespForResourceError(repo.Mark, "AddProject "+err.Error())
	}
	return NewRespForSuccess1(H{
		"project": p,
	})
}

func (repo *ProjectServerImpl) UpdateProject(id int64, content string) *ServerResponse {
	if id == 0 || content == "" {
		return NewRespForParamError1(repo.Mark, "UpdateProject id or content empty")
	}
	p, err := database.GetProjectByID(id)
	if err != nil {
		return NewRespForResourceError(repo.Mark, "UpdateProject "+err.Error())
	}

	e := p.EditProjectContent(content)
	if e != nil {
		return NewRespForResourceError(repo.Mark, "UpdateProject "+e.Error())
	}

	return NewRespForSuccess1(H{})
}

func (repo *ProjectServerImpl) DelProject(id int) *ServerResponse {

	return NewRespForSuccess("", nil)

}
