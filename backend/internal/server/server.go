package server

import (
	"fmt"
	"log"
	"net/http"
)

// Creates the router and start local server
func Start() {

	router := createRouter()

	fmt.Println("Starting server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))

}
