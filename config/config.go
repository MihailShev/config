package config

import (
	"fmt"
	"github.com/spf13/viper"
	"log"
	"time"
)

type Config struct {
	Addr           string
	ReadTimeout    time.Duration
	WriteTimeout   time.Duration
	MaxHeaderBytes int
}

const (
	prefix        = "FRAME"
	environment   = "env"
	defaultConfig = "default"
	configPath    = "../config"
)

var conf = &Config{}

var targetConfig = map[string]string{
	"dev":  "dev",
	"prod": "prod",
}

func readDefault() {
	viper.SetConfigName(defaultConfig)

	read()
	unmarshal()
}

func readTargetConfig() {
	env := viper.GetString(environment)
	fmt.Println(env)
	configName, ok := targetConfig[env]

	if ok {
		viper.SetConfigName(configName)
		read()
		unmarshal()
	}

}

func unmarshal() {
	err := viper.Unmarshal(conf)

	if err != nil {
		log.Fatal(err)
	}
}

func read() {
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatal(err)
	}
}

func ReadConfig() {
	viper.AddConfigPath(configPath)
	viper.AutomaticEnv()
	viper.SetEnvPrefix(prefix)

	readDefault()
	readTargetConfig()
}

func GetConfig() Config {
	return *conf
}
