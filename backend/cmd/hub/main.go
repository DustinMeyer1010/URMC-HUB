package main

import (
	"github.com/LostProgrammer1010/URMC-HUB/internal/global"
	"github.com/LostProgrammer1010/URMC-HUB/internal/server"
)

func main() {
	global.LoadEnv()
	//db.Init()
	server.Start()

	//models.FindShareDrive("medg_palliative")
}
