package cmd

import (
	"circulate/server"
	"circulate/usecase"
	"io"
	"log"
	"os"
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
	logFile, err := os.OpenFile("log.txt", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}
	mw := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(mw)

	usecase.Setup()

	tasks := []func(){
		server.RunWindowsServer,
		server.RunTcpServer,
	}

	wg.Add(len(tasks))
	for _, task := range tasks {
		go task()
	}
	log.Println("Circulate started.")

	wg.Wait()
}
