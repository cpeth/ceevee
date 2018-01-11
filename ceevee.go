package main

import (
	"fmt"
	"net/http"
)

//Serve starts the web server taking the port to run on as a parameter
// ex: Serve(":8080") to start the port on 8080
func Serve(port string) {
	http.ListenAndServe(port, nil)
}

func main() {
	fmt.Println("Starting CeeVee Server..")
	Serve(":8080")
}
