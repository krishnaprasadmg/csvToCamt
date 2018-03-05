package utils

import (
	"log"

	"github.com/spf13/viper"
)

var config *viper.Viper

func PanicOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func ParseConfigFile(file string) {
	config = viper.New()
	config.SetConfigFile(file)
	config.SetConfigType("yaml")
	PanicOnError(config.ReadInConfig())
}
