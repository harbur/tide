package main // import "github.com/harbur/tide/cmd/tide"

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

// install flags & args
var (
	// verbose enables verbose output
	verbose bool
	// watch enables watch mode
	watch bool
	// delete enables automatic deletion on exit
	deletion bool
	// manifest file
	manifest_file string
	// input file
	input_file string
)

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

func fatalf(format string, args ...interface{}) {
	fmt.Printf("fatal: %s\n", fmt.Sprintf(format, args...))
	os.Exit(0)
}

func init() {
	RootCommand.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose output")
}

func main() {
	RootCommand.Execute()
}
