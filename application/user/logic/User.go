package userLogic

import userModel "github.com/alvin0918/gin_api/application/user/model"

func UserLogin(username string, password string) (isTrue bool, err error) {

	var (
		user *userModel.User
		data map[string]string
	)

	user = &userModel.User{
		Password:    password,
		Username:    username,
	}

	if data, err = user.CheckUserAndPassword(); err != nil {
		return
	}

	if len(data) > 0 {
		isTrue = true
	} else {
		isTrue = false
	}

	return
}
