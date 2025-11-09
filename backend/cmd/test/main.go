package main

import (
	"fmt"

	"github.com/LostProgrammer1010/URMC-HUB/internal/service"
)

func main() {
	err := service.PingPrinter("172.16.244.107")

	fmt.Println(err)
}
