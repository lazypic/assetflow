package main

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func addSound(db dynamodb.DynamoDB) error {
	hasBool, err := hasItem(db, *flagTable, *flagAdd, *flagCreateDate)
	if err != nil {
		return err
	}
	if hasBool {
		return errors.New("The data already exists. Can not add data")
	}
	item := Sound{
		Typ:           *flagAdd,
		CreateDate:    *flagCreateDate,
		Product:       *flagProduct,
		ProductStatus: *flagProductStatus,
		ProductUser:   *flagUser,
		MonetaryUnit:  *flagMonetaryUnit,
		Cost:          *flagCost,
		PurchaseDate:  *flagPurchaseDate,
		Description:   *flagDescription,
	}

	dynamodbJSON, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return err
	}

	data := &dynamodb.PutItemInput{
		Item:      dynamodbJSON,
		TableName: aws.String(*flagTable),
	}
	_, err = db.PutItem(data)
	if err != nil {
		return err
	}
	return nil
}
