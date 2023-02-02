package cmd

import (
	"circulate/tcp"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(setLayout)
	setLayout.AddCommand(rows)
	setLayout.AddCommand(columns)
	setLayout.AddCommand(previous)
	setLayout.AddCommand(next)
}

var setLayout = &cobra.Command{
	Use:       "set-layout",
	ValidArgs: []string{"rows", "columns", "next", "previous"},
}

var rows = &cobra.Command{
	Use: "rows",

	Run: func(cmd *cobra.Command, args []string) {
		tcp.SendCommand("set-layout rows")
	},
}

var columns = &cobra.Command{
	Use: "columns",

	Run: func(cmd *cobra.Command, args []string) {
		tcp.SendCommand("set-layout columns")
	},
}

var previous = &cobra.Command{
	Use: "previous",

	Run: func(cmd *cobra.Command, args []string) {
		tcp.SendCommand("set-layout previous")
	},
}

var next = &cobra.Command{
	Use: "next",

	Run: func(cmd *cobra.Command, args []string) {
		tcp.SendCommand("set-layout next")
	},
}
