package utils

import (
	"fmt"
	"os"

	"net/http"

	"encoding/json"

	"io/ioutil"
)

const (
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

	cacheDirPath := config.GetString("iban.cachePath")

	if _, err := os.Stat(cacheDirPath); os.IsNotExist(err) {
		os.Mkdir(cacheDirPath, os.ModePerm)
	}

	for _, iban := range ibans {
		fileName := cacheDirPath + "/" + iban + ".txt"

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
