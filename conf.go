package main

import (
	"log"

	"github.com/spf13/viper"
)

func Bite(key string) string {
	viper.SetConfigName("config") // config.env
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal("Error while reading config file: ", err)
	}
	value, ok := viper.Get(key).(string)
	if !ok {
		log.Fatal("Invalid type assertion")
	}
	return value

}
