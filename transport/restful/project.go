package restful

// import (
// 	"github.com/freelifer/gohelper/pkg/log"
// 	"github.com/freelifer/gohelper/server"
// 	"github.com/gin-gonic/gin"
// 	"strconv"
// )

// func getProjects(c *gin.Context) {
// 	mark := CreateMark()
// 	server := server.NewProjectServer(mark)
// 	respone := server.List()
// 	if !respone.IsSuccess() {
// 		log.Info(mark, respone.Errmsg)
// 	}
// 	Success1(c, respone.Data)
// }

// func postProject(c *gin.Context) {
// 	mark := CreateMark()
// 	server := server.NewProjectServer(mark)

// 	name := c.PostForm("name")
// 	respone := server.AddProject(name)

// 	if !respone.IsSuccess() {
// 		log.Info(mark, respone.Errmsg)
// 	}
// 	Success1(c, respone.Data)
// }

// func getProject(c *gin.Context) {
// 	mark := CreateMark()
// 	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
// 	server := server.NewProjectServer(mark)
// 	respone := server.GetProject(id)

// 	if !respone.IsSuccess() {
// 		log.Info(mark, respone.Errmsg)
// 	}
// 	Success1(c, respone.Data)
// }

// func updateProject(c *gin.Context) {
// 	mark := CreateMark()
// 	pServer := server.NewProjectServer(mark)

// 	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
// 	content := c.PostForm("content")
// 	respone := pServer.UpdateProject(id, content)

// 	if !respone.IsSuccess() {
// 		log.Info(mark, respone.Errmsg)
// 	}
// 	Success1(c, respone.Data)
// }
