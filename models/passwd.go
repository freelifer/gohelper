package models

import (
	"github.com/go-xorm/xorm"
)

/* 密码信息 */
type PasswdInfo struct {
	Id       int64
	Uid      int64
	IconId   int64
	Title    string
	Username string
	Passwd   string
	Created  int64 `xorm:"created"`
	Updated  int64 `xorm:"updated"`
}

type PasswdInfoBean struct {
	Id       int64
	Icon     string
	Title    string
	Username string
	Passwd   string
	Created  int64
	Updated  int64
}

func AddPasswdInfo(p *PasswdInfo) (err error) {
	sess := x.NewSession()
	defer sess.Close()
	if err = sess.Begin(); err != nil {
		return err
	}

	if _, err = sess.Insert(p); err != nil {
		return err
	}

	return sess.Commit()
}

func GetUserPasswdsByWxUserId(userID int64) ([]*PasswdInfoBean, error) {
	sess := x.NewSession()
	return getUserPasswdsByWxUserId(sess, userID)
}

func getUserPasswdsByWxUserId(sess *xorm.Session, userID int64) ([]*PasswdInfoBean, error) {
	passwds := make([]*PasswdInfoBean, 0, 10)
	sql := `SELECT passwd_info.id, 
	passwd_info.title, passwd_info.username, 
	passwd_info.passwd, passwd_info.created, 
	passwd_info.updated, icon_info.url as icon FROM passwd_info, icon_info`
	return passwds, sess.SQL(sql).Where("`passwd_info`.uid=?", userID).And("`passwd_info.icon_id=?`", "`icon_info`.id").Find(&passwds)
	// return passwds, sess.Where("uid=?", userID).Find(&passwds)
}
