package main

import (
	"fmt"
	"github.com/spf13/viper"
	"net/http"
)



func mainRoute(w http.ResponseWriter, req *http.Request) {
	_, _ = fmt.Fprintf(w, "Hello")
}

func testRouter(w http.ResponseWriter, req *http.Request) {
	_, _ = fmt.Fprintf(w, "test")
}

func main() {
	//viper.SetConfigFile("./config/default.yml")
	viper.SetDefault("env", "default")
	viper.SetEnvPrefix("MICRO")
	viper.AutomaticEnv()

	env := viper.GetString("env")

	fmt.Println(env)
	fmt.Println(envConfig[env])

	viper.AddConfigPath("./config/default.yml")
	err := viper.ReadInConfig()

	if err != nil {
		fmt.Println("err", err)
	}

	v := viper.GetString("Server.Addr")
	fmt.Println("v", v)
	//http.HandleFunc("/", mainRoute)
	//http.HandleFunc("/test", testRouter)
	//err = http.ListenAndServe(":3009", nil)
	//log.Fatal(err)
}
