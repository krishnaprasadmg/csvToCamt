package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/sunspikes/csvToCamt/utils"
)

func init() {
	flag.Usage = func() {
		fmt.Printf("Usage of %s:\n", os.Args[0])
		fmt.Printf("%s -c config.yaml file1.csv file2.csv...\n", os.Args[0])
		flag.PrintDefaults()
	}
}

func main() {
	var configFile string
	var skipHeader bool
	flag.StringVar(&configFile, "c", "", "Config file to use, see config.yaml for example")
	flag.BoolVar(&skipHeader, "s", false, "Skip CSV header lines, off by default")
	flag.Parse()

	if configFile == "" || len(flag.Args()) == 0 {
		flag.Usage()
		fmt.Fprintln(os.Stderr, "\nInsufficient arguments!")
		os.Exit(1)
	}

	utils.ParseConfigFile(configFile)

	investorData := utils.LoadInvestors(flag.Args(), skipHeader)
	transactionData, totalAmount := utils.BuildTransactions(investorData)

	camtDoc := utils.NewCamtDocument()
	camtDoc.AddHeaders(totalAmount, len(transactionData))
	camtDoc.AddTransactionData(transactionData)
	camtDoc.PrintDocument()
}
