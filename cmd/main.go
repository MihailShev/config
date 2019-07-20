package main

import (
	"fmt"
	"net/http"
)

func mainRoute(w http.ResponseWriter, req *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello")
}

func testRouter(w http.ResponseWriter, req *http.Request) {
	_, _ = fmt.Fprintf(w, "test")
}

func main() {
	con := getConfig()
	fmt.Println(con)
	//http.HandleFunc("/", mainRoute)
	//http.HandleFunc("/test", testRouter)
	//err = http.ListenAndServe(":3009", nil)
	//log.Fatal(err)
}
