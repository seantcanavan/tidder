package user

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
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

func FromAvm(avm map[string]*dynamodb.AttributeValue) (*User, error) {

	u := &User{}

	if val, ok := avm["Id"]; ok {
		u.Id = *val.S
	}

	if val, ok := avm["First"]; ok {
		u.First = *val.S
	}

	if val, ok := avm["Last"]; ok {
		u.Last = *val.S
	}

	if val, ok := avm["Name"]; ok {
		u.Name = *val.S
	}

	if val, ok := avm["Email"]; ok {
		u.Email = *val.S
	}

	if !IsValidUser(u) {
		return nil, fmt.Errorf("created invalid user %v from map %v", u, avm)
	}

	return u, nil
}

func ToAvm(u *User) (map[string]*dynamodb.AttributeValue, error) {

	if !IsValidUser(u) {
		return nil, fmt.Errorf("cannot convert invalid user %v to map", u)
	}

	avm := make(map[string]*dynamodb.AttributeValue)

	if u.Id != "" {
		avm["Id"] = &dynamodb.AttributeValue{S: aws.String(u.Id)}
	}

	if u.First != "" {
		avm["First"] = &dynamodb.AttributeValue{S: aws.String(u.First)}
	}

	if u.Last != "" {
		avm["Last"] = &dynamodb.AttributeValue{S: aws.String(u.Last)}
	}

	if u.Email != "" {
		avm["Email"] = &dynamodb.AttributeValue{S: aws.String(u.Email)}
	}

	if u.Name != "" {
		avm["Name"] = &dynamodb.AttributeValue{S: aws.String(u.Name)}
	}

	return avm, nil
}

func ToAvmUpdate(u *User) (map[string]*dynamodb.AttributeValueUpdate, error) {

	if !IsValidUser(u) {
		return nil, fmt.Errorf("cannot convert invalid user %v to map", u)
	}

	avm := make(map[string]*dynamodb.AttributeValueUpdate)

	// do not marshal the Id since it cannot be part of a valid update command from dynamo

	if u.First != "" {
		avm["First"] = &dynamodb.AttributeValueUpdate{
			Value:  &dynamodb.AttributeValue{S: aws.String(u.First)},
			Action: aws.String("PUT"),
		}
	}

	if u.Last != "" {
		avm["Last"] = &dynamodb.AttributeValueUpdate{
			Value:  &dynamodb.AttributeValue{S: aws.String(u.Last)},
			Action: aws.String("PUT"),
		}
	}

	if u.Email != "" {
		avm["Email"] = &dynamodb.AttributeValueUpdate{
			Value:  &dynamodb.AttributeValue{S: aws.String(u.Email)},
			Action: aws.String("PUT"),
		}
	}

	if u.Name != "" {
		avm["Name"] = &dynamodb.AttributeValueUpdate{
			Value:  &dynamodb.AttributeValue{S: aws.String(u.Name)},
			Action: aws.String("PUT"),
		}
	}

	return avm, nil
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
