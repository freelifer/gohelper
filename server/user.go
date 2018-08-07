package server

import (
	"github.com/freelifer/gohelper/server/repository/database"
)

type UserServerImpl struct {
}

func (repo *UserServerImpl) Login(name, passwd string) *ServerResponse {
	if name == "" || passwd == "" {
		return NewRespForParamError("Login name or password empty", nil)
	}
	u, err := database.GetUserByName(name)
	if err != nil {
		return NewRespForParamError("Login "+err.Error(), nil)
	}
	if !u.ValidatePassword(passwd) {
		return NewResp(LOGIN_ERROR, "Login validate password fail", nil)
	}
	return NewRespForSuccess("", nil)
}

func () CreateUser(name, passwd string) *ServerResponse {
	if name == "" || passwd == "" {
		return NewRespForParamError("Login name or password empty", nil)
	}
	return nil
}
