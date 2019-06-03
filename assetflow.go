package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	// db setting

	// mode
	flagAdd    = flag.Bool("add", false, "add mode")
	flagUpdate = flag.Bool("update", false, "update mode")
	flagRm     = flag.Bool("rm", false, "rm mode")

	// key
	flagType       = flag.String("type", "", "item type")              // partition key
	flagCreateDate = flag.String("createdate", "", "item create date") // sort key

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
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage:\n")
	fmt.Fprintf(os.Stderr, "\tassetflow -add ...")
	os.Exit(2)

	/*
		assetflow -add -type hw -user woong -cost 6500000 -product imac
		assetflow -add -type sw -cost 6500000 -product hwp
		assetflow -add -type account -url https://test.com -id woong -pw test -cost 6500000 -monthlypayment
		assetflow -add -type realestate -cost 6500000 -address 주소 -monthlypayment
		assetflow -add -type other -product 안마의자 -cost 35000 -monthlypayment
	*/
}

func main() {
	log.SetPrefix("assetflow: ")
	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	fmt.Println(*flagAdd)
	fmt.Println(*flagType)
}
