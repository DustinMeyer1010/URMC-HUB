package ad

import (
	"fmt"
	"sync"

	"github.com/LostProgrammer1010/URMC-HUB/internal/global"
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
