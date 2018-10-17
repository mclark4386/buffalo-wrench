package cmd

import (
	"github.com/spf13/cobra"
)

// wrenchCmd represents the buffalo-wrench command
var wrenchCmd = &cobra.Command{
	Use:   "buffalo-wrench",
	Short: "tools for working with wrench",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func init() {
	rootCmd.AddCommand(wrenchCmd)
}
