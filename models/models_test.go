package models_test

import (
	. "github.com/freelifer/gohelper/models"
	. "github.com/smartystreets/goconvey/convey"
	"testing"

	"github.com/freelifer/gohelper/pkg/settings"
	_ "github.com/mattn/go-sqlite3"
)

func Init_DB() {
	settings.DatabaseCfg.Type = "sqlite3"
	settings.DatabaseCfg.Path = "../data/doc.db"
	Setup()
}

func Test_WxUser(t *testing.T) {
	Init_DB()
	defer DropTables()
	Convey("model WxUser测试", t, func() {
		Convey("AddWxUser", func() {
			wxUser := WxUser{WxOpenid: "abcdefg"}
			err := AddWxUser(&wxUser)
			So(err, ShouldBeNil)
		})
		Convey("GetWxUserByID", func() {
			u, err := GetWxUserByID(1)
			t.Log(u)
			So(u.Id, ShouldEqual, 1)
			So(err, ShouldBeNil)
		})
		Convey("GetWxUserByOpenId", func() {
			u, err := GetWxUserByOpenId("abcdefg")
			t.Log(u)
			So(u.Id, ShouldEqual, 1)
			So(err, ShouldBeNil)
		})
		Convey("GetWxUserByOpenId not", func() {
			u, err := GetWxUserByOpenId("abcdefgaa")
			t.Log(u)
			So(err, ShouldNotBeNil)
		})
		Convey("ExistWxUserByOpenId", func() {
			exist, err := ExistWxUserByOpenId("abcdefg")
			So(exist, ShouldEqual, true)
			So(err, ShouldBeNil)
		})
		Convey("ExistWxUserByOpenId not", func() {
			exist, err := ExistWxUserByOpenId("abcdefg11")
			So(exist, ShouldEqual, false)
			So(err, ShouldBeNil)
		})
		Convey("GetUserPasswds", func() {
			u, err := GetWxUserByOpenId("abcdefg")
			err = u.GetUserPasswds()
			So(err, ShouldBeNil)
			t.Log(u.PasswdInfos)
		})
	})
}
