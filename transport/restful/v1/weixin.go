package v1

import (
	"github.com/freelifer/coolgo/app/service"
	"github.com/freelifer/coolgo/models"
	"github.com/freelifer/coolgo/pkg/redis"
	"github.com/freelifer/coolgo/pkg/utils"
	"github.com/gin-gonic/gin"
)

// @Summary 微信登录
// @Tags user
// @Produce json
// @Param body body string true "微信code" default(xxxxxx)
// @Success 200 {string} json "{"code":200,"data":{"session_id":"xxxxxxxxxxx"},"msg":"ok"}"
// @Router /v1/wxlogin [get]
func WeiXinLogin(c *gin.Context) {
	code := c.Query("code")

}
