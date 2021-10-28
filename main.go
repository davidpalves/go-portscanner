package main

import (
	"flag"
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
	hostname := flag.String("hostname", "127.0.0.1", "Host that will be scanned. E.g.: 127.0.0.1")
	protocol := flag.String("protocol", "tcp", "Protocol used on the scan. E.g.: tpc")
	lowestPort := flag.Int("lowest-port", 0, "Lowest port used on the scan. E.g.: 0")
	highestPort := flag.Int("highest-port", 65535, "Highest port used on the scan. E.g.: 65535")
	concurrentOperators := flag.Int64("concurrent-operations", 32, "How many operations will occur concurrently. E.g.: 32")

	flag.Parse()

	var config = port.NewScanConfig(*hostname, *protocol, *lowestPort, *highestPort, *concurrentOperators)

	startTime := time.Now()
	results := port.StartScan(config)
	elapsedTime := time.Since(startTime).Seconds()

	printResult(config.Hostname, results, elapsedTime)
}
