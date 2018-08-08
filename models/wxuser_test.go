package models

import (
	"testing"
)

func Test_GetWxUserByOpenId(t *testing.T) {
	init_setup()
	_, err := GetWxUserByOpenId("kzhu")
	if err != nil {
		t.Error(err)
	}
}

func Test_CreateWxUserWhenNoExist(t *testing.T) {
	init_setup()
	_, err := CreateWxUserWhenNoExist("kzhu")
	if err != nil {
		t.Error(err)
	}
}
