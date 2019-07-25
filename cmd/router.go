package main

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
)

func mainRoute(w http.ResponseWriter, req *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello!")
}

func testRouter(w http.ResponseWriter, req *http.Request) {
	_, _ = fmt.Fprintf(w, "test")
}

func makeRouter(logger *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {

		logger.Info(req.Method, zap.String("url", req.URL.Path))

		switch req.URL.Path {
		case "/":
			mainRoute(w, req)
		case "/test":
			testRouter(w, req)
		}
	}
}
