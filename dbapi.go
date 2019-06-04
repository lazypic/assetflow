package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func tableStruct(tableName string) *dynamodb.CreateTableInput {
	return &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Typ"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("CreateDate"),
				AttributeType: aws.String("S"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Typ"),
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("CreateDate"),
				KeyType:       aws.String("RANGE"),
			},
		},
		BillingMode: aws.String(dynamodb.BillingModePayPerRequest), // ondemand
		TableName:   aws.String(tableName),
	}
}

func validTable(db dynamodb.DynamoDB, tableName string) bool {
	input := &dynamodb.ListTablesInput{}
	isTableName := false
	// 한번에 최대 100개의 테이블만 가지고 올 수 있다.
	// 한 리전에 최대 256개의 테이블이 존재할 수 있다.
	// https://docs.aws.amazon.com/ko_kr/amazondynamodb/latest/developerguide/Limits.html
	for {
		result, err := db.ListTables(input)
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case dynamodb.ErrCodeInternalServerError:
					fmt.Fprintf(os.Stderr, "%s %s\n", dynamodb.ErrCodeInternalServerError, err.Error())
				default:
					fmt.Fprintf(os.Stderr, "%s\n", aerr.Error())
				}
			} else {
				fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			}
			return false
		}

		for _, n := range result.TableNames {
			if *n == tableName {
				isTableName = true
				break
			}
		}
		if isTableName {
			break
		}
		input.ExclusiveStartTableName = result.LastEvaluatedTableName

		if result.LastEvaluatedTableName == nil {
			break
		}
	}
	return isTableName
}
