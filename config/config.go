package config

import "github.com/spf13/viper"

type Config struct {
	PORT            string
	DB_USER         string
	DB_PASS         string
	DB_URL          string
	DB_NAME         string
	JWT_SIGNING_KEY string
}

var ENV *Config

func LoadConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName(".env")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&ENV); err != nil {
		panic(err)
	}
}
