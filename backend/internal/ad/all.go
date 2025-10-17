package ad

import (
	"fmt"
	"sync"

	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
)

func AllSearch(search string) (result models.AllResults, err error) {
	result.Users = make([]models.UserSimpleInfo, 0)
	result.Computers = make([]models.ComputerSimpleInfo, 0)
	result.Groups = make([]models.GroupSimpleInfo, 0)
	result.Printers = make([]models.PrinterSimpleInfo, 0)
	result.Shares = make([]models.ShareDriveSimpleInfo, 0)
	var wg sync.WaitGroup
	ch := make(chan any, 5)

	wg.Add(5)
	go thread(&wg, ch, func() any {
		result, err := SearchAllComputers(search)
		if err != nil {
			return make([]models.ComputerSimpleInfo, 0)
		}
		return result
	})
	go thread(&wg, ch, func() any {
		result, err := SearchAllUsers(search)
		if err != nil {
			return make([]models.UserSimpleInfo, 0)
		}
		return result
	})
	go thread(&wg, ch, func() any {
		result, err := SearchAllGroups(search)
		if err != nil {
			return make([]models.GroupSimpleInfo, 0)
		}
		return result
	})
	go thread(&wg, ch, func() any {
		result, err := SearchAllPrinters(search)
		if err != nil {
			return make([]models.PrinterSimpleInfo, 0)
		}
		return result
	})
	go thread(&wg, ch, func() any {
		result, err := SearchallShareDrives(search)
		if err != nil {
			return make([]models.ShareDriveSimpleInfo, 0)
		}
		return result
	})

	wg.Wait()
	close(ch)

	for results := range ch {
		switch results := results.(type) {
		case []models.ShareDriveSimpleInfo:
			result.Shares = results
		case []models.ComputerSimpleInfo:
			result.Computers = results
		case []models.UserSimpleInfo:
			result.Users = results
		case []models.GroupSimpleInfo:
			result.Groups = results
		case []models.PrinterSimpleInfo:
			result.Printers = results
		default:
			err = fmt.Errorf("unknown type: %T", results)
		}

	}

	return
}

func thread(wg *sync.WaitGroup, ch chan any, task func() any) {
	defer wg.Done()
	ch <- task()
}
