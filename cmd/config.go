package main

import (
	"github.com/spf13/viper"
	"log"
)

const (
	envPrefix     = "MICRO"
	devEnv        = "default"
	environment   = "env"
	defaultConfig = "default.yml"
	configPath    = "./config"

	// config keys
	addr           = "Addr"
	readTimeout    = "ReadTimeout"
	writeTimeout   = "WriteTimeout"
	maxHeaderBytes = "MaxHeaderBytes"
)

var envConfig = map[string]string{
	"dev":  "dev.yml",
	"prod": "prod.yml",
}

type config struct {
	addr           string
	readTimeout    int
	writeTimeout   int
	idleTimeout    int
	maxHeaderBytes int
}

func getConfig() *config {
	initReader()
	con := &config{}
	readDefault(con)
	readConfig(con)

	return con
}

func readDefault(con *config) {
	viper.SetConfigFile(defaultConfig)
	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalln("Cannot read default config")
	}

	fillConfig(con)
}

func readConfig(con *config) {
	env := viper.GetString(environment)
	configFile, ok := envConfig[env]

	if !ok {
		log.Fatalf("Cannot find config file %s\n", configFile)
	}

	viper.SetConfigFile(configFile)

	fillConfig(con)
}

func fillConfig(con *config) {
	addr := viper.GetString(addr)
	readTimeout := viper.GetInt(readTimeout)
	writeTimeout := viper.GetInt(writeTimeout)
	maxHeaderBytes := viper.GetInt(maxHeaderBytes)

	if addr != "" {
		con.addr = addr
	}

	if readTimeout > 0 {
		con.readTimeout = readTimeout
	}

	if writeTimeout > 0 {
		con.writeTimeout = writeTimeout
	}

	if maxHeaderBytes > 0 {
		con.maxHeaderBytes = maxHeaderBytes
	}
}

func initReader() {
	viper.AddConfigPath(configPath)
	viper.SetDefault(environment, devEnv)
	viper.SetEnvPrefix(envPrefix)
	viper.AutomaticEnv()
}
