package cmd

import (
	"circulate/circulate/core"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(startCommand)
}

var startCommand = &cobra.Command{
	Use:   "start",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		core.Start()
		// [TODO] add short version
	},
}
