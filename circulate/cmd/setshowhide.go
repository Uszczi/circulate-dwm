package cmd

import (
	"circulate/circulate/core"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(setShowHide)
}

var setShowHide = &cobra.Command{
	Use: "toogle",
	Run: func(cmd *cobra.Command, args []string) {
		core.SendCommand("toogle")
	},
}
