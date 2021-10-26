package port

import (
	"net"
	"strconv"
	"time"
)

type ScanResult struct {
	Port  int
	State string
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

	var results []ScanResult

	for i := 0; i <= 1024; i++ {
		results = append(results, ScanPort("tcp", hostname, i))
	}

	for i := 0; i <= 1024; i++ {
		results = append(results, ScanPort("udp", hostname, i))
	}

	return results
}

func WideScan(hostname string) []ScanResult {
	var results []ScanResult

	for i := 0; i <= 49152; i++ {
		results = append(results, ScanPort("udp", hostname, i))
	}

	for i := 0; i <= 49152; i++ {
		results = append(results, ScanPort("tcp", hostname, i))
	}

	return results
}
