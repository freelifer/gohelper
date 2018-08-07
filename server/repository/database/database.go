package database

import (
	"fmt"
	"github.com/freelifer/gohelper/pkg/settings"
	"github.com/go-xorm/xorm"
	"os"
	"path"
)

var (
	x         *xorm.Engine
	tables    []interface{}
	HasEngine bool

	EnableSQLite3 bool
)

func init() {
	tables = append(tables, new(User), new(Project), new(ProjectContent))
}

func NewEngine() (err error) {
	x, err = getEngine()
	if err != nil {
		return fmt.Errorf("Fail to connect to database: %v", err)
	}

	if err = x.StoreEngine("InnoDB").Sync2(tables...); err != nil {
		return fmt.Errorf("sync database struct error: %v\n", err)
	}

	return nil
}

func getEngine() (*xorm.Engine, error) {
	connStr := ""
	// var Param string = "?"
	// if strings.Contains(settings.DatabaseCfg.HostName, Param) {
	// 	Param = "&"
	// }
	switch settings.DatabaseCfg.Type {
	case "mysql":
		if settings.DatabaseCfg.Host[0] == '/' {
			connStr = fmt.Sprintf("%s:%s@unix(%s)/%s?charset=utf8",
				settings.DatabaseCfg.User, settings.DatabaseCfg.Passwd, settings.DatabaseCfg.Host, settings.DatabaseCfg.Name)
		} else {
			connStr = fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8",
				settings.DatabaseCfg.User, settings.DatabaseCfg.Passwd, settings.DatabaseCfg.Host, settings.DatabaseCfg.Name)
		}
		// var engineParams = map[string]string{"rowFormat": "DYNAMIC"}
		// return xorm.NewEngineWithParams(settings.DatabaseCfg.Type, connStr, engineParams)
	case "sqlite3":
		// if !EnableSQLite3 {
		// 	return nil, errors.New("This binary version does not build support for SQLite3.")
		// }
		if err := os.MkdirAll(path.Dir(settings.DatabaseCfg.Path), os.ModePerm); err != nil {
			return nil, fmt.Errorf("Fail to create directories: %v", err)
		}
		connStr = "file:" + settings.DatabaseCfg.Path + "?cache=shared&mode=rwc"
	default:
		return nil, fmt.Errorf("Unknown database type: %s", settings.DatabaseCfg.Type)
	}
	fmt.Println(connStr)
	return xorm.NewEngine(settings.DatabaseCfg.Type, connStr)
}
