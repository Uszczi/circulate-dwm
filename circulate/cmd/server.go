package cmd

import (
	"circulate/circulate/core"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(server)
}

var server = &cobra.Command{
	Use:   "start",
	Short: "",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		core.Main()
		// [TODO] add short version
	},
}
