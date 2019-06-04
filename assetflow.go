package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var (
	// db setting
	flagRegion  = flag.String("region", "ap-northeast-2", "AWS region name")
	flagProfile = flag.String("profile", "lazypic", "AWS Credentials profile name")
	flagTable   = flag.String("table", "assetflow", "AWS Dynamodb table name")

	// mode and partition key
	flagAdd    = flag.String("add", "", "type addition mode")
	flagUpdate = flag.String("update", "", "type update mode")
	flagRm     = flag.String("rm", "", "type remove mode")

	// sort key
	flagCreateDate = flag.String("createdate", "", "item create date")

	// attributes
	flagProduct        = flag.String("product", "", "product name")
	flagProductStatus  = flag.String("productstatus", "", "product status")
	flagUser           = flag.String("user", "", "product user name")
	flagSn             = flag.String("sn", "", "product serial number")
	flagCost           = flag.Float64("cost", 0.0, "product cost")
	flagMonetaryUnit   = flag.String("monetaryunit", "KRW", "price monetary unit")
	flagPurchaseDate   = flag.String("purchasedate", "", "product purchase date")
	flagDescription    = flag.String("description", "", "description")
	flagURL            = flag.String("url", "", "url")
	flagID             = flag.String("id", "", "id")
	flagPassword       = flag.String("password", "", "password")
	flagContractDate   = flag.String("contractdate", "", "contract date")
	flagAddress        = flag.String("address", "", "address")
	flagMonthlyPayment = flag.Bool("monthlypayment", false, "monthly payment")
	flagYearlyPayment  = flag.Bool("yearlypayment", false, "yearly payment")
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage:\n")
	fmt.Fprintf(os.Stderr, "\tassetflow -add ...")
	os.Exit(2)

	/*
		assetflow -add hw -user woong -cost 6500000 -product imac
		assetflow -add sw -cost 6500000 -product hwp
		assetflow -add account -url https://test.com -id woong -pw test -cost 6500000 -monthlypayment
		assetflow -add realestate -cost 6500000 -address 주소 -monthlypayment
		assetflow -add other -product 안마의자 -cost 35000 -monthlypayment
	*/
}

func main() {
	log.SetPrefix("assetflow: ")
	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
		Config:            aws.Config{Region: aws.String(*flagRegion)},
		Profile:           *flagProfile,
	}))
	db := dynamodb.New(sess)

	// 테이블이 존재하는지 점검하고 없다면 테이블을 생성한다.
	if !validTable(*db, *flagTable) {
		_, err := db.CreateTable(tableStruct(*flagTable))
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
		fmt.Println("Created table:", *flagTable)
		fmt.Println("Please try again in one minute.")
		os.Exit(0)
	}

	if *flagAdd != "" {
		fmt.Println(*flagAdd)
	}
}
