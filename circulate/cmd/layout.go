package cmd

import (
	"circulate/circulate/core"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(setLayout)
	setLayout.AddCommand(rows)
	setLayout.AddCommand(columns)
}

var setLayout = &cobra.Command{
	Use:   "set-layout",
	Short: "",
	Long:  "",
}

var rows = &cobra.Command{
	Use:   "rows",
	Short: "",
	Long:  "",

	Run: func(cmd *cobra.Command, args []string) {
		core.Start()
	},
}

var columns = &cobra.Command{
	Use:   "columns",
	Short: "",
	Long:  "",

	Run: func(cmd *cobra.Command, args []string) {
		core.Start()
	},
}
