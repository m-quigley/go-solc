package mod

import (
	"os/exec"
	"path/filepath"
	"strings"
)

var Root string

func init() {
	stdout, _ := exec.Command("go", "env", "GOMOD").Output()

	// Trim spaces and check if the output ends with "/go.mod"
	modPath := strings.TrimSpace(string(stdout))
	var ok bool
	Root, ok = strings.CutSuffix(modPath, "/go.mod")
	if !ok {
		Root, ok = strings.CutSuffix(modPath, `\go.mod`) // Handle Windows path
	}

	if ok {
		// Normalize the path separators for compatibility across platforms
		Root = filepath.FromSlash(Root)
	} else {
		Root = ""
	}
}
