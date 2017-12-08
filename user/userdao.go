package user

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

const TABLENAME = "User"

func CreateUser(u *User) (*dynamodb.PutItemOutput, error) {
	avm, marshalErr := dynamodbattribute.MarshalMap(u)
	if marshalErr != nil {
		return nil, marshalErr
	}

	dynamo := getDynamoDb()

	pii := &dynamodb.PutItemInput{
		Item:      avm,
		TableName: aws.String(TABLENAME),
	}

	return dynamo.PutItem(pii)
}

func ReadUser(id string) (*dynamodb.GetItemOutput, error) {
	avm := make(map[string]*dynamodb.AttributeValue)
	avm["Id"] = &dynamodb.AttributeValue{
		S: aws.String(id),
	}

	gii := &dynamodb.GetItemInput{
		Key: avm,
		TableName: aws.String(TABLENAME),
	}

	dynamo := getDynamoDb()

	return dynamo.GetItem(gii)
}

func UpdateUser(u *User) (*dynamodb.UpdateItemOutput, error) {

	keyMap := make(map[string]*dynamodb.AttributeValue)
	keyMap["Id"] = &dynamodb.AttributeValue{
		S: aws.String(u.Id),
	}

	avm, marshalErr := ToAvm(u)
	if marshalErr != nil {
		return nil, marshalErr
	}

	dynamo := getDynamoDb()

	uii := &dynamodb.UpdateItemInput{
		Key: keyMap,
		AttributeUpdates: avm,
		TableName: aws.String(TABLENAME),
	}

	return dynamo.UpdateItem(uii)
}

func DeleteUser(id string) (*dynamodb.DeleteItemOutput, error) {
	avm := make(map[string]*dynamodb.AttributeValue)
	avm["Id"] = &dynamodb.AttributeValue{
		S: aws.String(id),
	}

	dynamo := getDynamoDb()

	dii := &dynamodb.DeleteItemInput{
		Key:       avm,
		TableName: aws.String(TABLENAME),
	}

	return dynamo.DeleteItem(dii)

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
		TableName: aws.String(TABLENAME),
	}

	svc := getDynamoDb()

	return svc.DescribeTable(req)
}