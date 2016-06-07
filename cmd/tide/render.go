package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
)

const renderDesc = `
This command renders a chart archive.

The render argument must be either a relative
path to a chart directory or the name of a
chart in the current working directory.
`

var renderCmd = &cobra.Command{
	Use:   "view [CHART]",
	Short: "view a chart archive.",
	Long:  renderDesc,
	RunE:  runRender,
}

func runRender(cmd *cobra.Command, args []string) error {
	log.SetOutput(ioutil.Discard)
	setupRenderEnv(args)
	manifest, _ := readManifest(installArg)
	fmt.Printf("%s\n", manifest)
	return nil
}

func setupRenderEnv(args []string) {
	if len(args) > 0 {
		installArg = args[0]
	} else {
		fatalf("This command needs at least one argument, the name of the chart.")
	}
}

func init() {
	renderCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose render")
	renderCmd.Flags().StringVarP(&manifest_file, "file", "f", "", "view manifest file")
	RootCommand.AddCommand(renderCmd)
}
