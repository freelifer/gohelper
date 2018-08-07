package log

import (
	"fmt"
	"github.com/freelifer/gohelper/pkg/settings"
	"github.com/freelifer/gohelper/pkg/utils"
	"os"
	"time"
)

const (
	PROJECT_NAME = "doc"
)

var (
	DefaultOut = os.Stdout
)

func init() {
	if settings.LogCfg.Type == "file" {
		DefaultOut, _ = os.OpenFile(settings.LogCfg.Path, os.O_WRONLY|os.O_CREATE|os.O_SYNC|os.O_APPEND, 0755)
	}
}

func CreateMark() string {
	mark, _ := utils.RandomString(6)
	return mark
}

func Info(mark string, msg string) {
	end := time.Now()
	fmt.Fprintf(DefaultOut, "[API] [%s] %v %s %s\n",
		PROJECT_NAME,
		end.Format("2006/01/02 15:04:05"),
		mark,
		msg,
	)
}
