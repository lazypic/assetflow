package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	// db setting

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

	fmt.Println(*flagAdd)
}
