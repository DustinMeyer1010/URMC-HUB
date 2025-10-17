package ad

import (
	"sort"
	"sync"

	"github.com/LostProgrammer1010/URMC-HUB/internal/global"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
)

func LockoutInfoData(user string) (matches []models.LockOutStatus) {
	var servers = [...]string{global.SERVER1, global.SERVER2, global.SERVER3, global.SERVER4, global.SERVER5, global.SERVER6, global.SERVER7, global.SERVER8, global.SERVER9, global.SERVER10}

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

	sort.Slice(matches, func(i, j int) bool {
		return matches[i].Name < matches[j].Name
	})

	return
}

func ServerLockout(server string, user string) (models.LockOutStatus, error) {

	lockout := models.LockOutStatus{}

	results, err := SearchByCategory(
		"user",
		"SAMAccountName",
		user,
		"badPwdCount",
		"badPasswordTime",
	)

	if err != nil || results == nil {
		return lockout, err
	}

	entry := results.Entries[0]

	return models.ToLockOutStatus(server, entry), err
}
