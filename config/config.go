package config

import "github.com/spf13/viper"

var config *viper.Viper

func Init() {
	config = viper.New()
	config.SetConfigFile(".env")
	config.ReadInConfig()

	config.SetDefault("PORT", 8080)
}

func GetConfig() *viper.Viper {
	return config
}
