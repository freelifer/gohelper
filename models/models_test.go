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

		Convey("GetWxUserByOpenId", func() {
			u, err := GetWxUserByOpenId("abcdefg")
			t.Log(u)
			So(u.Id, ShouldEqual, 1)
			So(err, ShouldBeNil)
		})
	})
}
