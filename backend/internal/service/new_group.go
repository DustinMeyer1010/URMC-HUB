package service

import (
	"encoding/json"
	"strings"
	"sync"

	"github.com/LostProgrammer1010/URMC-HUB/internal/ad"
)

// Retrieves a group's attributes by DistinguishedName and returns them as JSON.
// If attributes is empty, it returns the DistinguishedName. Returns a NOT_FOUND error, if
// the group is not found
func GetGroup(dn string, attributes ...string) ([]byte, error) {

	attr, _ := ad.LookupGroup(dn, attributes...)

	jsonData, _ := json.Marshal(attr)

	return jsonData, nil

}

// Retrieves a group's attributes by DistinguishedName and returns them as []byte.
// Returns a NOT_FOUND error if group is not found
func GetGroupAvaiableAttributes(dn string) ([]byte, error) {
	attr, _ := ad.LookupGroup(dn, "*")

	var allAttributesNames strings.Builder
	for k := range attr {
		allAttributesNames.WriteString(k + "\n")
	}
	return []byte(allAttributesNames.String()), nil
}

func GetGroupMembers(dn string, start, end int) ([]byte, error) {

	membersDN, cError := ad.GetGroupMembers(dn, start, end)

	if cError != nil {
		return []byte(""), cError
	}

	members := new([]map[string][]string)

	workers := 8
	resutls := make(chan map[string][]string, workers*2)
	var wg sync.WaitGroup
	jobs := make(chan string, workers*2)

	go createJobs(membersDN["members"], jobs)

	for range workers {
		wg.Add(1)
		go startLookupUsersWorkers(&wg, jobs, resutls)
	}

	go cleanUpResultsChannel(&wg, resutls)

	for user := range resutls {
		*members = append(*members, user)
	}

	jsonData, _ := json.Marshal(members)

	return jsonData, nil
}

func cleanUpResultsChannel(wg *sync.WaitGroup, results chan map[string][]string) {
	wg.Wait()
	close(results)
}

func createJobs(membersDN []string, jobs chan string) {

	for _, m := range membersDN {
		jobs <- m
	}
	close(jobs)
}

func startLookupUsersWorkers(wg *sync.WaitGroup, jobs chan string, results chan map[string][]string) {
	defer wg.Done()
	for dn := range jobs {
		user, _ := ad.LookupUser(dn, "cn", "username", "netid", "urid", "dn")

		results <- user
	}
}
