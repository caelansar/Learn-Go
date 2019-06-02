package service

import (
	"errors"
	"fmt"
	"math/rand"
	"webapp/model"
	"webapp/utils"
)

type UserService struct {
}

func (s *UserService) Register (
	username,
	plainpwd string) (user model.User, err error) {

	// check if username has been register already
	var (
		u model.User
	)
	DbEngin.Debug().Where("username=?", username).First(&u)
	if u.ID > 0 {
		return u, errors.New("user exists")
	}
	user.Username = username
	user.Salt = fmt.Sprintf("%06d", rand.Int31n(10000))
	fmt.Println("salt: ", user.Salt)
	user.Pwd = utils.MakePasswd(plainpwd, user.Salt)
	DbEngin.Debug().Create(&user)
	return user, err
}

func (s *UserService) Login (username, password string) (tokenStr string, err error) {
	var (
		user model.User
	)
	DbEngin.Debug().Where("username=?", username).First(&user)
	if user.ID > 0 {
		// check password
		if utils.ValidatePasswd(password, user.Salt, user.Pwd) {
			//generate token for user
			tokenStr, err = utils.GenerateToken(user)
			return

		} else {
			return "", errors.New("wrong password")
		}

	}else {
		return "", errors.New("no such user")
	}

}






