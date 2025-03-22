package config

import (
	"log"

	"github.com/The-EpaG/Go-BasicServer/config"
	"github.com/spf13/viper"
)

func Init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading the configs")
	}

	viper.SetConfigName("secret")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("config/secret")

	if err := viper.MergeInConfig(); err != nil {
		log.Fatal("Error reading the secrets")
	}

	err := viper.Unmarshal(&configurations)
	if err != nil {
		log.Fatal("unable to decode secrets into struct")
	}
}

// to be used only for testing / custom initializations, never during requests
func Set(cfg *config.Config) {
	configurations = *cfg
}
