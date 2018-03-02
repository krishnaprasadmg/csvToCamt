package utils

import (
	"os"

	"encoding/csv"
	"io"

	"strings"

	"strconv"
)

type Investor struct {
	name   string
	iban   string
	amount float64
}

func LoadInvestors(files []string) []Investor {
	var amount float64
	investors := make([]Investor, 0)

	for _, file := range files {
		investorFile, err := os.OpenFile(file, os.O_RDONLY, os.ModePerm)
		PanicOnError(err)
		defer investorFile.Close()

		reader := csv.NewReader(investorFile)
		reader.Comma = ';'
		reader.FieldsPerRecord = 3

		for {
			record, err := reader.Read()

			if err == io.EOF {
				break
			} else {
				PanicOnError(err)
			}

			amount, _ = strconv.ParseFloat(strings.Replace(record[2], ",", ".", -1), 64)
			investors = append(investors, Investor{record[0], record[1], -1.0 * amount})
		}

	}

	return investors
}
