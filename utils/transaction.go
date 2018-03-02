package utils

import (
	"fmt"

	"strconv"
)

type Transaction struct {
	name     string
	iban     string
	amount   string
	endToEnd string
	bic      string
	bankName string
}

func (t *Transaction) AddToAmount(amount float64) {
	currentAmount, _ := strconv.ParseFloat(t.amount, 64)
	t.amount = fmt.Sprintf("%.2f", currentAmount+amount)
}

func BuildTransactions(investors []Investor) (map[string]*Transaction, float64) {
	var transactions = make(map[string]*Transaction)
	var iban IBAN
	var total float64

	i := 1

	for _, investor := range investors {
		iban = Lookup(investor.iban)[investor.iban]
		total += investor.amount

		if _, ok := transactions[investor.iban]; !ok {
			transactions[investor.iban] = &Transaction{
				investor.name,
				investor.iban,
				"0.0",
				fmt.Sprintf("%d", i),
				iban.BankData.Bic,
				iban.BankData.Name,
			}
			i++
		}
		transactions[investor.iban].AddToAmount(investor.amount)
	}

	return transactions, total
}
