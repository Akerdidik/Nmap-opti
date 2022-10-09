package nmapper

import "fmt"

func Helper() {
	fmt.Println("Usage: go run main.go <commands>\nCommands: --help = to see what commands you can use;\n --target to specify target\n --port to specify port(0,port)")
}

func Unavailable(target string) {
	fmt.Printf("Target %s is currently sleep or not exist!\n", target)
}
