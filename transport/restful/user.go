package restful

// import (
// 	"github.com/freelifer/gohelper/pkg/log"
// 	"github.com/freelifer/gohelper/server"
// 	"github.com/gin-gonic/gin"
// )

// func login(c *gin.Context) {
// 	mark := CreateMark()
// 	name := c.PostForm("name")
// 	passwd := c.PostForm("passwd")
// 	userServer := server.NewUserServer()
// 	response := userServer.Login(name, passwd)
// 	if response.IsSuccess() {
// 		Success(c)
// 	} else {
// 		log.Info(mark, response.Errmsg)
// 		Error(c, response.Errno, mark)
// 	}
// }
