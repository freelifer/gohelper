package e

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorJSON(c *gin.Context, err Err) {
	// log err.Error() 内部错误
	code := err.Code()

	c.JSON(http.StatusOK, gin.H{
		"errcode": code,
		"errmsg":  GetMsg(code),
		// "errmsg": err.Error(),
	})
}

func SuccessJSON(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"errcode": 0,
		"errmsg":  "",
		"data":    data,
	})
}
