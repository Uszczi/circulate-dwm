package cmd

import (
	"circulate/circulate/tcp"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(send)
}

var send = &cobra.Command{
	Use:   "ping",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
        tcp.SendCommand()
	},
}
