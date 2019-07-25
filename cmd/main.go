package main

import (
	"frame/pkg/config"
	"log"
	"net/http"
	"time"

	"go.uber.org/zap"
)

func main() {
	conf := config.GetConfig()
	logger := makeLogger(conf.Environment)
	defer logger.Sync()

	server := makeServer(conf)

	http.HandleFunc("/", makeRouter(logger))

	logger.Sugar().Info("Server started at port", conf.Addr)

	err := server.ListenAndServe()

	if err != nil {
		logger.Fatal("Failed to start server", zap.Error(err))
	}
}

func makeLogger(env string) *zap.Logger {
	var err error
	var logger *zap.Logger

	if env == config.EnvType.Dev {
		logger, err = zap.NewDevelopment()
	} else {
		logger, err = zap.NewProduction()
	}

	if err != nil {
		log.Fatal(err)
	}

	return logger
}

func makeServer(conf config.Config) *http.Server {
	return &http.Server{
		Addr:           conf.Addr,
		ReadTimeout:    conf.ReadTimeout * time.Second,
		WriteTimeout:   conf.WriteTimeout * time.Second,
		MaxHeaderBytes: conf.MaxHeaderBytes,
	}
}
