package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of Ravictl",
	Long:  `All software has versions. This is Ravictl's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print("v0.1\n")
	},
}
