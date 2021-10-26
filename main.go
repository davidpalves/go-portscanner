package main

import (
	"fmt"
	"time"

	"github.com/davidpalves/go-portscanner/port"
)

func main() {
	fmt.Println("====== Go Port Scanner ======")
	startTime := time.Now()
	results := port.InitialScan("localhost", false)
	fmt.Printf("Elapsed time %fs\n", time.Since(startTime).Seconds())
	fmt.Println(results)
	startTime = time.Now()
	results = port.WideScan("localhost", false)
	fmt.Printf("Elapsed time %fs\n", time.Since(startTime).Seconds())
	fmt.Println(results)
}
