package cmd

import (
	"circulate/server"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(setLayout)
	setLayout.AddCommand(rows)
	setLayout.AddCommand(columns)
	setLayout.AddCommand(floating)
	setLayout.AddCommand(previous)
	setLayout.AddCommand(next)
}

// [TODO] find a better way for this

var setLayout = &cobra.Command{
	Use:       "set-layout",
	ValidArgs: []string{"rows", "columns", "next", "previous", "kupa"},

	Run: func(cmd *cobra.Command, args []string) {
		server.SendCommand("set-layout " + args[0])
	},
}

var rows = &cobra.Command{
	Use: "rows",

	Run: func(cmd *cobra.Command, args []string) {
		server.SendCommand("set-layout rows")
	},
}

var columns = &cobra.Command{
	Use: "columns",

	Run: func(cmd *cobra.Command, args []string) {
		server.SendCommand("set-layout columns")
	},
}

var floating = &cobra.Command{
	Use: "floating",

	Run: func(cmd *cobra.Command, args []string) {
		server.SendCommand("set-layout floating")
	},
}

var previous = &cobra.Command{
	Use: "previous",

	Run: func(cmd *cobra.Command, args []string) {
		server.SendCommand("set-layout previous")
	},
}

var next = &cobra.Command{
	Use: "next",

	Run: func(cmd *cobra.Command, args []string) {
		server.SendCommand("set-layout next")
	},
}
