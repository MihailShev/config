package router

import (
	"fmt"
	"go.uber.org/zap"
	"net/http"
)

func MakeRouter(logger *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var response string

		logger.Info(req.Method, zap.String("url", req.URL.Path))

		switch req.URL.Path {
		case "/":
			w.WriteHeader(200)
			response = "Hello!"
		case "/test":
			w.WriteHeader(200)
			response = "test"
		default:
			w.WriteHeader(404)
			response = "Route not found"
		}

		_, err := fmt.Fprintln(w, response)

		if err != nil {
			logger.Error("", zap.Error(err))
		}
	}
}
