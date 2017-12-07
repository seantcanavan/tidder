package user

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var TABLE_NAME = "User"

func DeleteUser(id string) (*dynamodb.DeleteItemOutput, error) {
	avm := make(map[string]*dynamodb.AttributeValue)
	avm["Id"] = &dynamodb.AttributeValue{
		S: aws.String(id),
	}

	svc := getDynamoDb()

	dii := &dynamodb.DeleteItemInput{
		Key:       avm,
		TableName: aws.String(TABLE_NAME),
	}

	return svc.DeleteItem(dii)

}

func AddUser(user *User) (*dynamodb.PutItemOutput, error) {
	avm, marshalErr := dynamodbattribute.MarshalMap(user)
	if marshalErr != nil {
		return nil, marshalErr
	}

	svc := getDynamoDb()

	pii := &dynamodb.PutItemInput{
		Item:      avm,
		TableName: aws.String(TABLE_NAME),
	}

	return svc.PutItem(pii)
}

func getDynamoDb() *dynamodb.DynamoDB {
	newSession, err := session.NewSession(&aws.Config{
		Region:      aws.String("us-west-2"),
		Credentials: credentials.NewSharedCredentials("", "sean-personal"),
	})

	if err != nil {
		panic(fmt.Sprintf("unable to connect to dynamo: %v", err.Error()))
	}

	return dynamodb.New(newSession)
}

func DescribeTable() (*dynamodb.DescribeTableOutput, error) {
	req := &dynamodb.DescribeTableInput{
		TableName: aws.String(TABLE_NAME),
	}

	svc := getDynamoDb()

	return svc.DescribeTable(req)
}