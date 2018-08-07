package server

import (
	"github.com/freelifer/gohelper/server/repository/database"
	"github.com/gin-gonic/gin/json"
	_ "github.com/mattn/go-sqlite3"
)

const (
	// root error no
	SUCCESS        = 0
	SERVER_ERROR   = 10001
	PARAMS_ERROR   = 10002
	RESOURCE_ERROR = 10003

	// user error no
	LOGIN_ERROR       = 30001
	CREATE_USER_ERROR = 30002
)

type H map[string]interface{}

type ServerResponse struct {
	Errno  int
	Errmsg string
	Data   []byte
}

type Server struct {
	Mark string
}

type UserServer interface {
	Login(name, passwd string) *ServerResponse
}

type ProjectServer interface {
	List() *ServerResponse
	AddProject(name string) *ServerResponse
	DelProject(id int) *ServerResponse
	UpdateProject(id int64, content string) *ServerResponse
	GetProject(id int64) *ServerResponse
}

func init() {
	err := database.NewEngine()
	if err != nil {
		panic(err)
	}
}

// -------ServerResponse start-----------
func NewResp(errno int, msg string, data []byte) *ServerResponse {
	sp := ServerResponse{Errno: errno, Errmsg: msg, Data: data}
	return &sp
}

func NewRespForSuccess(msg string, data []byte) *ServerResponse {
	sp := ServerResponse{Errno: SUCCESS, Errmsg: msg, Data: data}
	return &sp
}

func NewRespForSuccess1(object H) *ServerResponse {
	object["errno"] = SUCCESS
	jsonBytes, _ := json.Marshal(object)
	sp := ServerResponse{Errno: SUCCESS, Errmsg: "", Data: jsonBytes}
	return &sp
}

func NewRespForParamError1(mark, msg string) *ServerResponse {
	jsonBytes, _ := json.Marshal(H{
		"errno":   PARAMS_ERROR,
		"request": mark,
	})
	sp := ServerResponse{Errno: PARAMS_ERROR, Errmsg: msg, Data: jsonBytes}
	return &sp
}

func NewRespForResourceError(mark, msg string) *ServerResponse {
	jsonBytes, _ := json.Marshal(H{
		"errno":   RESOURCE_ERROR,
		"request": mark,
	})
	sp := ServerResponse{Errno: RESOURCE_ERROR, Errmsg: msg, Data: jsonBytes}
	return &sp
}
func NewRespForParamError(msg string, data []byte) *ServerResponse {
	sp := ServerResponse{Errno: PARAMS_ERROR, Errmsg: msg, Data: data}
	return &sp
}

func (resp *ServerResponse) IsSuccess() bool {
	return resp.Errno == SUCCESS
}

// -------ServerResponse end-----------

func NewUserServer() UserServer {
	impl := UserServerImpl{}
	return &impl
}

func NewProjectServer(mark string) ProjectServer {
	impl := ProjectServerImpl{}
	impl.Mark = mark
	return &impl
}
