package nmapper

import (
	"fmt"
	"net"
	"time"
)

func Scanner(host string, ports []string) {
	for _, port := range ports {
		timeout := time.Second / 2
		conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
		if err != nil {
			fmt.Println("Troubles with connection!\n", err)
		}
		if conn != nil {
			defer conn.Close()
			fmt.Println("Port Open: ", net.JoinHostPort(host, port))
		}
	}
}

func TcpGather(target string, ports []string) map[string]string {
	results := make(map[string]string)
	for _, port := range ports {
		addr := net.JoinHostPort(target, port)
		conn, err := net.DialTimeout("tcp", addr, time.Second)
		if err != nil {
			results[port] = "failed"
		} else {
			if conn != nil {
				results[port] = "open"
				_ = conn.Close()
			} else {
				results[port] = "close"
			}
		}
	}
	return results
}
