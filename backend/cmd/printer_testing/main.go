package main

import (
	"bytes"
	"fmt"
	"os/exec"
)

func main() {
	printer := `\\Print-HPc\ISD013`

	cmd := exec.Command("powershell", "-Command",
		fmt.Sprintf(`Get-PrintJob -PrinterName "%s" | Select-Object Id,DocumentName,UserName,PagesPrinted,TotalPages`, printer),
	)

	var out bytes.Buffer
	cmd.Stdout = &out

	err := cmd.Run()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(out.String())
}
