package cmd

import (
	"circulate/server"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(moveToLayout)
	rootCmd.AddCommand(switchToWorkspace)
	rootCmd.AddCommand(debugWorkspace)
	rootCmd.AddCommand(clearWorkspace)
}

var switchToWorkspace = &cobra.Command{
	Use:     "switch-to-workspace <workspace>",
	Aliases: []string{"stw"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			return
		}
		workspace := args[0]

		server.SendCommand("switch-to-workspace" + " " + workspace)
	},
}

var moveToLayout = &cobra.Command{
	Use:     "move-to-workspace",
	Aliases: []string{"mtw"},

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			return
		}
		workspace := args[0]
		server.SendCommand("move-to-workspace" + " " + workspace)
	},
}

var debugWorkspace = &cobra.Command{
	Use: "debug-workspace",

	Run: func(cmd *cobra.Command, args []string) {
		server.SendCommand("debug-workspace")
	},
}

var clearWorkspace = &cobra.Command{
	Use: "clear-workspace",

	Run: func(cmd *cobra.Command, args []string) {
		server.SendCommand("clear-workspace")
	},
}
