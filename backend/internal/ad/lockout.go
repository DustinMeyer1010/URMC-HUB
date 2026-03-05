package ad

import (
	"fmt"
	"sort"
	"sync"

	"github.com/LostProgrammer1010/URMC-HUB/internal/customError"
	"github.com/LostProgrammer1010/URMC-HUB/internal/global"
	"github.com/LostProgrammer1010/URMC-HUB/internal/logger"
	"github.com/LostProgrammer1010/URMC-HUB/internal/models"
)

// Deprecated: This is being replace with LookupLockoutStatus
func LockoutInfoData(user string) (matches []models.LockOutStatus) {
	var servers = global.AllServers()

	var wg sync.WaitGroup

	for _, server := range servers {
		wg.Add(1)
		go func() {
			defer wg.Done()
			result, err := ServerLockout(server, user)
			if err != nil {
				logger.Error(err)
				return
			}
			matches = append(matches, result)
		}()
	}

	wg.Wait()

	sort.Slice(matches, func(i, j int) bool {
		return matches[i].Name < matches[j].Name
	})

	return
}

// Deprecated: This is being replace with gatherLockoutStatusAttributes
func ServerLockout(server string, user string, attr ...string) (models.LockOutStatus, *customError.Error) {

	l, cError := ConnectToServer("LDAP://" + server)

	if cError != nil {
		return models.LockOutStatus{}, cError
	}
	defer l.Close()
	defer l.Unbind()

	/*
		"badPwdCount",
		"badPasswordTime",
	*/

	config := DefaultSearchConfig().SetFilter(fmt.Sprintf("(&(objectClass=user)(SAMAccountName=%s*))", user)).SetAttributes(attr)

	results, ldapError := config.Search()

	if ldapError != nil {
		logger.Error(ldapError)
		cError := customError.LDAP_ERROR.NewError(ldapError)
		return models.LockOutStatus{}, &cError
	}

	if results == nil {
		cError := customError.NOT_FOUND.NewMessage(fmt.Sprintf("NO USER FOUND FOR: %s", user))
		return models.LockOutStatus{}, &cError
	}

	entry := results.Entries[0]

	return models.ToLockOutStatus(server, entry), nil
}
