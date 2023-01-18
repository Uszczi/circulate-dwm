package cmd

import (
	"github.com/asaskevich/EventBus"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(send)
}

var send = &cobra.Command{
	Use:   "send",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		server := EventBus.NewServer(":2010", "/_server_bus_", EventBus.New())
		server.Start()
		server.EventBus().Publish("main:calculator", 4, 6)
		server.Stop()
	},
}
