package v1

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type User struct {
	Username, Passwd string
}

// https://www.cnblogs.com/hackyo/p/7992174.html
// @Summary 登录
// @Tags user
// @Accept json
// @Produce json
// @Param  account body User true "User"
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /v1/tokens [post]
func Login(c *gin.Context) {
	var reqInfo User
	err := c.BindJSON(&reqInfo)
	if err != nil {
		log.Printf("%v", err)
		return
	}
	log.Println(reqInfo)
}

// @Summary 注册
// @Tags user
// @Accept application/x-www-form-urlencoded
// @Accept json
// @Produce json
// @Param body body string true "用户名 密码" default(username=1&passwd=2)
// @Success 200 {string} json "{"code":200,"data":{},"msg":"ok"}"
// @Router /v1/register [post]
func Register(c *gin.Context) {
	var reqInfo User
	c.Bind(&reqInfo)
	// username := c.DefaultPostForm("username", "")
	// passwd := c.DefaultPostForm("passwd", "")
	log.Println(reqInfo)
	c.JSON(http.StatusOK, gin.H{"username": "username", "passwd": "passwd"})
}
