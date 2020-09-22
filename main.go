package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	gomrepo "github.com/kemokemo/gomrepo/lib"
)

func main() {
	os.Exit(run(os.Args[1:]))
}

var (
	out    = os.Stdout
	outErr = os.Stderr
)

// pkginfo is the info of packages.
type pkginfo struct {
	id  string
	ver string
	lic string
}

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

	// The result of 'go list -m all' command enumerates the dependent modules.
	// Format: {identifier of the dependent module} {version}
	//         ex) cloud.google.com/go v0.26.0
	cmd := exec.Command("go", "list", "-m", "all")
	var cmdOut bytes.Buffer
	cmd.Stdout = &cmdOut
	err = cmd.Run()
	if err != nil {
		fmt.Fprintln(outErr, fmt.Sprintf("failed to run command: %v", err))
		return 1
	}

	modules := strings.Split(cmdOut.String(), "\n")
	pkgs := make(chan pkginfo)
	tokens := make(chan struct{}, 10)
	var counter int
	cl := gomrepo.NewGomClient()

	for _, module := range modules[1:] {
		fields := strings.Fields(module)
		if len(fields) < 2 {
			continue
		}
		counter++
		go func(id, ver string) {
			tokens <- struct{}{}
			lic, err := cl.GetLicense(id)
			if err != nil {
				fmt.Fprintln(outErr, "failed to get license info:", err)
			}
			<-tokens
			pkgs <- pkginfo{id, ver, lic}
		}(fields[0], fields[1])
	}

	for counter > 0 {
		pkg := <-pkgs
		fmt.Fprintln(out, fmt.Sprintf("%s %s %s", pkg.id, pkg.ver, pkg.lic))
		counter--
	}

	return 0
}
