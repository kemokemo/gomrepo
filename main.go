/*
Copyright © 2020 kemokemo <t2wonderland@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/
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
	format     string
	versionOut bool
	helpOut    bool
)

func init() {
	flag.StringVar(&format, "f", "markdown", "format to output")
	flag.StringVar(&format, "format", "markdown", "format to output")
	flag.BoolVar(&versionOut, "v", false, "print the version number")
	flag.BoolVar(&versionOut, "version", false, "print the version number")
	flag.BoolVar(&helpOut, "h", false, "print the help")
	flag.BoolVar(&helpOut, "help", false, "print the help")

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
	if versionOut {
		fmt.Fprintf(out, "%s.%s\n", Version, Revision)
		return 0
	}
	if helpOut {
		flag.PrintDefaults()
		fmt.Fprintf(out, "\nAvailable format list:\n %v\n\nThey are not case-sensitive.\n", availableFormats)
		return 0
	}

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

var availableFormats = []string{"markdown", "md", "html", "html", "asciidoc", "ascii", "textile"}

func getFormatter(f string) gomrepo.Formatter {
	f = strings.ToLower(f)
	switch f {
	case "markdown", "md":
		return gomrepo.MD
	case "html", "htm":
		return gomrepo.HTML
	case "asciidoc", "ascii":
		return gomrepo.ASCIIDoc
	case "textile":
		return gomrepo.Textile
	default:
		return gomrepo.MD
	}
}
