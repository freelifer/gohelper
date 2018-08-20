package v1

import (
	"github.com/freelifer/gohelper/pkg/e"
	"github.com/gin-gonic/gin"
)

// @Summary 密码列表
// @Tags passwd
// @Produce json
// @Success 200 {string} json "{"code":200,"data":{"session_id":"xxxxxxxxxxx"},"msg":"ok"}"
// @Router /v1/passwds [get]
func PasswdList(c *gin.Context) {
	e.SuccessJSON(c, gin.H{"PasswdList": "PasswdList"})
}

// @Summary 密码详情
// @Tags passwd
// @Produce json
// @Success 200 {string} json "{"code":200,"data":{"session_id":"xxxxxxxxxxx"},"msg":"ok"}"
// @Router /v1/passwds [get]
func GetPasswd(c *gin.Context) {
	e.SuccessJSON(c, gin.H{"GetPasswd": "GetPasswd"})
}

// @Summary 更新密码信息
// @Tags passwd
// @Produce json
// @Success 200 {string} json "{"code":200,"data":{"session_id":"xxxxxxxxxxx"},"msg":"ok"}"
// @Router /v1/passwds [get]
func EditPasswd(c *gin.Context) {
	e.SuccessJSON(c, gin.H{"EditPasswd": "EditPasswd"})
}

// @Summary 创建密码
// @Tags passwd
// @Produce json
// @Success 200 {string} json "{"code":200,"data":{"session_id":"xxxxxxxxxxx"},"msg":"ok"}"
// @Router /v1/passwds [get]
func AddPasswd(c *gin.Context) {
	e.SuccessJSON(c, gin.H{"AddPasswd": "AddPasswd"})
}
