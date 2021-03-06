package main

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func addAccount(db dynamodb.DynamoDB) error {
	hasBool, err := hasItem(db, *flagTable, *flagAdd, *flagCreateDate)
	if err != nil {
		return err
	}
	if hasBool {
		return errors.New("The data already exists. Can not add data")
	}
	item := Account{
		Typ:            *flagAdd,
		CreateDate:     *flagCreateDate,
		Product:        *flagProduct,
		ProductURL:     *flagURL,
		ID:             *flagID,
		Password:       *flagPW,
		MonetaryUnit:   *flagMonetaryUnit,
		Cost:           *flagCost,
		PurchaseDate:   *flagPurchaseDate,
		Description:    *flagDescription,
		MonthlyPayment: *flagMonthlyPayment,
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
