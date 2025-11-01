package service

import (
	"fmt"

	"github.com/LostProgrammer1010/URMC-HUB/internal/ad"
	"github.com/LostProgrammer1010/URMC-HUB/internal/db"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
)

func Login(user models.UserLogin) (err error) {

	if err = ad.Login(user); err != nil {
		fmt.Println(err)
		return
	}

	if err = db.AgentDatabaseInit(); err != nil {
		fmt.Println(err)
		return
	}

	return
}
