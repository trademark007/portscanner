package main

import (
	"fmt"

	"github.com/trademark007/portscanning/port"
)

func main() {
	fmt.Println("Port Scanning")
	results := port.InitialScan("localhost")
	fmt.Println(results)
}
