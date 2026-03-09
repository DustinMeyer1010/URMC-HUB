package service

import (
	"github.com/LostProgrammer1010/URMC-HUB/internal/ad"
	"github.com/LostProgrammer1010/URMC-HUB/internal/utils"
	excel "github.com/xuri/excelize/v2"
)

func GetAllMembers(group string) ([]string, error) {

	return ad.GetAllMembers(group)
}

func GetExcelOfAllMembers(group string) *excel.File {
	membersDNs, _ := GetAllMembers(group)
	members := ad.PullMembersInformation(membersDNs)
	return utils.CreateAllMembersExcel(members)
}
