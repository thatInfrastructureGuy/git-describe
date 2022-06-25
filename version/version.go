package version

import (
	"log"
	"os/exec"
	"strings"
)

// GetVersion gets version by executing "git describe --abbrev=7 --dirty"
// Needs git binary present.
// In absense of git binary, version is set to "UNKNOWN"
func GetVersion() string {
	git, err := exec.LookPath("git")
	if err != nil {
		log.Println("Error finding git:", err)
		return "UNKNOWN"
	}

	output, err := exec.Command(git, "describe", "--dirty", "--abbrev=7").Output()
	if err != nil {
		log.Println("Error executing git describe:", err)
		return "UNKNOWN"
	}

	return strings.Trim(string(output), " \n\t\r")
}
