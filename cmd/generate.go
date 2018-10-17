package cmd

import (
	"context"

	"github.com/gobuffalo/genny"
	"github.com/gobuffalo/genny/movinglater/gotools"
	"github.com/mclark4386/buffalo-wrench/genny/wrench"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var generateOptions = struct {
	*wrench.Options
	dryRun bool
}{
	Options: &wrench.Options{},
}

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "wrench",
	Short: "generates a new wrench",
	RunE: func(cmd *cobra.Command, args []string) error {
		r := genny.WetRunner(context.Background())

		if generateOptions.dryRun {
			r = genny.DryRunner(context.Background())
		}

		g, err := wrench.New(generateOptions.Options)
		if err != nil {
			return errors.WithStack(err)
		}
		r.With(g)

		g, err = gotools.GoFmt(r.Root)
		if err != nil {
			return errors.WithStack(err)
		}
		r.With(g)

		return r.Run()
	},
}

func init() {
	generateCmd.Flags().BoolVarP(&generateOptions.dryRun, "dry-run", "d", false, "run the generator without creating files or running commands")
	wrenchCmd.AddCommand(generateCmd)
}
