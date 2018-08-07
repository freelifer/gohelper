package repository

import (
	"github.com/freelifer/gohelper/server/repository/database"
)

const (
	// root error no
	SUCCESS           = 0
	SERVER_ERROR      = 10001
	POST_HAS_NO_PARAM = 10002

	// user error no
	LOGIN_ERROR       = 30001
	CREATE_USER_ERROR = 30002
)

type Repository struct {
	Mark  string
	errno int
	data  []byte
}

type UserRepository struct {
	Repository
}

func (repo *UserRepository) GetErrNo() int {
	return repo.errno
}

func (repo *UserRepository) GetData() []byte {
	return repo.data
}

func (repo *UserRepository) Login(name, passwd string) {
	if name == "" || passwd == "" {
		// log.Info(mark, "Login name or password empty")
		repo.errno = POST_HAS_NO_PARAM
		return
	}
	u, err := database.GetUserByName(name)
	if err != nil {
		// log.Info(mark, "Login "+err.Error())
		repo.errno = POST_HAS_NO_PARAM
		return
	}
	if !u.ValidatePassword(passwd) {
		// log.Info(mark, "Login validate password fail")
		repo.errno = LOGIN_ERROR
		return
	}
	repo.errno = SUCCESS

}
