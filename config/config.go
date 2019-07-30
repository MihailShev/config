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
	Environment    string
}

const (
	prefix        = "frame"
	envVar        = "env"
	defaultConfig = "default"
	configPath    = "../config"
)

var EnvType = struct {
	Dev  string
	Prod string
}{Dev: "DEV", Prod: "PROD"}

var configFiles = map[string]string{
	EnvType.Dev:  "dev",
	EnvType.Prod: "prod",
}

var conf = &Config{}

func GetConfig() Config {
	if conf.Environment == "" {
		readConfig()
	}
	return *conf
}

func readConfig() {
	defineEnv()
	readDefault()
	readTargetConfig()
}

func defineEnv() {
	viper.AutomaticEnv()
	viper.SetEnvPrefix(prefix)

	env := viper.GetString(envVar)

	switch env {
	case EnvType.Prod:
		conf.Environment = EnvType.Prod
	default:
		conf.Environment = EnvType.Dev
	}
}

func readDefault() {
	viper.AddConfigPath(configPath)
	viper.SetConfigName(defaultConfig)

	read()
	unmarshal()
}

func readTargetConfig() {
	configName, ok := configFiles[conf.Environment]

	if ok {
		viper.SetConfigName(configName)
		read()
		unmarshal()
	} else {
		log.Fatal("Cannot read target config", configName)
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
