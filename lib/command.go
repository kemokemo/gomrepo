package gomrepo

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// GetModuleList returns the module list (name,version).
func GetModuleList(path string) ([]string, error) {
	err := os.Chdir(path)
	if err != nil {
		return nil, fmt.Errorf("failed to change directory to '%s': %v", path, err)
	}

	// The result of 'go list -m all' command enumerates the dependent modules.
	// Format: {identifier of the dependent module} {version}
	//         ex) cloud.google.com/go v0.26.0
	cmd := exec.Command("go", "list", "-m", "all")
	var cmdOut bytes.Buffer
	cmd.Stdout = &cmdOut
	err = cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to run command: %v", err)
	}

	return strings.Split(cmdOut.String(), "\n")[1:], nil
}
