package cmd

import (
	"circulate/server"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(focusRoot)
}

var focusRoot = &cobra.Command{
	Use:       "focus",
	ValidArgs: []string{"next"},

	Run: func(cmd *cobra.Command, args []string) {
		server.SendCommand("focus " + args[0])
	},
}
