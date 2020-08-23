package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

func main() {
	os.Exit(run(os.Args[1:]))
}

var (
	out    = os.Stdout
	outErr = os.Stderr
)

func run(args []string) int {
	var dirPath string
	if len(args) > 0 {
		dirPath = filepath.Clean(args[0])
	} else {
		dirPath = filepath.Clean(".")
	}

	err := os.Chdir(dirPath)
	if err != nil {
		fmt.Fprintln(outErr, fmt.Sprintf("failed to change directory to '%s': %v", dirPath, err))
		return 1
	}

	cmd := exec.Command("go", "list", "-m", "all")
	var cmdOut bytes.Buffer
	cmd.Stdout = &cmdOut
	err = cmd.Run()
	if err != nil {
		fmt.Fprintln(outErr, fmt.Sprintf("failed to run command: %v", err))
	}

	fmt.Fprintf(out, "result of 'go list -m all' command:\n%v\n", cmdOut.String())
	return 0
}
