package services

import "github.com/ngoctutrang/micro-store/domain/users"

func CreateUser(user users.User) (*users.User, error) {
	return &user, nil
}
