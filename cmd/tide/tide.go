package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var stdout = os.Stdout
var helmHome string

// flagVerbose is a signal that the user wants additional output.
var flagVerbose bool

var globalUsage = `The Kubernetes package renderer

Tide is a standalone CLI tool to render Kubernetes packages with the Helm format.

It is a stateless tool, which uses the current directory to render the templates
and either outputs them for further usage (e.g. pipeline with kubectl or other
scripting methods) or installs them directly to Kubernetes.

Common actions from this point on include:

- tide search: search for charts
- tide fetch: download a chart to your local directory to view
- tide install: upload the chart to Kubernetes
- tide list: list releases of charts
`

// RootCommand is the top-level command for Helm.
var RootCommand = &cobra.Command{
	Use:   "tide",
	Short: "The Kubernetes package renderer.",
	Long:  globalUsage,
}

func init() {
	RootCommand.PersistentFlags().BoolVarP(&flagVerbose, "verbose", "v", false, "enable verbose output")
}

func main() {
	RootCommand.Execute()
}

func checkArgsLength(expectedNum, actualNum int, requiredArgs ...string) error {
	if actualNum != expectedNum {
		arg := "arguments"
		if expectedNum == 1 {
			arg = "argument"
		}
		return fmt.Errorf("This command needs %v %s: %s", expectedNum, arg, strings.Join(requiredArgs, ", "))
	}
	return nil
}
