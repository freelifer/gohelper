package models

type UserName struct {
	Id      int64
	Name    string
	Created int64 `xorm:"created"`
	Updated int64 `xorm:"updated"`
}
