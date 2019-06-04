package main

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func addHw(db dynamodb.DynamoDB) error {
	b, err := hasItem(db, *flagTable, *flagAdd, *flagCreateDate)
	if err != nil {
		return err
	}
	if b {
		return errors.New("The data already exists. Can not add data")
	}
	item := Hw{
		Typ:            *flagAdd,
		CreateDate:     *flagCreateDate,
		Product:        *flagProduct,
		ProductStatus:  *flagProductStatus,
		ProductUser:    *flagUser,
		MonetaryUnit:   *flagMonetaryUnit,
		Cost:           *flagCost,
		PurchaseDate:   *flagPurchaseDate,
		Description:    *flagDescription,
		MonthlyPayment: *flagMonthlyPayment,
	}

	// 데이터 저장
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
