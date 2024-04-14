package main

import (
	"log"
	"github.com/nitin3339/ecom/cmd/api" // Importing the API package
)

func main() {
	// Creating a new API server instance with address ":8080" and nil database connection
	server := api.NewAPIServer(":8080", nil)
	if err := server.Run(); err != nil { // Running the server and checking for errors
		log.Fatal(err)
	}
}