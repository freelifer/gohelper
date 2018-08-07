package settings

import (
	"testing"
)

func Test_Setup(t *testing.T) {
	Setup()

	if AppCfg.Name != "gohelper" {
		t.Error("settings parse conf/app.conf AppCfg.Name != gohelper")
	}

	if DatabaseCfg.Type != "mysql" {
		t.Error("settings parse conf/app.conf DatabaseCfg.Type != mysql")
	}

	if DatabaseCfg.Host != "127.0.0.1:3306" {
		t.Error("settings parse conf/app.conf DatabaseCfg.Host != 127.0.0.1:3306")
	}

	if DatabaseCfg.Name != "blog" {
		t.Error("settings parse conf/app.conf DatabaseCfg.Name != blog")
	}

	if DatabaseCfg.User != "root" {
		t.Error("settings parse conf/app.conf DatabaseCfg.User != root")
	}

	if DatabaseCfg.Passwd != "rootroot" {
		t.Error("settings parse conf/app.conf DatabaseCfg.Passwd != rootroot")
	}

	if RedisCfg.Host != "127.0.0.1:6379" {
		t.Error("settings parse conf/app.conf RedisCfg.Host != 127.0.0.1:6379")
	}

	if RedisCfg.Password != "" {
		t.Error("settings parse conf/app.conf RedisCfg.Password != ")
	}

	if RedisCfg.MaxIdle != 30 {
		t.Error("settings parse conf/app.conf RedisCfg.MaxIdle != 30")
	}

	if RedisCfg.MaxActive != 30 {
		t.Error("settings parse conf/app.conf RedisCfg.MaxActive != 30")
	}

	if RedisCfg.IdleTimeout != 200 {
		t.Error("settings parse conf/app.conf RedisCfg.IdleTimeout != 200")
	}

	t.Log(AppCfg.Version)
}
