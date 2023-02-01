package cmd

import (
	// usecase "circulate/circulate/usecase"
	"circulate/circulate/core"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(setLayout)
	setLayout.AddCommand(rows)
	setLayout.AddCommand(columns)
}

var setLayout = &cobra.Command{
	Use:       "set-layout",
	ValidArgs: []string{"rows", "columns"},
}

var rows = &cobra.Command{
	Use: "rows",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run rows")
		core.SendCommand("rows")
		// usecase.SetRowLayout()
	},
}

var columns = &cobra.Command{
	Use: "columns",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run columns")
		core.SendCommand("columns")
		// usecase.SetColumnLayout()
	},
}
