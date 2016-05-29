package main

import (
	"github.com/howeyc/fsnotify"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
)

const applyDesc = `
This command applies a chart archive.

The apply argument must be either a relative
path to a chart directory or the name of a
chart in the current working directory.
`

var applyCmd = &cobra.Command{
	Use:   "up [CHART]",
	Short: "Apply a chart archive to Kubernetes.",
	Long:  applyDesc,
	RunE:  runApply,
}

func runApply(cmd *cobra.Command, args []string) error {
	if deletion && !watch {
		fatalf("--delete can only be used in combination with --watch")
		return nil
	}
	setupApplyEnv(args)
	if watch {
		// log.SetOutput(ioutil.Discard)
		manifest, _ := readManifest(installArg)
		execute("apply", manifest)
		watchForApply()
	} else {
		log.SetOutput(ioutil.Discard)
		manifest, _ := readManifest(installArg)
		execute("apply", manifest)
	}
	return nil
}

func watchForApply() {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	done := make(chan bool)

	// Process events
	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				log.Println("event:", ev)
				log.SetOutput(ioutil.Discard)
				manifest, _ := readManifest(installArg)
				execute("apply", manifest)
			case err := <-watcher.Error:
				log.Println("error:", err)
			}
		}
	}()
	log.Printf("dir %s", installArg)
	err = watcher.Watch(installArg)
	if err != nil {
		log.Fatal(err)
	}

	err = watcher.Watch(installArg + "/templates")
	if err != nil {
		log.Fatal(err)
	}

	// Exit on Ctrl-C
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, os.Interrupt)
	go func() {
		<-sigs

		// Delete Chart if --delete flag is true
		if deletion {
			manifest, _ := readManifest(installArg)
			execute("delete", manifest)
		}
		done <- true
	}()

	// Hang so program doesn't exit
	<-done

	/* ... do stuff ... */
	watcher.Close()
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
	applyCmd.Flags().BoolVarP(&watch, "watch", "w", false, "watch directory for changes")
	applyCmd.Flags().BoolVarP(&deletion, "delete", "d", false, "Automatically delete the chart when it exits")
	RootCommand.AddCommand(applyCmd)
}
