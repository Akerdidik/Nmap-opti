package main

import (
	"fmt"
	"nmapper"
	"os/exec"
	"strconv"
)

func main() {
	nmapper.Argcheck()

	for i := 0; i <= nmapper.Port; i++ {
		t := strconv.Itoa(i)
		nmapper.Ports = append(nmapper.Ports, t)
	}
	resp := exec.Command("ping", "-c", "1", nmapper.Target)
	err := resp.Run()
	if err != nil {
		nmapper.Unavailable(nmapper.Target)
		return
	}
	nmapper.Scanner(nmapper.Target, nmapper.Ports)
	results := nmapper.TcpGather(nmapper.Target, nmapper.Ports)
	fmt.Println(results)
}
