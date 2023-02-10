package cmd

import (
	"circulate/server"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(setLayout)
}

var setLayout = &cobra.Command{
	Use:       "set-layout",
	ValidArgs: []string{"columns", "rows", "floating", "previous", "next"},

	Run: func(cmd *cobra.Command, args []string) {
		server.SendCommand("set-layout " + args[0])
	},
}
