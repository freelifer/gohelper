package models

import (
	"github.com/go-xorm/xorm"
)

/* 密码信息 */
type PasswdInfo struct {
	Id       int64
	Uid      int64
	Username string
	Passwd   string
	Icon     string
	Created  int64 `xorm:"created"`
	Updated  int64 `xorm:"updated"`
}

func GetUserPasswdsByWxUserId(userID int64) ([]*PasswdInfo, error) {
	sess := x.NewSession()
	return getUserPasswdsByWxUserId(sess, userID)
}

func getUserPasswdsByWxUserId(sess *xorm.Session, userID int64) ([]*PasswdInfo, error) {
	passwds := make([]*PasswdInfo, 0, 10)
	return passwds, sess.Where("`passwd_info`.uid=?", userID).Find(&passwds)
}
