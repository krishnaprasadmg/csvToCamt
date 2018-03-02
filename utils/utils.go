package utils

import (
	"log"

	"github.com/spf13/viper"
)

func PanicOnError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func GetConfig(file string) *viper.Viper {
	config := viper.New()
	config.SetConfigFile(file)
	config.SetConfigType("yaml")
	PanicOnError(config.ReadInConfig())

	return config
}
