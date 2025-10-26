package ad

import (
	"fmt"
	"sort"
	"sync"

	"github.com/LostProgrammer1010/URMC-HUB/internal/global"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
)

func LockoutInfoData(user string) (matches []models.LockOutStatus) {
	var servers = global.AllServers()

	var wg sync.WaitGroup // create a wait group

	for _, server := range servers { // loop through servers
		wg.Add(1)
		go func() {
			defer wg.Done()
			result, err := ServerLockout(server, user)
			if err != nil {
				return
			}
			matches = append(matches, result) // append results
		}()
	}

	wg.Wait()

	fmt.Println(matches)

	sort.Slice(matches, func(i, j int) bool {
		return matches[i].Name < matches[j].Name
	})

	return
}

func ServerLockout(server string, user string) (models.LockOutStatus, error) {

	l, err := ConnectToServer("LDAP://" + server)

	if err != nil {
		return models.LockOutStatus{}, err
	}

	defer l.Close()

	config := SearchConfig(
		fmt.Sprintf("(&(objectClass=user)(SAMAccountName=%s*))", user),
		"badPwdCount",
		"badPasswordTime",
	)

	results, _ := config.Search(l)
	if results == nil {
		return models.LockOutStatus{}, fmt.Errorf("no entries found")
	}
	entry := results.Entries[0]

	return models.ToLockOutStatus(server, entry), err
}
