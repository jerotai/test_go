package config

import (
	"github.com/spf13/viper"
)

//建構子
type (
	Config struct {
	}
)

func (c *Config) Init() {
	viper.AddConfigPath("routes/config")
}

func (c *Config)GetConfig(key string) interface{} {
	return viper.Get(key)
}