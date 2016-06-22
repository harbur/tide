package main // import "github.com/harbur/tide/cmd/tide"

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
)

const viewDesc = `
This command displays a chart archive.

The view argument must be either a relative
path to a chart directory or the name of a
chart in the current working directory.
`

var viewCmd = &cobra.Command{
	Use:   "view [CHART]",
	Short: "view a chart archive.",
	Long:  viewDesc,
	RunE:  runView,
}

func runView(cmd *cobra.Command, args []string) error {
	if len(args) == 0 {
		fatalf("This command needs at least one argument, the name of the chart.")
	}

	log.SetOutput(ioutil.Discard)

	for _, arg := range args {
		debug("Viewing %s", arg)
		manifest, _ := readManifest(arg)
		fmt.Printf("%s\n", manifest)
	}

	return nil
}

func init() {
	viewCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose view")
	viewCmd.Flags().StringVarP(&manifest_file, "file", "f", "", "view manifest file")
	viewCmd.Flags().StringVarP(&profile, "profile", "p", "", "activate profile to read alternative values input")
	RootCommand.AddCommand(viewCmd)
}
