package user

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/satori/go.uuid"
	"github.com/seantcanavan/tidder/tools"
	"log"
	"github.com/seantcanavan/tidder/test"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type User struct {
	Id       string `json: "id"`
	First    string `json: "first"`
	Last     string `json: "last"`
	Name     string `json: "name"`
	Email    string `json: "email"`
	Password []byte `JSON: "password"`
}

func FromAvmArray(avm []map[string]*dynamodb.AttributeValue) ([]*User, error) {
	list := make([]*User, len(avm))

	for index, element := range avm {
		from, fromErr := FromAvm(element)
		if fromErr != nil {
			return []*User{}, fmt.Errorf("issue parsing item FromAvm: %v", fromErr.Error())
		}
		list[index] = from
	}

	return list, nil
}

func FromAvm(avm map[string]*dynamodb.AttributeValue) (*User, error) {
	user := &User{}

	err := dynamodbattribute.UnmarshalMap(avm, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func ToAvm(u *User) (map[string]*dynamodb.AttributeValue, error) {
	if !IsValidUser(u) {
		return nil, fmt.Errorf("cannot convert invalid user %v to map", u)
	}

	avm, err := dynamodbattribute.MarshalMap(u);
	if err != nil {
		return nil, err
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

func New(name, email string) (*User, error) {

	if !tools.IsValidName(name) {
		return nil, fmt.Errorf("cannot create new user with an invalid user name : %v", name)
	}

	if !tools.IsValidEmail(email) {
		return nil, fmt.Errorf("cannot create new user with an invalid email address: %v", email)
	}

	newUser := new(User)
	newUser.Name = name
	newUser.Id = uuid.NewV4().String()
	newUser.Email = email

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

func TestUser() *User {
	return &User{
		Id: uuid.NewV4().String(),
		Name: test.RandomAlphaMixed(10),
		Last: test.RandomAlphaMixed(10),
		First: test.RandomAlphaMixed(10),
		Email: test.RandomEmail(),
	}
}