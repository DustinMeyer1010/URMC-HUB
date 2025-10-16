package main

import (
	"github.com/LostProgrammer1010/URMC-HUB/internal/db"
	"github.com/LostProgrammer1010/URMC-HUB/internal/server"
)

func main() {
	db.Init()
	server.Start()
}
