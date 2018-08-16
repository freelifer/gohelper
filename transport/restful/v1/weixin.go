package v1

import (
	"github.com/freelifer/gohelper/service/wx_service"
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

	s := wx_service.WxService{Code: code}
	err := s.Login()

	if err != nil {
		errorJSON(c, 100, err.Error())
		return
	}
	successJSON(c, gin.H{"session_id": s.SessionId})
}

func successJSON(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"errcode": 0,
		"errmsg":  "",
		"data":    data,
	})
}
func errorJSON(c *gin.Context, errcode int, errmsg string) {
	c.JSON(http.StatusOK, gin.H{
		"errcode": errcode,
		"errmsg":  errmsg,
	})
}
