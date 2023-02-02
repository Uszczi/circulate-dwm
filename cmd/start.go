package cmd

import (
	"circulate/core"
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

	wg.Add(1)

	fmt.Println("Start Tcp Server")
	go core.RunTcpServer()
	fmt.Println("Start Windows Server")
	// go core.RunWindowsServer()

	wg.Wait()
}