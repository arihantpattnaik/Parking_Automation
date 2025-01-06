package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/parking_automation/router"
)

func main() {
	fmt.Println("I AM THE MAIN FILE")
	r := router.Router()

	fmt.Println("Server is getting started...")

	// Start the server on port 4000 and listen for any errors
	log.Fatal(http.ListenAndServe(":4000", r))

	// This line will not be executed due to log.Fatal above,
	// but if you handle errors differently, you could log this
	fmt.Println("Listening at port 4000...")
}
