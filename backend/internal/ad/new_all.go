package ad

import (
	"sync"

	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
)

func SearchAll(searchValue string) models.AllResultsNew {

	results := new(models.AllResultsNew)

	var wg sync.WaitGroup
	wg.Add(5)
	go SearchUsers(searchValue, &results.Users, &wg)
	go SearchGroups(searchValue, &results.Groups, &wg)
	go SearchComputers(searchValue, &results.Computers, &wg)
	go SearchPrinters(searchValue, &results.Printers, &wg)
	go SearchDrives(searchValue, &results.Drives, &wg)
	wg.Wait()

	return *results

}

func SearchUsers(searchValue string, users *[]map[string][]string, wg *sync.WaitGroup) {
	defer wg.Done()
	attributes := []string{"username", "netid", "dn", "cn", "email", "urid"}
	SearchAllUserNew(users, searchValue, attributes...)
}

func SearchGroups(searchValue string, groups *[]map[string][]string, wg *sync.WaitGroup) {
	defer wg.Done()
	attributes := []string{"samaccountname", "dn", "cn", "information", "description"}
	SearchAllGroupsNew(groups, searchValue, attributes...)
}

func SearchComputers(searchValue string, computers *[]map[string][]string, wg *sync.WaitGroup) {
	defer wg.Done()
	attributes := []string{"cn", "dn", "os", "information", "description"}
	SearchAllComputersNew(computers, searchValue, attributes...)
}

func SearchPrinters(searchValue string, printers *[]models.PrinterSimpleInfo, wg *sync.WaitGroup) {
	defer wg.Done()
	*printers, _ = SearchAllPrinters(searchValue)
}

func SearchDrives(searchValue string, drives *[]models.DriveSimpleInfo, wg *sync.WaitGroup) {
	defer wg.Done()
	*drives, _ = SearchAllDrives(searchValue)
}
