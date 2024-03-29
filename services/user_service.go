package services

import (
	"github.com/ngoctutrang/micro-store/domain/users"
	"github.com/ngoctutrang/micro-store/utils/errors"
)

func CreateUser(user users.User) (*users.User, *errors.RestErr) {
	if err := user.Validate(); err != nil {
		return nil, err
	}

	if err := user.Save(); err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUser(userId int64) (*users.User, *errors.RestErr) {
	user := users.User{Id: userId}

	if err := user.Get(); err != nil {
		return nil, err
	}

	return &user, nil
}
