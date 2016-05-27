package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
)

const installDesc = `
This command installs a chart archive.

The install argument must be either a relative
path to a chart directory or the name of a
chart in the current working directory.
`

var installCmd = &cobra.Command{
	Use:   "install [CHART]",
	Short: "install a chart archive.",
	Long:  installDesc,
	RunE:  runInstall,
}

func runInstall(cmd *cobra.Command, args []string) error {
	log.SetOutput(ioutil.Discard)
	setupInstallEnv(args)
	manifest, _ := readManifest(installArg)
	execute("create", manifest)
	return nil
}

func setupInstallEnv(args []string) {
	if len(args) > 0 {
		installArg = args[0]
	} else {
		fatalf("This command needs at least one argument, the name of the chart.")
	}
}

func fatalf(format string, args ...interface{}) {
	fmt.Printf("fatal: %s\n", fmt.Sprintf(format, args...))
	os.Exit(0)
}

func init() {
	installCmd.Flags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose install")
	RootCommand.AddCommand(installCmd)
}
