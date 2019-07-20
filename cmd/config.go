package main

import (
	"github.com/spf13/viper"
	"log"
)

const (
	envPrefix = "MICRO"
	devEnv = "default"
	prodEnv = "prod"
	enviroment = "env"
	defaultConfig = "default.yml"
	configPath = "./config"
	// config keys
	addr = "Addr"
)

var envConfig = map[string]string{
	"dev": "dev.yml",
	"prod": "prod.yml",
}

type server struct {
	addr string
	readTimeout int
	writeTimeout int
	idleTimeout int
	maxHeaderBytes int
}

type config struct {
	server
}

func readDefault() {
	viper.SetConfigFile(defaultConfig)
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalln("Cannot read default config")
	}

	con := &config{}

	fillConfig(con)
}

func fillConfig(con *config)  {
	addr := viper.GetString(addr)

	if addr != "" {
		con.addr = addr
	}
}

func readConfig() {
	env := viper.GetString(enviroment)
	configFile, ok := envConfig[env]

	if !ok {
		log.Fatalf("Cannot find config file %s\n", configFile)
	}
}

func initReader() {
	viper.AddConfigPath(configPath)
	viper.SetDefault(enviroment, devEnv)
	viper.SetEnvPrefix(envPrefix)
	viper.AutomaticEnv()
}