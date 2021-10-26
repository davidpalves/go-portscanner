package main

import (
	"fmt"

	"github.com/davidpalves/go-portscanner/port"
)

func main() {
	fmt.Println("Scanneando Porta")
	results := port.InitialScan("localhost")
	fmt.Println(results)

}
