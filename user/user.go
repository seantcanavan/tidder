package user

import (
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/seantcanavan/tidder/tools"
	"log"
)

type User struct {
	Id    string `json: "id"`
	First string `json: "first"`
	Last  string `json: "last"`
	Name  string `json: "name"`
	Email string `json: "email"`
}

func New(name, emailAddress string) (*User, error) {

	if !tools.IsValidUserName(name) {
		return nil, fmt.Errorf("cannot create new user with an invalid user name : %v", name)
	}

	if !tools.IsValidEmail(emailAddress) {
		return nil, fmt.Errorf("cannot create new user with an invalid email address: %v", emailAddress)
	}

	newUser := new(User)
	newUser.Name = name
	newUser.Id = uuid.NewV4().String()
	newUser.Email = emailAddress

	return newUser, nil
}

func (u *User) SetEmail(email string) {
	if tools.IsValidEmail(email) {
		u.Email = email
	} else {
		log.Printf("cannot set invalid email %v to user %v", email, u)
	}
}

func (u *User) GetFullName() string { return tools.StrTrimConcat(u.First, u.Last) }
