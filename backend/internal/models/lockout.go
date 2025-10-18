package models

import (
	"strconv"
	"time"

	"github.com/go-ldap/ldap/v3"
)

// Lockout information for the frontend to display
type LockOutStatus struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
	Time  string `json:"time"`
}

// Convers entry with server to a lockout struct
func ToLockOutStatus(server string, entry *ldap.Entry) LockOutStatus {
	count, _ := strconv.Atoi(entry.GetAttributeValue("badPwdCount"))
	return LockOutStatus{
		Name:  server,
		Count: count,
		Time:  timeConvert(entry.GetAttributeValue("badPasswordTime")),
	}
}

// Converts the time return from the ldap server into a readable time
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
