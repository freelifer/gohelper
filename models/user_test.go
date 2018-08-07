package models

import (
	"testing"

	"github.com/freelifer/gohelper/pkg/settings"
	_ "github.com/mattn/go-sqlite3"
)

func init_setup() {
	settings.DatabaseCfg.Type = "sqlite3"
	settings.DatabaseCfg.Path = "../data/doc.db"
	Setup()
}

// go test models/*.go -v -o .
func Test_Setup(t *testing.T) {
	init_setup()
}

func Test_ExistUserByName_NO(t *testing.T) {
	init_setup()
	exist, err := ExistUserByName("kzhu")
	if exist {
		t.Error("has kzhu user")
	}

	if err != nil {
		t.Error(err)
	}
}

func Test_AddUser(t *testing.T) {
	init_setup()

	u := &User{Name: "kzhu", Passwd: "123456"}
	err := AddUser(u)

	if err != nil {
		t.Error(err)
	}
	t.Log(u)
}

func Test_ExistUserByName(t *testing.T) {
	init_setup()
	exist, err := ExistUserByName("kzhu")
	if !exist {
		t.Error("not has kzhu user")
	}

	if err != nil {
		t.Error(err)
	}
}

func Test_GetUserByID(t *testing.T) {
	init_setup()
	u, err := GetUserByID(1)

	if err != nil {
		t.Error(err)
	}

	if u.Name != "kzhu" {
		t.Error("Error name")
	}
}
func Test_GetUserByName(t *testing.T) {
	init_setup()
	u, err := GetUserByName("kzhu")

	if err != nil {
		t.Error(err)
	}

	if u.Id != 1 {
		t.Error("Error id")
	}
}

func Test_ValidatePassword(t *testing.T) {
	init_setup()
	u, err := GetUserByName("kzhu")

	if err != nil {
		t.Error(err)
	}

	success := u.ValidatePassword("123456")

	if !success {
		t.Error("Error password")
	}

	success = u.ValidatePassword("1234567")

	if success {
		t.Error("Error password")
	}
}
