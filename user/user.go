package user

import (
	"github.com/satori/go.uuid"
	"strings"
	"github.com/seantcanavan/tidder/tools"
	"fmt"
)

type User struct {
	Id           string `json: "id"`
	FirstName    string `json: "firstName"`
	LastName     string `json: "lastName"`
	Name         string `json: "name"`
	EmailAddress string `json: "emailAddress"`
}

func NewUser(name, emailAddress string) (*User, error) {

	if !tools.IsValidUserName(name) {
		return nil, fmt.Errorf("cannot create new user with an invalid user name : %v", name)
	}

	if !tools.IsValidEmail(emailAddress) {
		return nil, fmt.Errorf("cannot create new user with an invalid email address: %v", emailAddress)
	}

	newUser := new(User)
	newUser.Name = name
	newUser.Id = uuid.NewV4().String()
	newUser.EmailAddress = emailAddress

	return newUser, nil
}

func (u *User) GetEmailAddress() string {
	return u.EmailAddress
}

func (u *User) SetEmailAddress(emailAddress string) {
	u.EmailAddress = strings.ToLower(emailAddress)
}

func (u *User) GetFullName() string {
	return tools.StrTrimConcat(u.FirstName, u.LastName)
}



