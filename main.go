package main

import (
	"fmt"

	"github.com/davidpalves/go-portscanner/port"
)

func main() {
	fmt.Println("Scanneando Porta")
	results := port.InitialScan("localhost")
	fmt.Println(results)
	fmt.Println("\n============================\n")
	wideScanResults := port.WideScan("localhost")
	fmt.Println(wideScanResults)
}
