package v1

import (
	"github.com/freelifer/gohelper/pkg/e"
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
	if len(code) == 0 {
		e.ErrorJSON(c, e.WX_CODE_EMPTY)
		return
	}
	s := wx_service.WxService{Code: code}
	err := s.Login()

	if err != nil {
		e.ErrorJSON(c, err)
		return
	}
	e.SuccessJSON(c, gin.H{"session_id": s.SessionId})
}
