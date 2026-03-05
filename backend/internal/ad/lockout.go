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

// Look up Lockout | Bad Password information for the provided user.
// This does a search on all the servers on URMC. This utiles go routes
// for faster searches
func LookupLockoutStatus(userDN string) []models.LockOutStatus {

	lockoutCh := make(chan models.LockOutStatus, 15)

	connectToServers(userDN, lockoutCh)
	lockoutStatuses := drainLockoutChannel(lockoutCh)

	return lockoutStatuses

}

// Reads the lockout channel to pull out the avaiable items in the channel
func drainLockoutChannel(lockoutCh chan models.LockOutStatus) []models.LockOutStatus {
	lockoutStatuses := []models.LockOutStatus{}
	for r := range lockoutCh {
		lockoutStatuses = append(lockoutStatuses, r)
	}

	return lockoutStatuses
}

// Connects to each server getting the lockout information
func connectToServers(userDN string, lockoutCh chan models.LockOutStatus) {
	var wg sync.WaitGroup
	for _, s := range global.AllServers() {
		wg.Add(1)
		go gatherLockoutStatus(s, userDN, lockoutCh, &wg)
	}

	go closeLockoutCh(&wg, lockoutCh)
}

// Cleans up the channel after the wait group finishs
func closeLockoutCh(wg *sync.WaitGroup, lockoutCh chan models.LockOutStatus) {
	wg.Wait()
	close(lockoutCh)
}

// Run the ldap search on each of the server to get the lockout status information
func gatherLockoutStatus(server string, userDN string, lockout chan models.LockOutStatus, waitGroup *sync.WaitGroup) {

	defer waitGroup.Done()
	results, ldapError := gatherLockoutStatusAttributes(server, userDN)

	if ldapError != nil {
		return
	}

	lockout <- results

}

func gatherLockoutStatusAttributes(server, userDN string) (models.LockOutStatus, error) {
	searchConfig := DefaultSearchConfig().
		SetBaseDN(userDN).
		SetAttributes([]string{"badPwdCount", "badPasswordTime"})

	searchResults, ldapError := searchConfig.Search()

	if ldapError != nil {
		fmt.Println(ldapError)
		return models.LockOutStatus{Name: server}, ldapError
	}

	if searchResults == nil || len(searchResults.Entries) == 0 {
		return models.LockOutStatus{Name: server}, fmt.Errorf("no entries found")
	}

	entry := searchResults.Entries[0]

	return models.ToLockOutStatus(server, entry), nil
}

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
