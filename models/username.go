package models

type UserName struct {
	Id      int64
	Name    string
	Created int64 `xorm:"created"`
	Updated int64 `xorm:"updated"`
}

// add
func AddUserName(u *UserName) (err error) {
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

// delete

// update
func EditUserName(u *UserName) (err error) {
	if _, err = x.Id(u.Id).Update(u); err != nil {
		return err
	}

	return nil
}

// get
func GetUserNameById(id int64) (*UserName, error) {
	u := new(UserName)
	has, err := x.Id(id).Get(u)
	if err != nil {
		return nil, err
	} else if !has {
		return nil, fmt.Errorf("username does not exist [id: %d, name: %s]", id, "")
	}
	return u, nil
}
