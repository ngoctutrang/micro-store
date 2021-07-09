package users

import (
	"fmt"

	"github.com/ngoctutrang/micro-store/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	result := usersDB[user.Id]

	if result == nil {
		return errors.NewNotFoundErr(
			fmt.Sprintf("user %d not found", user.Id),
		)
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated
	return nil
}

func (user *User) Save() *errors.RestErr {
	currentUser := usersDB[user.Id]

	if currentUser != nil {
		if user.Email == currentUser.Email {
			return errors.NewBadRequestErr(fmt.Sprintf("user %s already exists", currentUser.Email))
		}
		return errors.NewBadRequestErr(fmt.Sprintf("user %d already exists", currentUser.Id))
	}

	usersDB[user.Id] = user

	return nil
}
