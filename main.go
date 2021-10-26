package main

import (
	"fmt"
)

func main() {
	fmt.Println("Scanneando Porta")
	open := port.InitialScan("tcp", "localhost", 8080)

	fmt.Printf("porta aberta: %t\n", open)

}
