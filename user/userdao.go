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

	pii := &dynamodb.PutItemInput{
		Item:      avm,
		TableName: aws.String(TABLENAME),
	}

	return getDynamoDb().PutItem(pii)
}

func ReadUser(id string) (*dynamodb.GetItemOutput, error) {
	avm := make(map[string]*dynamodb.AttributeValue)
	avm["Id"] = &dynamodb.AttributeValue{ S: aws.String(id), }

	gii := &dynamodb.GetItemInput{
		Key:       avm,
		TableName: aws.String(TABLENAME),
	}

	return getDynamoDb().GetItem(gii)
}

func UpdateUser(u *User) (*dynamodb.UpdateItemOutput, error) {
	keyMap := make(map[string]*dynamodb.AttributeValue)
	keyMap["Id"] = &dynamodb.AttributeValue{ S: aws.String(u.Id), }

	avm, marshalErr := ToAvmUpdate(u)
	if marshalErr != nil {
		return nil, marshalErr
	}

	uii := &dynamodb.UpdateItemInput{
		Key:              keyMap,
		AttributeUpdates: avm,
		TableName:        aws.String(TABLENAME),
	}

	return getDynamoDb().UpdateItem(uii)
}

func DeleteUser(id string) (*dynamodb.DeleteItemOutput, error) {
	avm := make(map[string]*dynamodb.AttributeValue)
	avm["Id"] = &dynamodb.AttributeValue{ S: aws.String(id), }

	dii := &dynamodb.DeleteItemInput{
		Key:       avm,
		TableName: aws.String(TABLENAME),
	}

	return getDynamoDb().DeleteItem(dii)
}

// index operations
func ReadUserByEmail(email string) ([]*User, error) {
	qi := &dynamodb.QueryInput{
		TableName: aws.String(TABLENAME),
		IndexName: aws.String("Email-index"),
		KeyConditions: map[string]*dynamodb.Condition{
			"Email": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue {
					{
						S: aws.String(email),
					},
				},
			},
		},
	}

	qo, queryErr := getDynamoDb().Query(qi)
	if queryErr != nil {
		return make([]*User, 0), queryErr
	}

	return FromAvmArray(qo.Items)
}

func BatchReadUsersById(ids[] string) ([]*User, error) {

	var keysToQuery []map[string]*dynamodb.AttributeValue

	for _, element := range ids {
		currentMap := make(map[string]*dynamodb.AttributeValue)
		currentMap["id"] = &dynamodb.AttributeValue{
			S: aws.String(element),
		}

		keysToQuery = append(keysToQuery, currentMap)
	}

	kaa := &dynamodb.KeysAndAttributes{
		ConsistentRead: aws.Bool(true),
		Keys: keysToQuery,
	}

	ri := make(map[string]*dynamodb.KeysAndAttributes)
	ri[TABLENAME] = kaa

	bgii := &dynamodb.BatchGetItemInput{
		RequestItems: ri,
	}

	svc := getDynamoDb()

	bgio, bgioErr := svc.BatchGetItem(bgii)

	if bgioErr != nil {
		return nil, bgioErr
	}
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
