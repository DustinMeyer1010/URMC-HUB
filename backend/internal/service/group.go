package service

import (
	"encoding/json"
	"fmt"
	"strings"
	"sync"
	"time"

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

	startTime := time.Now()

	membersDN, cError := ad.GetGroupMembers(dn, start, end)

	if cError != nil {
		return nil, cError
	}

	members := new([]map[string][]string)
	workers := 16
	results := make(chan map[string][]string, len(membersDN))
	var wg sync.WaitGroup
	jobs := make(chan string, workers*2)

	go createJobs(membersDN["members"], jobs)
	startWorkers(workers, &wg, jobs, results) // Start the worker to start searches and putting items in the reuslts channel
	go cleanUpResultsChannel(&wg, results)

	for user := range results {
		*members = append(*members, user)
	}

	fmt.Println(time.Since(startTime))

	jsonData, _ := json.Marshal(members)

	return jsonData, nil
}

// Closes the results channel once all the workers have complete everything in teh results chan
func cleanUpResultsChannel(wg *sync.WaitGroup, results chan map[string][]string) {
	wg.Wait()
	close(results)
}

// Starts the number of workers based on the number provided.
func startWorkers(nWorkers int, wg *sync.WaitGroup, jobs chan string, results chan map[string][]string) {
	for i := range nWorkers {
		wg.Add(1)
		go startLookupUsersWorkers(wg, jobs, results, i)
	}
}

// Puts all memberDN into a chan for workers to grab from
func createJobs(membersDN []string, jobs chan string) {

	for _, m := range membersDN {
		jobs <- m
	}
	close(jobs)
}

func startLookupUsersWorkers(wg *sync.WaitGroup, jobs chan string, results chan map[string][]string, id int) {
	defer wg.Done()
	for dn := range jobs {
		fmt.Printf("Worker %d working on %s\n", id, dn)
		user, _ := ad.LookupUser(dn, "cn", "username", "netid", "urid", "dn")

		results <- user
	}
}

// TODO: Handle the adding user to group
func UserAddtOGroup(userDN []string, groupDN string) ([]byte, error) {
	return []byte(""), nil
}

// TODO: Handle removing user from group
func UserRemoveFromGroup(userDN []string, groupDN string) ([]byte, error) {
	return []byte(""), nil
}
