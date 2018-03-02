package main

import (
	"flag"

	"log"

	"github.com/sunspikes/csvToCamt/utils"
)

type Investor struct {
	name   string
	iban   string
	amount float64
}

func main() {
	configFile := flag.String("c", "", "Config file to use")
	flag.Parse()

	if *configFile == "" {
		log.Fatal("You must supply a config file, use -h for usage")
	}

	if len(flag.Args()) == 0 {
		log.Fatal("You must supply path to CSV files")
	}

	config := utils.GetConfig(*configFile)

	investorData := utils.LoadInvestors(flag.Args())
	transactionData, totalAmount := utils.BuildTransactions(investorData)

	camtDoc := utils.NewCamtDocument(config)
	camtDoc.AddHeaders(totalAmount, len(transactionData))
	camtDoc.AddTransactionData(transactionData)
	camtDoc.PrintDocument()
}
