package cmd

import (
	"circulate/core"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(send)
}

var send = &cobra.Command{
	Use: "ping",
	Run: func(cmd *cobra.Command, args []string) {
		core.SendCommand()
	},
}
