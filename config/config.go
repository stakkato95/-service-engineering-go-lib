package config

import (
	"github.com/spf13/viper"
	"github.com/stakkato95/service-engineering-go-lib/logger"
)

func Init(config interface{}, empty interface{}) {
	viper.AddConfigPath(".")
	viper.SetConfigName("app")
	viper.SetConfigType("env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		logger.Panic("config not found")
	}

	if err := viper.Unmarshal(config); err != nil {
		logger.Panic("config can not be read")
	}

	if config == empty {
		logger.Panic("config is emtpy")
	}
}
