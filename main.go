package main

import (
	"fmt"
	"time"

	"github.com/davidpalves/go-portscanner/port"
	"github.com/fatih/color"
)

var blue = color.New(color.FgBlue, color.Bold).SprintFunc()
var red = color.New(color.FgRed, color.Bold).SprintFunc()
var reverseCyan = color.New(color.FgBlack, color.BgCyan, color.Bold).SprintFunc()
var magenta = color.New(color.FgMagenta)
var yellow = color.New(color.FgHiYellow)

func printResult(hostname string, results []port.ScanResult, elapsedTime float64) {

	magenta.Printf("The open ports in %s ", reverseCyan(hostname))
	magenta.Println("are:")
	for _, result := range results {
		fmt.Printf("- %s:%s\n", blue(hostname), red(result.Port))
	}
	yellow.Printf("\nElapsed time: %.2fs", elapsedTime)
	fmt.Println("")
}

func main() {
	var hostname string = "127.0.0.1"

	startTime := time.Now()
	results := port.StartScan(hostname, "tcp", 0, 49152)
	elapsedTime := time.Since(startTime).Seconds()

	printResult(hostname, results, elapsedTime)
}
