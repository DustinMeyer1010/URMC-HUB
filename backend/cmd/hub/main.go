package main

import (
	"github.com/LostProgrammer1010/URMC-HUB/internal/global"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
)

func main() {
	//db.Init()
	//server.Start()
	global.LoadEnv()
	models.FindShareDrive("medg_palliative")
}
