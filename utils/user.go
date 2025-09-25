package utils

import (
	"errors"
	"firstbee/models"
)

type LoginReq struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

func SaveUser(user *models.User) (*models.User, error) {
	hashed, err := HashPassword([]byte(user.Password))
	if err != nil {
		return nil, err
	}
	user.Password = hashed
	id, err := models.O.Insert(user)
	if err != nil {
		return nil, err
	}
	user.Id = uint64(id)
	return user, nil
}

func GetUserByEmail(email string) (*models.User, error) {
	user := models.User{Email: email}
	err := models.O.Read(&user, "Email")
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func Authenticate(login *LoginReq) (*models.User, error) {
	user, err := GetUserByEmail(login.Email)
	if err != nil {
		return nil, err
	}
	if CheckPassword([]byte(user.Password), []byte(login.Password)) {
		return user, nil
	} else {
		err = errors.New("password validation failed.")
		return nil, err
	}
}
