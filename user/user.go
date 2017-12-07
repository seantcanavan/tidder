package user

import (
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/seantcanavan/tidder/tools"
	"log"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type User struct {
	Id    string `json: "id"`
	First string `json: "first"`
	Last  string `json: "last"`
	Name  string `json: "name"`
	Email string `json: "email"`
}

func FromAvm(avm map[string]*dynamodb.AttributeValue) (*User, error) {

	user := &User{}

	if val, ok := avm["Id"]; ok {
		user.Id = *val.S
	}

	if val, ok := avm["First"]; ok {
		user.First = *val.S
	}

	if val, ok := avm["Last"]; ok {
		user.Last = *val.S
	}

	if val, ok := avm["Name"]; ok {
		user.Name = *val.S
	}

	if val, ok := avm["Email"]; ok {
		user.Email = *val.S
	}

	if !IsValidUser(user) {
		return nil, fmt.Errorf("created invalid user %v from map %v", user, avm)
	}

	return user, nil
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

func IsValidUser(u *User) bool {
	if !tools.IsValidEmail(u.Email) {
		return false
	}

	if !tools.IsValidUUID(u.Id) {
		return false
	}

	if u.Name == "" {
		return false
	}

	return true
}