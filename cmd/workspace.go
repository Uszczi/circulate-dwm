package cmd

import (
	"circulate/core"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(moveToLayout)
	rootCmd.AddCommand(switchToWorkspace)
	rootCmd.AddCommand(debugWorkspace)
}

var switchToWorkspace = &cobra.Command{
	Use:     "switch-to-workspace <workspace>",
	Aliases: []string{"stw"},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			return
		}
		workspace := args[0]

		fmt.Println("switch-to-workspace" + " " + workspace)
		core.SendCommand("switch-to-workspace" + " " + workspace)
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
		core.SendCommand("move-to-workspace" + " " + workspace)
	},
}

var debugWorkspace = &cobra.Command{
	Use: "debug-workspace",

	Run: func(cmd *cobra.Command, args []string) {
		core.SendCommand("debug-workspace")
	},
}
