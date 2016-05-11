package main

import (
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
)

const uninstallDesc = `
This command uninstalls a chart archive.

The uninstall argument must be either a relative
path to a chart directory or the name of a
chart in the current working directory.
`

var uninstallCmd = &cobra.Command{
	Use:   "uninstall [CHART]",
	Short: "uninstall a chart archive.",
	Long:  uninstallDesc,
	RunE:  runUninstall,
}

func runUninstall(cmd *cobra.Command, args []string) error {
	log.SetOutput(ioutil.Discard)
	setupInstallEnv(args)
	manifest, _ := readManifest(installArg)
	execute("delete", manifest)
	return nil
}

func setupUninstallEnv(args []string) {
	if len(args) > 0 {
		installArg = args[0]
	} else {
		fatalf("This command needs at least one argument, the name of the chart.")
	}

}

func init() {
	uninstallCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose render")
	RootCommand.AddCommand(uninstallCmd)
}
