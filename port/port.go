package port

import (
	"fmt"
	"net"
	"runtime"
	"strconv"
	"sync"
	"time"
)

var mu sync.Mutex
var wg sync.WaitGroup

type ScanResult struct {
	Port  int
	State string
}

func Rotinas(start int, end int) {

}

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

func InitialScan(hostname string) []ScanResult {
	start := time.Now()
	var results []ScanResult
	var result ScanResult
	i := 0
	portRange := 1024
	routinesCount := runtime.NumCPU()
	indexRange := portRange / routinesCount
	wg.Add(routinesCount)
	for j := 1; j <= routinesCount; j++ {
		go func() {
			currentIndex := j
			for i < indexRange*currentIndex {
				mu.Lock()
				result = ScanPort("tcp", hostname, i)
				i++
				mu.Unlock()
				if result.State == "Open" {
					results = append(results, result)
				}
			}
			wg.Done()
		}()
	}

	wg.Wait()
	// for i := 0; i <= 1024; i++ {
	// 	results = append(results, ScanPort("udp", hostname, i))
	// }
	fmt.Println("initial scan com rotinas ", time.Since(start))
	return results
}

func WideScan(hostname string) []ScanResult {
	fmt.Println("wide scan start $$$$$$")
	start := time.Now()

	var results []ScanResult

	// for i := 0; i <= 49152; i++ {
	// 	go ScanPort("udp", hostname, i)

	// 	results = append(results)
	// }

	for i := 0; i <= 49152; i++ {
		fmt.Printf("olalala %d \n", i)
		result := ScanPort("tcp", hostname, i)
		if result.State == "Open" {
			results = append(results, result)
		}
	}
	fmt.Println("WideScan sem rotinas ", time.Since(start))

	return results
}
