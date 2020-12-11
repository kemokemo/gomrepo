package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	gomrepo "github.com/kemokemo/gomrepo/lib"
)

// flags
var (
	format string
)

func init() {
	flag.StringVar(&format, "format", "markdown", "format to output")
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
	if len(flag.Args()) > 0 {
		dirPath = filepath.Clean((flag.Args())[0])
	} else {
		dirPath = filepath.Clean(".")
	}

	modules, err := gomrepo.GetModuleList(dirPath)
	if err != nil {
		fmt.Fprintln(outErr, "failed to get module list: ", err)
		return 1
	}

	cl := gomrepo.NewGomClient()
	err = cl.GetLicenseList(out, modules, getFormatter(format))
	if err != nil {
		fmt.Fprintln(outErr, "failed to get license list: ", err)
		return 1
	}

	return 0
}

func getFormatter(f string) gomrepo.Formatter {
	f = strings.ToLower(f)
	switch f {
	case "markdown", "md":
		return gomrepo.MD
	case "html", "htm":
		return gomrepo.HTML
	case "asciidoc", "ascii":
		return gomrepo.ASCIIDoc
	default:
		return gomrepo.MD
	}
}
