package main // import "github.com/harbur/tide/cmd/tide"

import (
	"fmt"
	"io"
	"os"
	"os/exec"
)

func execute(command string, manifest string) {
	cmd := exec.Command("kubectl", command, "-f", "-")
	debug("hello there %s", manifest)
	stdin, err := cmd.StdinPipe()

	if err != nil {
		fmt.Println(err)
	}

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err = cmd.Start(); err != nil {
		fmt.Println("An error occured: ", err)
	}

	io.WriteString(stdin, manifest)
	stdin.Close()
	cmd.Wait()
}
