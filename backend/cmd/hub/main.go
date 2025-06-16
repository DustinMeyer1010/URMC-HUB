package main

import (
	"github.com/LostProgrammer1010/URMC-HUB/internal/db"
)

func main() {
	db.CreateDatabase()
	startServer()
}
