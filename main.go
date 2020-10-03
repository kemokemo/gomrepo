package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	gomrepo "github.com/kemokemo/gomrepo/lib"
)

// flags
var (
	format string
)

func init() {
	flag.StringVar(&format, "format", "", "format to output")
	flag.Parse()
}

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

	modules, err := gomrepo.GetModuleList(dirPath)
	if err != nil {
		fmt.Fprintln(outErr, "failed to get module list: ", err)
		return 1
	}

	cl := gomrepo.NewGomClient()
	licenses, err := cl.GetLicenseList(modules, gomrepo.MD)
	if err != nil {
		fmt.Fprintln(outErr, "failed to get license list: ", err)
		return 1
	}

	fmt.Fprintln(out, licenses)

	return 0
}
