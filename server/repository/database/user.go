package database

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
	Created int64  `xorm:"created"`
	Updated int64  `xorm:"updated"`
}

// CreateUser add
func CreateUser(u *User) (err error) {
	isExist, err := IsUserExist(0, u.Name)
	if err != nil {
		return err
	} else if isExist {
		return fmt.Errorf("user already exists [name: %s]", u.Name)
	}

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

// IsUserExist checks if given user name exist,
// the user name should be noncased unique.
// If uid is presented, then check will rule out that one,
// it is used when update a user name in settings page.
func IsUserExist(uid int64, name string) (bool, error) {
	if len(name) == 0 {
		return false, nil
	}
	return x.Where("id != ?", uid).Get(&User{Name: name})
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
