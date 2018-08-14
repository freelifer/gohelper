package models

import (
	"github.com/go-xorm/xorm"
)

/* 密码信息 */
type PasswdInfo struct {
	Id       int64
	Uid      int64
	IconId   int64
	Icon     string `xorm:"-"`
	Title    string
	Username string
	Passwd   string
	Created  int64 `xorm:"created"`
	Updated  int64 `xorm:"updated"`
}

// 创建新密码信息
// 根据Unid 判断是否是新的用户名
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

func GetUserPasswdsByWxUserId(userID int64) ([]*PasswdInfo, error) {
	sess := x.NewSession()
	return getUserPasswdsByWxUserId(sess, userID)
}

func getUserPasswdsByWxUserId(sess *xorm.Session, userID int64) ([]*PasswdInfo, error) {
	passwds := make([]*PasswdInfo, 0, 10)
	sql := `SELECT p.id,p.uid,p.title,p.username,p.passwd,p.created,p.updated,i.url as icon
	FROM passwd_info as p
	LEFT JOIN icon_info as i
	ON p.icon_id=i.id
	WHERE p.uid=?`

	return passwds, sess.SQL(sql, userID).Find(&passwds)
}
