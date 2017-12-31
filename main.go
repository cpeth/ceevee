package main

import (
	"fmt"

	"github.com/ceevee/ceevee"
)

func main() {
	fmt.Println("Starting CeeVee Server..")
	ceevee.Serve(":8080")
}
