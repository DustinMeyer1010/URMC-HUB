package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/LostProgrammer1010/URMC-HUB/router"
)

// Creates the router and start local server
func Start() {

	router := router.Create()

	fmt.Println("Starting server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))

}
