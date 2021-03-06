package main

import (
	"fibonacci-api/app"
	"fmt"
	"log"
	"os"
)

func main() {
	//Check that the port was passed
	if len(os.Args) < 2 {
		log.Fatal(fmt.Sprintf("Please provide port on which to start the server...ex: go run main.go 8000"))
	} else {
		port := os.Args[1]
		app.Start(port)
	}
}
