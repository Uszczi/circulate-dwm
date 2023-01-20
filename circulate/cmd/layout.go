package cmd

import (
	usecase "circulate/circulate/usecase"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(setLayout)
	setLayout.AddCommand(rows)
	setLayout.AddCommand(columns)
}

var setLayout = &cobra.Command{
	Use:   "set-layout",
}

var rows = &cobra.Command{
	Use:   "rows",

	Run: func(cmd *cobra.Command, args []string) {
		usecase.SetRowLayout()
	},
}

var columns = &cobra.Command{
	Use:   "columns",

	Run: func(cmd *cobra.Command, args []string) {
        usecase.SetColumnLayout()
	},
}
