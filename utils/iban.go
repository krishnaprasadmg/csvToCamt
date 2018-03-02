package utils

import (
	"fmt"
	"os"

	"net/http"

	"encoding/json"

	"io/ioutil"
)

const (
	IBANCachePath = "./iban"
	IBANLookupURL = "https://openiban.com/validate/%s?getBIC=true"
)

type bankData struct {
	BankCode string `json:"bankCode"`
	Name     string `json:"name"`
	Zip      string `json:"zip"`
	City     string `json:"city"`
	Bic      string `json:"bic"`
}

type IBAN struct {
	Iban         string   `json:"iban"`
	BankData     bankData `json:"bankData"`
	CheckResults interface{}
	Valid        interface{}
	Messages     interface{}
}

func Lookup(ibans ...string) map[string]IBAN {
	var ibanData = make(map[string]IBAN)
	var body []byte

	if _, err := os.Stat(IBANCachePath); os.IsNotExist(err) {
		os.Mkdir(IBANCachePath, os.ModePerm)
	}

	for _, iban := range ibans {
		fileName := IBANCachePath + "/" + iban + ".txt"

		if _, err := os.Stat(fileName); os.IsNotExist(err) {
			res, err := http.Get(fmt.Sprintf(IBANLookupURL, iban))
			PanicOnError(err)
			defer res.Body.Close()

			body, err = ioutil.ReadAll(res.Body)
			PanicOnError(err)

			ioutil.WriteFile(fileName, body, os.ModePerm)
		} else {
			body, err = ioutil.ReadFile(fileName)
			PanicOnError(err)
		}

		var ibanItem IBAN
		err := json.Unmarshal(body, &ibanItem)
		PanicOnError(err)

		ibanData[iban] = ibanItem
	}

	return ibanData
}
