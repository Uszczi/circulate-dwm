package main

import (
	"circulate/circulate/core"
	"circulate/circulate/tcp"
	"os"
	"sync"
)

var wg sync.WaitGroup

func main() {
	wg.Add(2)
	arg := os.Args[1]
	if arg == "start" {
		go tcp.Run()
		go core.Main()

		wg.Wait()

	} else {
		tcp.SendCommand()
	}
}
