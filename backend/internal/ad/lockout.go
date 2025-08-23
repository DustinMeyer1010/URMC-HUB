package ad

import (
	"fmt"
	"sort"
	"strconv"
	"sync"
	"time"

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
			matches = append(matches, *result) // append results
		}()
	}

	wg.Wait()

	sort.Slice(matches, func(i, j int) bool {
		return matches[i].Name < matches[j].Name
	})

	return
}

func ServerLockout(server string, user string) (*models.LockOutStatus, error) {
	conn, err := connectToLDAP()

	if err != nil {
		return nil, err
	}
	defer conn.Unbind()
	defer conn.Close()

	ldapConfig := SearchConfig(
		[]string{"badPwdCount", "badPasswordTime"},
		fmt.Sprintf("(&(objectClass=user)(SAMAccountName=%s*))", user),
	)

	results, err := ldapConfig.Search(conn)

	if err != nil {
		return nil, err
	}

	entry := results.Entries[0]

	count, _ := strconv.Atoi(entry.GetAttributeValue("badPwdCount"))

	return &models.LockOutStatus{
		Name:  server,
		Count: count,
		Time:  timeConvert(entry.GetAttributeValue("badPasswordTime")),
	}, nil
}

func timeConvert(input string) (output string) {
	ts, _ := strconv.Atoi(input)
	// Nanoseconds since 1601-01-01
	ticks := int64(ts)
	// Calculate seconds and nanoseconds offset from Unix epoch (1970)
	seconds := ticks/10000000 - 11644473600
	nanoseconds := (ticks % 10000000) * 100
	// Create time.Time object
	t := time.Unix(seconds, nanoseconds)
	if !t.IsDST() {
		t = t.Add(time.Hour)
	}
	if t.Format("2006") == "1600" {
		output = "None"
	} else {
		// Format the time as a string
		output = t.Format("01/02/2006 15:04:05")
	}
	return
}
