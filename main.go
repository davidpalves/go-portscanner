package main

import (
	"fmt"
	"time"

	"github.com/davidpalves/go-portscanner/port"
	"github.com/fatih/color"
)

func printResult(hostname string, results []port.ScanResult, elapsedTime float64) {
	blue := color.New(color.FgBlue, color.Bold).SprintFunc()
	red := color.New(color.FgRed, color.Bold).SprintFunc()
	reverseCyan := color.New(color.FgBlack, color.BgCyan, color.Bold).SprintFunc()
	magenta := color.New(color.FgMagenta)
	yellow := color.New(color.FgHiYellow)

	magenta.Printf("The open ports in %s", reverseCyan(hostname))
	magenta.Println(" are:")
	for _, result := range results {
		fmt.Printf("- %s:%s\n", blue(hostname), red(result.Port))
	}
	yellow.Printf("\nElapsed time: %.2fs", elapsedTime)
	fmt.Println("")
}

func main() {
	var hostname string = "127.0.0.1"

	startTime := time.Now()
	results := port.WideScan(hostname, false)
	elapsedTime := time.Since(startTime).Seconds()

	printResult(hostname, results, elapsedTime)
}
