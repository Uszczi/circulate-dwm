package cmd

import (
	"circulate/tcp"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(setShowHide)
}

var setShowHide = &cobra.Command{
	Use: "toogle",
	Run: func(cmd *cobra.Command, args []string) {
		tcp.SendCommand("toogle")
	},
}
