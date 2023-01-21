package cmd

import (
	"circulate/circulate/core"
	"circulate/circulate/tcp"
	"fmt"
	"sync"

	"github.com/spf13/cobra"
)

var wg sync.WaitGroup

func init() {
	rootCmd.AddCommand(startCommand)
}

var startCommand = &cobra.Command{
	Use:   "start",
	Run: func(cmd *cobra.Command, args []string) {
		start()
	},
}

func start(){
    fmt.Println("Start circulate")

	wg.Add(2)

    go tcp.Run()
    go core.Main()

    wg.Wait()
}



