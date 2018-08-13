package models

import (
	"fmt"
)

type UserName struct {
	Id      int64
	Uid     int64
	Name    string
	Created int64 `xorm:"created"`
	Updated int64 `xorm:"updated"`
}

func CreateUserNameWhenNoExist(uid int64, name string) (*UserName, error) {
	if len(name) == 0 {
		return nil, fmt.Errorf("username len is 0")
	}

	un, err := GetUserNameByName(uid, name)
	if un != nil {
		return un, nil
	}
	un2 := UserName{Uid: uid, Name: name}
	err = AddUserName(&un2)
	if err != nil {
		return nil, err
	}
	return &un2, nil
}

// add
func AddUserName(u *UserName) (err error) {
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

// delete

// update
func EditUserName(u *UserName) (err error) {
	if _, err = x.Id(u.Id).Update(u); err != nil {
		return err
	}

	return nil
}

// get
func GetUserNameById(id int64) (*UserName, error) {
	u := new(UserName)
	has, err := x.Id(id).Get(u)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, fmt.Errorf("username does not exist [id: %d, name: %s]", id, "")
	}
	return u, nil
}

func GetUserNameByName(uid int64, name string) (*UserName, error) {
	if len(name) == 0 {
		return nil, fmt.Errorf("username does not exist [id: %d, name: %s]", 0, name)
	}
	u := &UserName{Uid: uid, Name: name}
	has, err := x.Get(u)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, fmt.Errorf("username does not exist [id: %d, name: %s]", 0, name)
	}
	return u, nil
}
