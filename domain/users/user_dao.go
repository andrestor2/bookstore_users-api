package users

import (
	"fmt"
	"github.com/andrestor2/bookstore_users-api/utils/date_utils"
	"github.com/andrestor2/bookstore_users-api/utils/errors"
)

var (
	usersDB = make(map[int64]*User)
)

func (user *User) Get() *errors.RestErr {
	result := usersDB[user.Id]
	if result == nil {
		return errors.NewNotFoundError(fmt.Sprintf("User %d not found", user.Id))
	}
	user.Id = result.Id
	user.FirstName = result.FirstName
	user.LastName = result.LastName
	user.Email = result.Email
	user.DateCreated = result.DateCreated

	return nil
}

func (user *User) Save() *errors.RestErr {
	current := usersDB[user.Id]
	if current != nil {
		if current.Email == user.Email {
			return errors.NewBadRequestError(fmt.Sprintf("Email %d already exist", user.Id))
		}
		return errors.NewBadRequestError(fmt.Sprintf("User %d already exist", user.Id))
	}

	user.DateCreated = date_utils.GetNowString()
	usersDB[user.Id] = user
	return nil
}