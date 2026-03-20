package service

import (
	"github.com/LostProgrammer1010/URMC-HUB/internal/ad"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
)

func Login(user models.UserLogin) error {

	if cError := ad.Login(user); cError != nil {
		return cError
	}

	return nil
}
