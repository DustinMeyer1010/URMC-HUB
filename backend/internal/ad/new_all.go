package ad

import (
	"fmt"
	"sync"
)

func AllSearchNew(searchValue string) {
	users := make([]map[string][]string, 0)

	var wg sync.WaitGroup
	wg.Add(1)
	go SearchUsers(searchValue, &users, &wg)

	wg.Wait()

	fmt.Println(users)

}

func SearchUsers(searchValue string, users *[]map[string][]string, wg *sync.WaitGroup) {
	defer wg.Done()
	attributes := []string{"username", "netid", "dn", "cn"}
	SearchAllUserNew(users, searchValue, attributes...)

}
