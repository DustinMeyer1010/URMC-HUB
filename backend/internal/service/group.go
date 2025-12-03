package service

import (
	"github.com/LostProgrammer1010/URMC-HUB/internal/ad"
	"github.com/LostProgrammer1010/URMC-HUB/internal/customError"
)

func GetAllMembers(group string) ([]string, *customError.Error) {

	return ad.GetAllMembers(group)
}
