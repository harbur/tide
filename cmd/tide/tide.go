package main

import (
	"github.com/spf13/cobra"
)

// flagVerbose is a signal that the user wants additional output.
var flagVerbose bool

var globalUsage = `The Kubernetes package renderer

Tide is a standalone CLI tool to render Kubernetes packages with the Helm format.

It is a stateless tool, which uses the current directory to render the templates
and either outputs them for further usage (e.g. pipeline with kubectl or other
scripting methods) or installs them directly to Kubernetes.`

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
