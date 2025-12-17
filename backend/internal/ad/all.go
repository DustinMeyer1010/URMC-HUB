package ad

import (
	"sync"

	"github.com/LostProgrammer1010/URMC-HUB/internal/customError"
	"github.com/LostProgrammer1010/URMC-HUB/internal/logger"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
	"github.com/go-ldap/ldap/v3"
)

func AllSearch(search string) (models.AllResults, *customError.Error) {
	result := models.AllResults{}
	result.Users = make([]models.UserSimpleInfo, 0)
	result.Computers = make([]models.ComputerSimpleInfo, 0)
	result.Groups = make([]models.GroupSimpleInfo, 0)
	result.Printers = make([]models.PrinterSimpleInfo, 0)
	result.Shares = make([]models.DriveSimpleInfo, 0)

	ch := make(chan any, 5)
	startGoRouteSearch(search, ch)

	for results := range ch {
		switch results := results.(type) {
		case []models.DriveSimpleInfo:
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
			logger.Errorf("No types found %+v", results)
		}

	}

	return result, nil
}

func thread(wg *sync.WaitGroup, ch chan any, task func() any) {
	defer wg.Done()
	ch <- task()
}

func startGoRouteSearch(search string, ch chan any) {
	var wg sync.WaitGroup

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
		result, err := SearchAllDrives(search)
		if err != nil {
			return make([]models.DriveSimpleInfo, 0)
		}
		return result
	})

	wg.Wait()
	close(ch)
}

func PullMembersInformation(membersDNs []string) []models.UserSimpleInfo {
	const workers = 8

	jobs := make(chan string)
	results := make(chan models.UserSimpleInfo, workers*2)

	var wg sync.WaitGroup

	for range workers {
		wg.Add(1)
		go SearchForObject(&wg, results, jobs)
	}

	go createJobs(membersDNs, jobs)
	go closeResults(&wg, results)

	members := []models.UserSimpleInfo{}
	for username := range results {
		members = append(members, username)
	}

	return members
}

func createJobs(membersDNs []string, jobs chan string) {

	for _, dn := range membersDNs {
		jobs <- dn
	}
	close(jobs)
}

func closeResults(wg *sync.WaitGroup, results chan models.UserSimpleInfo) {
	wg.Wait()
	close(results)
}

func SearchForObject(wg *sync.WaitGroup, results chan models.UserSimpleInfo, jobs chan string) {
	defer wg.Done()

	conn, err := connectToLDAP()
	if err != nil {
		return
	}
	defer conn.Unbind()
	defer conn.Close()

	for dn := range jobs {
		var user models.UserSimpleInfo
		sr, err := conn.Search(ldap.NewSearchRequest(
			dn,
			ldap.ScopeBaseObject,
			ldap.NeverDerefAliases,
			0, 0, false,
			"(|(&(objectCategory=person)(objectClass=user))(objectClass=group))",
			[]string{"cn", "name", "sAMAccountName", "distinguishedName", "uid", "mail", "urid", "objectCategory"},
			nil,
		))

		if err == nil && sr != nil && len(sr.Entries) > 0 {
			user.FillAttributes(sr.Entries[0])
		} else {
			user.OU = dn
		}
		results <- user
	}
}
