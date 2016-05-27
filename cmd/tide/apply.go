package main

import (
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
)

const applyDesc = `
This command applies a chart archive.

The apply argument must be either a relative
path to a chart directory or the name of a
chart in the current working directory.
`

var applyCmd = &cobra.Command{
	Use:   "apply [CHART]",
	Short: "apply a chart archive.",
	Long:  applyDesc,
	RunE:  runApply,
}

func runApply(cmd *cobra.Command, args []string) error {
	log.SetOutput(ioutil.Discard)
	setupInstallEnv(args)
	manifest, _ := readManifest(installArg)
	execute("apply", manifest)
	return nil
}

func setupApplyEnv(args []string) {
	if len(args) > 0 {
		installArg = args[0]
	} else {
		fatalf("This command needs at least one argument, the name of the chart.")
	}

}

func init() {
	applyCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose apply")
	RootCommand.AddCommand(applyCmd)
}
