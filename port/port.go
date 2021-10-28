package port

import (
	"context"
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

var sem = semaphore.NewWeighted(1024)
var tmpl string = `{{ yellow "Ports scanned:" }} {{ bar . "[" "#" ">" "." "]" | green}} {{percent .}} - {{rtime . "Remaining time: %s"}}`

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

func StartScan(hostname string, protocol string, lowestPort int, highestPort int) []ScanResult {
	var results []ScanResult
	var result ScanResult
	wg := sync.WaitGroup{}

	if protocol == "" {
		protocol = "tcp"
	}

	bar := pb.ProgressBarTemplate(tmpl).Start(highestPort)

	for i := lowestPort; i <= highestPort; i++ {
		wg.Add(1)
		sem.Acquire(context.TODO(), 1)
		go func(i int) {
			result = ScanPort(protocol, hostname, i)
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
