package config

import (
	"github.com/spf13/viper"
	"log"
)

type Config struct {
	HTTP
}

type HTTP struct {
	Port string
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	viper.AddConfigPath("/Users/homidov/Desktop/wb_l2/11")
	viper.SetConfigName("config")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalln(err)
	}
	cfg.Port = viper.GetString("port")
	return cfg, nil
}
