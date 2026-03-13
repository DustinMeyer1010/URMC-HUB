package utils

import (
	"os/exec"
)

func Ping(host string) (string, error) {
	cmd := exec.Command("ping", "-n", "4", host)
	out, err := cmd.CombinedOutput()

	return string(out), err
}

func NSLookUp(host string) (string, error) {
	cmd := exec.Command("nslookup", host)
	out, err := cmd.CombinedOutput()

	return string(out), err
}
