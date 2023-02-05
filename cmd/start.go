package cmd

import (
	"circulate/tcp"
	"fmt"
	"sync"

	"github.com/spf13/cobra"
)

var wg sync.WaitGroup

func init() {
	rootCmd.AddCommand(startCommand)
}

var startCommand = &cobra.Command{
	Use: "start",
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}

func start() {
	fmt.Println("Start circulate")

	tasks := []func(){
		// core.RunWindowsServer,
		tcp.RunTcpServer,
	}

	wg.Add(len(tasks))

	for _, task := range tasks {
		go task()
	}

	wg.Wait()
}
