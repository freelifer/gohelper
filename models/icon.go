package models

import (
// "github.com/go-xorm/xorm"
)

/* Icon信息 */
type IconInfo struct {
	Id      int64
	Name    string `xorm:"unique"`
	Url     string
	Letter  string
	Created int64 `xorm:"created"`
	Updated int64 `xorm:"updated"`
}

func GetIconInfos() ([]*IconInfo, error) {
	icons := make([]*IconInfo, 0, 10)
	return icons, x.Find(&icons)
}

func AddIconInfo(i *IconInfo) (err error) {
	sess := x.NewSession()
	defer sess.Close()
	if err = sess.Begin(); err != nil {
		return err
	}

	if _, err = sess.Insert(i); err != nil {
		return err
	}

	return sess.Commit()
}

func EditIconInfo(icon *IconInfo) (err error) {
	if _, err = x.Id(icon.Id).Update(icon); err != nil {
		return err
	}

	return nil
}
