package main

import (
	"fmt"
	"frame/config"
	"log"
	"net/http"
	"time"

	"go.uber.org/zap"
)

var slogger *zap.SugaredLogger

func mainRoute(w http.ResponseWriter, req *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello!")
}

func testRouter(w http.ResponseWriter, req *http.Request) {
	_, _ = fmt.Fprintf(w, "test")
}

func initConfig() {
	config.ReadConfig()
}

func initLogger() {
	logger, err := zap.NewDevelopment()

	if err != nil {
		log.Fatal(err)
	}

	slogger = logger.Sugar()
}

func main() {
	initConfig()
	initLogger()
	defer slogger.Sync()

	conf := config.GetConfig()

	server := &http.Server{
		Addr:           conf.Addr,
		ReadTimeout:    conf.ReadTimeout * time.Second,
		WriteTimeout:   conf.WriteTimeout * time.Second,
		MaxHeaderBytes: conf.MaxHeaderBytes,
	}

	http.HandleFunc("/", mainRoute)
	http.HandleFunc("/test", testRouter)

	err := server.ListenAndServe()

	if err != nil {
		slogger.Fatal(err)
	}
}
