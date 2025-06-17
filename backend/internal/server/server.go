package server

import (
	"fmt"
	"log"
	"net/http"
)

// Creates the router and start local server
func Start() {

	router := createRouter()

	fmt.Println("Starting server on port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))

}
