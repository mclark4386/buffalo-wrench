package cmd

import (
	"fmt"

	"github.com/mclark4386/buffalo-wrench/wrench"
	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "current version of wrench",
	RunE: func(cmd *cobra.Command, args []string) error {
		fmt.Println("wrench", wrench.Version)
		return nil
	},
}

func init() {
	wrenchCmd.AddCommand(versionCmd)
}
