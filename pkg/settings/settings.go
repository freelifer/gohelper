package settings

import (
	"github.com/go-ini/ini"
	"log"
	"time"
)

var (
	AppCfg struct {
		Name, Version, Port string
	}
	ServerCfg struct {
		RunMode      string
		HttpPort     int
		ReadTimeout  time.Duration
		WriteTimeout time.Duration
	}
	DatabaseCfg struct {
		Type, Host, Name, User, Passwd, Path string
	}

	LogCfg struct {
		Type, Path string
	}
	RedisCfg struct {
		Host        string
		Password    string
		MaxIdle     int
		MaxActive   int
		IdleTimeout time.Duration
	}
	WxCfg struct {
		Appid  string
		Secret string
	}
)

// func init() {
// 	AppCfg.Name = goconfig.MustValue("app", "name", "default")
// 	AppCfg.Version = "1.0"
// 	AppCfg.Port = "8888"

// 	DatabaseCfg.Type = goconfig.MustValue("database", "DB_TYPE", "sqlite3")
// 	DatabaseCfg.Host = goconfig.MustValue("database", "HOST", "")
// 	DatabaseCfg.Name = goconfig.MustValue("database", "NAME", "")
// 	DatabaseCfg.User = goconfig.MustValue("database", "USER", "")
// 	DatabaseCfg.Passwd = goconfig.MustValue("database", "PASSWD", "")
// 	DatabaseCfg.Path = goconfig.MustValue("database", "PATH", "data/doc.db")

// 	// log config
// 	LogCfg.Type = goconfig.MustValue("log", "TYPE", "stdout")
// 	LogCfg.Path = goconfig.MustValue("log", "PATH", "data/access.log")
// }

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.conf")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.conf': %v", err)
	}
	mapTo("app", &AppCfg)
	mapTo("server", &ServerCfg)
	mapTo("database", &DatabaseCfg)
	mapTo("redis", &RedisCfg)
	mapTo("weixin", &WxCfg)
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}
}
