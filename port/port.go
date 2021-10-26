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

var sem = semaphore.NewWeighted(1)
var tmpl string = `{{ yellow "Ports scanned:" }} {{ bar . "[" "#" ">" "." "]" | green}} {{percent .}} - {{rtime . "Remaining time: %s"}}`

func ScanPort(protocol, hostname string, port int) ScanResult {
	result := ScanResult{Port: port}
	address := hostname + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout(protocol, address, 60*time.Second)

	if err != nil {
		result.State = "Closed"
		return result
	}
	defer conn.Close()
	result.State = "Open"
	return result
}

func InitialScan(hostname string, runUDP bool) []ScanResult {
	var results []ScanResult
	var result ScanResult
	wg := sync.WaitGroup{}
	const limit = 1024

	if runUDP {
		bar := pb.ProgressBarTemplate(tmpl).Start64(limit)
		for i := 0; i <= limit; i++ {
			wg.Add(1)
			sem.Acquire(context.TODO(), 1)
			go func(i int) {
				result = ScanPort("udp", hostname, i)
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
	}

	bar := pb.ProgressBarTemplate(tmpl).Start(limit)

	for i := 0; i <= limit; i++ {
		wg.Add(1)
		sem.Acquire(context.TODO(), 1)
		go func(i int) {
			result = ScanPort("tcp", hostname, i)
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

func WideScan(hostname string, runUDP bool) []ScanResult {
	var results []ScanResult
	var result ScanResult
	wg := sync.WaitGroup{}

	const limit = 49152

	if runUDP {
		bar := pb.ProgressBarTemplate(tmpl).Start64(limit)
		for i := 0; i <= limit; i++ {
			wg.Add(1)
			sem.Acquire(context.TODO(), 1)
			go func(i int) {
				result = ScanPort("udp", hostname, i)
				defer sem.Release(1)
				defer wg.Done()
				if result.State == "Open" {
					results = append(results, result)
				}
			}(i)
		}

		wg.Wait()
		bar.Finish()
	}

	bar := pb.ProgressBarTemplate(tmpl).Start64(limit)

	for i := 0; i <= limit; i++ {
		wg.Add(1)
		sem.Acquire(context.TODO(), 1)
		go func(i int) {
			result = ScanPort("tcp", hostname, i)
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
