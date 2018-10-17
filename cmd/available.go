package cmd

import (
	"github.com/gobuffalo/buffalo-plugins/plugins/plugcmds"
	"github.com/mclark4386/buffalo-wrench/wrench"
)

var Available = plugcmds.NewAvailable()

func init() {
	Available.Add("root", wrenchCmd)
	Available.Listen(wrench.Listen)
	Available.Add("generate", generateCmd)
	Available.Mount(rootCmd)
}
