package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var (
	now = time.Now()
	// db setting
	flagRegion  = flag.String("region", "ap-northeast-2", "AWS region name")
	flagProfile = flag.String("profile", "lazypic", "AWS Credentials profile name")
	flagTable   = flag.String("table", "assetflow", "AWS Dynamodb table name")

	// mode and partition key
	flagAdd = flag.String("add", "", "type addition mode")
	//flagUpdate = flag.String("update", "", "type update mode")
	//flagRm     = flag.String("rm", "", "type remove mode")

	// sort key
	flagCreateDate = flag.String("createdate", now.Format(time.RFC3339), "item create date")

	// attributes
	flagProduct        = flag.String("product", "_", "product name")
	flagProductStatus  = flag.String("productstatus", "normal", "product status")
	flagUser           = flag.String("user", "lazypic", "product user name")
	flagSn             = flag.String("sn", "_", "product serial number")
	flagCost           = flag.Int64("cost", 0, "product cost")
	flagMonetaryUnit   = flag.String("monetaryunit", "KRW", "price monetary unit")
	flagPurchaseDate   = flag.String("purchasedate", now.Format(time.RFC3339), "product purchase date")
	flagDescription    = flag.String("description", "_", "description")
	flagURL            = flag.String("url", "https://", "url")
	flagID             = flag.String("id", "_", "id")
	flagPW             = flag.String("pw", "_", "password")
	flagContractDate   = flag.String("contractdate", "_", "contract date")
	flagAddress        = flag.String("address", "_", "address")
	flagMonthlyPayment = flag.Bool("monthlypayment", false, "monthly payment")
	flagYearlyPayment  = flag.Bool("yearlypayment", false, "yearly payment")
	flagFocalLength    = flag.String("focallength", "00mm", "lens focal length")
	flagMacAddress     = flag.String("macaddress", "00:00:00:00:00", "mac address")
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage:\n")
	fmt.Fprintf(os.Stderr, "\tassetflow -add hw -user woong -cost 1400000 -product macbook -description macOS개발머신\n")
	fmt.Fprintf(os.Stderr, "\tassetflow -add sw -user woong -cost 65000 -product hwp2014\n")
	fmt.Fprintf(os.Stderr, "\tassetflow -add account -url https://test.com -id woong -pw pass -cost 5000 -monthlypayment\n")
	fmt.Fprintf(os.Stderr, "\tassetflow -add hw -user woong -cost 1400000 -product macbook -description macOS개발머신\n")
	fmt.Fprintf(os.Stderr, "\tassetflow -add hw -user woong -cost 1400000 -product macbook -description macOS개발머신\n")
	os.Exit(2)
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

	// type 체크 필요함.
	// 모드 체크 필요함
	switch *flagAdd {
	case "hw":
		err := addHw(*db)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
	case "sw":
		err := addSw(*db)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
	case "camera":
		err := addCamera(*db)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
	case "rig":
		err := addRig(*db)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
	case "sound":
		err := addSound(*db)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
	case "lens":
		err := addLens(*db)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
	case "account":
		err := addAccount(*db)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
	case "other":
		err := addOther(*db)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%v\n", err)
			os.Exit(1)
		}
	}
}
