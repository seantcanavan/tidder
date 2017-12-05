package user

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"os"
	"os/user"
)

func GetUserByName(name string) {

}

func GetUserByEmailAddress(emailAddress string) {

}

func GetUserById(id string) {

}

func DeleteUser(id string) {

}

func AddUser(user *user.User) {

	svc := getDynamoDb()

	av, err := dynamodbattribute.MarshalMap(user)
	if err != nil {
		fmt.Println("Unable to marshal User map.")
		fmt.Println(err)
		os.Exit(1)
	}

	input := &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String("users"),
	}

	_, err = svc.PutItem(input)

	if err != nil {
		fmt.Println("Unable to PutItemn:")
		fmt.Println(err)
		os.Exit(1)
	}
}

func getDynamoDb() *dynamodb.DynamoDB {
	session, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-west-2"),
		Credentials: credentials.NewSharedCredentials("", "sean-personal"),
	})

	if err != nil {
		fmt.Println("Unable to connect to dynamodb.")
		fmt.Println(err)
		os.Exit(1)
	}

	return dynamodb.New(session)
}
