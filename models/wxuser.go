package models

import (
	"encoding/json"
	"fmt"
)

/* 微信用户 */
type WxUser struct {
	Id          int64
	WxOpenid    string `xorm:"unique"`
	Passwd      string
	PasswdInfos []*PasswdInfo `xorm:"-" json:"-"`
	Created     int64         `xorm:"created"`
	Updated     int64         `xorm:"updated"`
}

func (u *WxUser) GetUserPasswds() (err error) {
	u.PasswdInfos, err = GetUserPasswdsByWxUserId(u.Id)
	return err
}

func CreateWxUserWhenNoExist(openid string) (*WxUser, error) {
	if len(openid) == 0 {
		return nil, fmt.Errorf("wxuser openid len is 0")
	}

	wxUser, err := GetWxUserByOpenId(openid)
	if wxUser != nil {
		return wxUser, nil
	}
	wxUser2 := WxUser{WxOpenid: openid}
	err = AddWxUser(&wxUser2)
	if err != nil {
		return nil, err
	}
	return &wxUser2, nil
}

func EditWxUserPasswd(id int64, passwd string) (err error) {
	sess := x.NewSession()
	defer sess.Close()
	if err := sess.Begin(); err != nil {
		return err
	}
	if _, err = sess.Exec("UPDATE `wx_user` SET passwd=? WHERE id=?", passwd, id); err != nil {
		return err
	}

	return sess.Commit()
}

func ExistWxUserByOpenId(openid string) (bool, error) {
	if len(openid) == 0 {
		return false, nil
	}
	var user WxUser
	_, err := x.Select("id").Where("wx_openid = ?", openid).Get(&user)
	if err != nil {
		return false, err
	}

	if user.Id > 0 {
		return true, nil
	}
	return false, nil
}

func GetWxUserByID(id int64) (*WxUser, error) {
	u := new(WxUser)
	has, err := x.Id(id).Get(u)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, fmt.Errorf("wxuser does not exist [wxuser_id: %d, openid: %s]", id, "")
	}
	return u, nil
}

func GetWxUserByOpenId(openid string) (*WxUser, error) {
	if len(openid) == 0 {
		return nil, fmt.Errorf("wxuser does not exist [wxuser_id: %d, openid: %s]", 0, openid)
	}
	u := &WxUser{WxOpenid: openid}
	has, err := x.Get(u)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, fmt.Errorf("wxuser does not exist [wxuser_id: %d, openid: %s]", 0, openid)
	}
	return u, nil
}

func AddWxUser(u *WxUser) (err error) {
	sess := x.NewSession()
	defer sess.Close()
	if err = sess.Begin(); err != nil {
		return err
	}

	if _, err = sess.Insert(u); err != nil {
		return err
	}

	return sess.Commit()
}

func WxUserToJson(s *WxUser) (string, error) {
	b, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(b), nil
}

func JsonToWxUser(s string) (*WxUser, error) {
	var wxUser WxUser
	err := json.Unmarshal([]byte(s), &wxUser)
	if err != nil {
		return nil, err
	}
	return &wxUser, nil
}
