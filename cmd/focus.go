package cmd

import (
	"circulate/server"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(changeFocus)
}

var changeFocus = &cobra.Command{
	Use:       "focus",
	ValidArgs: []string{"next", "previous"},
	Run: func(cmd *cobra.Command, args []string) {
		server.SendCommand("focus " + args[0])
	},
}
