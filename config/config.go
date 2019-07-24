package config

import (
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
	prefix        = "frame"
	envVar        = "env"
	defaultConfig = "default"
	configPath    = "../config"
)

var conf = &Config{}
var environment string

var EnvironmentType = struct {
	Dev  string
	Prod string
}{Dev: "DEV", Prod: "PROD"}

func ReadConfig() {
	viper.AddConfigPath(configPath)

	readDefault()
	readTargetConfig()
}

func defineEnv() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix(prefix)
	env := viper.GetString(envVar)

	if env == EnvironmentType.Prod {
		environment = EnvironmentType.Prod
	} else {
		environment = EnvironmentType.Dev
	}
}

func readDefault() {
	viper.SetConfigName(defaultConfig)

	read()
	unmarshal()
}

func readTargetConfig() {
	viper.SetConfigName(GetEnv())
	read()
	unmarshal()
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

func GetConfig() Config {
	return *conf
}

func GetEnv() string {
	if environment == "" {
		defineEnv()
	}

	return environment
}
