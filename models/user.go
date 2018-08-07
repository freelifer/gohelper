package models

import (
	"fmt"
	"github.com/freelifer/gohelper/pkg/utils"
	"golang.org/x/crypto/pbkdf2"

	"crypto/sha256"
	"crypto/subtle"
)

type User struct {
	Id      int64
	Name    string `xorm:"UNIQUE NOT NULL"`
	Passwd  string
	Salt    string `xorm:"VARCHAR(10)"`
	Token   string `xor:"VARCHAR(32)"`
	Created int64  `xorm:"created"`
	Updated int64  `xorm:"updated"`
}

func ExistUserByName(name string) (bool, error) {
	if len(name) == 0 {
		return false, nil
	}
	var user User
	_, err := x.Select("id").Where("name = ?", name).Get(&user)
	if err != nil {
		return false, err
	}

	if user.Id > 0 {
		return true, nil
	}
	return false, nil
}

func AddUser(u *User) (err error) {
	if u.Salt, err = GetUserSalt(); err != nil {
		return err
	}

	u.EncodePasswd()

	sess := x.NewSession()
	defer sess.Close()
	if err = sess.Begin(); err != nil {
		return err
	}

	if _, err = sess.Insert(u); err != nil {
		return err
	}

	return sess.Commit()
}

func GetUserByID(id int64) (*User, error) {
	u := new(User)
	has, err := x.Id(id).Get(u)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, fmt.Errorf("user does not exist [user_id: %d, name: %s]", id, "")
	}
	return u, nil
}

func GetUserByName(name string) (*User, error) {
	if len(name) == 0 {
		return nil, fmt.Errorf("user does not exist [user_id: %d, name: %s]", 0, name)
	}
	u := &User{Name: name}
	has, err := x.Get(u)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, fmt.Errorf("user does not exist [user_id: %d, name: %s]", 0, name)
	}
	return u, nil
}

//ValidatePassword
// ValidatePassword checks if given password matches the one belongs to the user.
func (u *User) ValidatePassword(passwd string) bool {
	newUser := &User{Passwd: passwd, Salt: u.Salt}
	newUser.EncodePasswd()
	return subtle.ConstantTimeCompare([]byte(u.Passwd), []byte(newUser.Passwd)) == 1
}

// GetUserSalt returns a ramdom user salt token.
func GetUserSalt() (string, error) {
	return utils.RandomString(10)
}

// EncodePasswd encodes password to safe format.
func (u *User) EncodePasswd() {
	newPasswd := pbkdf2.Key([]byte(u.Passwd), []byte(u.Salt), 10000, 50, sha256.New)
	u.Passwd = fmt.Sprintf("%x", newPasswd)
}
