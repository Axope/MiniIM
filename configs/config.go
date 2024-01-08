package configs

import (
	"fmt"

	"github.com/spf13/viper"
)

type LogConfig struct {
	Level string `yaml:"level"`
	Path  string `yaml:"path"`
}

type MysqlConfig struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     uint   `yaml:"port"`
	DBname   string `yaml:"DBname"`
	Timeout  string `yaml:"timeout"`
}

type RabbitMQConfig struct {
	Addr         string `yaml:"addr"`
	User         string `yaml:"user"`
	Password     string `yaml:"password"`
	ExchangeName string `yaml:"exchangeName"`
	ExchangeType string `yaml:"exchangeType"`
}

type YamlConfig struct {
	Log   LogConfig
	Mysql MysqlConfig
	RabbitMQ RabbitMQConfig
}

var c YamlConfig

func init() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %s", err))
	}

	viper.Unmarshal(&c)
}

func GetConfig() YamlConfig {
	return c
}
