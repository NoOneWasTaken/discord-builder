package handlers

import (
	"discord-builder/constants"
	"fmt"
	"os/exec"
	"strings"
)

type BunHandler struct{}

func (d BunHandler) InstallDependencies(projectName string) {
	if _, err := exec.Command("bun", "--version").Output(); err != nil {
		if strings.Contains(err.Error(), "bun: command not found") || strings.Contains(err.Error(), "executable file not found in %PATH%") {
			fmt.Println(constants.ANSIBrightRed + "It looks like you don't have Bun Runtime installed. \nBun Runtime can be downloaded from:" + constants.ANSIBrightBlue + " https://bun.sh/docs/installation" + constants.ANSIReset)
			return
		}
	}

	cmd := exec.Command("bun", "i")
	cmd.Dir = projectName

	LogStdoutStream(cmd)

	fmt.Println(constants.ANSIBrightGreen + "\nDependencies have been successfully installed for Bun. Please configure the .env file and run the project. Happy coding! 🍱" + constants.ANSIReset)
}
