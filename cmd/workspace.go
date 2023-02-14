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
	Use:       "switch-to-workspace <workspace>",
	ValidArgs: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"},
	Run: func(cmd *cobra.Command, args []string) {
		server.SendCommand("switch-to-workspace " + args[0])
	},
}

var moveToLayout = &cobra.Command{
	Use:       "move-to-workspace",
	ValidArgs: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"},
	Run: func(cmd *cobra.Command, args []string) {
		server.SendCommand("move-to-workspace " + args[0])
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
