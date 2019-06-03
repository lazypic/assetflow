package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	flagAdd  = flag.Bool("add", false, "add a item")
	flagType = flag.String("type", "", "item type")
)

func usage() {
	fmt.Fprintf(os.Stderr, "usage:\n")
	fmt.Fprintf(os.Stderr, "\tassetflow -add ...")
	os.Exit(2)
}

func main() {
	log.SetPrefix("assetflow: ")
	log.SetFlags(0)
	flag.Usage = usage
	flag.Parse()

	fmt.Println(*flagAdd)
	fmt.Println(*flagType)
}
