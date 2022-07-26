package port

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/cheggaaa/pb/v3"
	"golang.org/x/sync/semaphore"
)

type ScanResult struct {
	Port  int
	State string
}

type ScanConfig struct {
	Hostname            string
	Protocol            string
	LowestPort          int
	HighestPort         int
	ConcurrentOperators int64
}

var tmpl string = `{{ yellow "Ports scanned:" }} {{ bar . "[" "-" ">" "." "]" | green}} {{percent .}} - {{rtime . "Remaining time: %s"}}`

func NewScanConfig(hostname, protocol string, lowestPort, highestPort int, concurrentOperators int64) ScanConfig {
	fmt.Printf("Initializing scan configuration")

	if hostname == "" {
		hostname = "127.0.0.1"
	}

	if protocol == "" {
		protocol = "tcp"
	}

	if highestPort == 0 {
		highestPort = 65535
	}

	if concurrentOperators == 0 {
		concurrentOperators = 32
	}

	var config = ScanConfig{
		Hostname:            hostname,
		Protocol:            protocol,
		LowestPort:          lowestPort,
		HighestPort:         highestPort,
		ConcurrentOperators: concurrentOperators,
	}
	return config
}

func ScanPort(protocol, hostname string, port int) ScanResult {
	result := ScanResult{Port: port}
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 500*time.Millisecond)

	if err != nil {
		result.State = "Closed"
		return result
	}
	defer conn.Close()
	result.State = "Open"
	return result
}

func StartScan(config ScanConfig) []ScanResult {
	var results []ScanResult
	var result ScanResult
	var sem = semaphore.NewWeighted(config.ConcurrentOperators)

	wg := sync.WaitGroup{}

	bar := pb.ProgressBarTemplate(tmpl).Start(config.HighestPort)

	for i := config.LowestPort; i <= config.HighestPort; i++ {
		wg.Add(1)
		sem.Acquire(context.TODO(), 1)
		go func(i int) {
			result = ScanPort(config.Protocol, config.Hostname, i)
			defer bar.Increment()
			defer sem.Release(1)
			defer wg.Done()
			if result.State == "Open" {
				results = append(results, result)
			}
		}(i)
	}

	wg.Wait()
	bar.Finish()

	return results
}
