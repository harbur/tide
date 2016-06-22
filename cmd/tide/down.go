package main // import "github.com/harbur/tide/cmd/tide"

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
	Use:   "down [CHART]",
	Short: "Delete a chart archive from Kubernetes.",
	Long:  uninstallDesc,
	RunE:  runUninstall,
}

func runUninstall(cmd *cobra.Command, args []string) error {
	log.SetOutput(ioutil.Discard)

	if len(args) == 0 {
		fatalf("This command needs at least one argument, the name of the chart.")
	}

	for _, arg := range args {
		manifest, _ := readManifest(arg)
		execute("delete", manifest)
	}
	return nil
}

func init() {
	uninstallCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose render")
	RootCommand.AddCommand(uninstallCmd)
}
