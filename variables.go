package nmapper

import "os"

var (
	Args   []string = os.Args[1:]
	Target string
	Flag   bool = false
	Port   int  = 65535
	Ports  []string
)
