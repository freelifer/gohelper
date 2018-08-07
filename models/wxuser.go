package models

import (
	"encoding/json"
	"time"
)

/* 微信用户 */
type WxUser struct {
	Id       int64
	WxOpenid string `xorm:"unique"`
	Created  int64  `xorm:"created"`
	Updated  int64  `xorm:"updated"`
}

func CreateWxUserWhenNoExist(openid string) (*User, error) {
	if len(openid) == 0 {
		return nil, fmt.Errorf("wxuser openid len is 0")
	}

	wxUser, err := GetWxUserByOpenId(openid)
	if wxUser != nil {
		return wxUser, nil
	}
	wxUser2 = WxUser{WxOpenid: openid}
	err := AddWxUser(&wxUser2)
	if err != nil {
		return nil, err
	}
	return &wxUser2, nil
}

func ExistWxUserByOpenId(openid string) (bool, error) {
	if len(openid) == 0 {
		return false, nil
	}
	var user WxUser
	_, err := x.Select("id").Where("name = ?", openid).Get(&user)
	if err != nil {
		return false, err
	}

	if user.Id > 0 {
		return true, nil
	}
	return false, nil
}

func GetWxUserByID(id int64) (*User, error) {
	u := new(WxUser)
	has, err := x.Id(id).Get(u)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, fmt.Errorf("wxuser does not exist [wxuser_id: %d, openid: %s]", id, "")
	}
	return u, nil
}

func GetWxUserByOpenId(openid string) (*User, error) {
	if len(openid) == 0 {
		return nil, fmt.Errorf("wxuser does not exist [wxuser_id: %d, openid: %s]", 0, name)
	}
	u := &WxUser{WxOpenid: openid}
	has, err := x.Get(u)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, fmt.Errorf("wxuser does not exist [wxuser_id: %d, name: %s]", 0, name)
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
