package main

import (
	"fmt"
	"net/http"

	"github.com/jeheskielSunloy77/simpleserver/api"
)

func main() {
	restServer := api.NewServer()
	fmt.Println("yay! The server is running on port 8000!")
	http.ListenAndServe(":8000", restServer)
}
