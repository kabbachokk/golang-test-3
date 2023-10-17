package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	MySQL MySQL
}

type MySQL struct {
	Addr string
	User string
	Pass string
	DB   string
}

// ParseConfig Parse config file
func GetConfig() (*Config, error) {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	viper.SetConfigName("config.yml")

	if err := viper.ReadInConfig(); err != nil {
		return nil, err
	}

	var c Config
	if err := viper.Unmarshal(&c); err != nil {
		return nil, err
	}

	return &c, nil
}
