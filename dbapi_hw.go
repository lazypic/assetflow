package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func addHw(db dynamodb.DynamoDB) {
	if hasItem(db, *flagTable, *flagAdd, *flagCreateDate) {
		fmt.Fprint(os.Stderr, "The data already exists. Can not add data.\n")
		os.Exit(1)
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
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	data := &dynamodb.PutItemInput{
		Item:      dynamodbJSON,
		TableName: aws.String(*flagTable),
	}
	_, err = db.PutItem(data)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
	fmt.Println("add item:")
	fmt.Println(item)
}
