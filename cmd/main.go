package main

import (
	"fmt"
	"frame/config"
	"log"
	"net/http"
	"time"

	"go.uber.org/zap"
)

func mainRoute(w http.ResponseWriter, req *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello!")
}

func testRouter(w http.ResponseWriter, req *http.Request) {
	_, _ = fmt.Fprintf(w, "test")
}

func initConfig() {
	config.ReadConfig()
}

var logger *zap.Logger

func initLogger() {
	var err error

	env := config.GetEnv()

	if env == config.EnvironmentType.Dev {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}

	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	initConfig()
	initLogger()
	defer logger.Sync()

	conf := config.GetConfig()

	server := &http.Server{
		Addr:           conf.Addr,
		ReadTimeout:    conf.ReadTimeout * time.Second,
		WriteTimeout:   conf.WriteTimeout * time.Second,
		MaxHeaderBytes: conf.MaxHeaderBytes,
	}

	http.HandleFunc("/", mainRoute)
	http.HandleFunc("/test", testRouter)
	logger.Debug("", zap.String("env", config.GetEnv()))

	err := server.ListenAndServe()

	if err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
	}
}
