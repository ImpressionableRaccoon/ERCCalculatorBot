package configs

import (
	"log"

	"github.com/spf13/viper"
)

type envConfigs struct {
	TelegramBotToken string `mapstructure:"TELEGRAM_BOT_TOKEN"`
}

func LoadEnvVariables() (config *envConfigs) {
	viper.AddConfigPath(".")
	viper.SetConfigName("bot")
	viper.SetConfigType("env")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("Error reading env file", err)
	}

	if err := viper.Unmarshal(&config); err != nil {
		log.Fatal(err)
	}

	return config
}
