package main

import (
	"circulate/circulate/tcp"
	"os"
)

func main() {
	arg := os.Args[1]
	if arg == "start" {
		tcp.Run()
	} else {
		tcp.SendCommand()
	}
	// cmd.Execute()
}
