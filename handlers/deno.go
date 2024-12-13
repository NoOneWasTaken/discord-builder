package handlers

import (
	"discord-builder/constants"
	"fmt"
	"os/exec"
	"strings"
)

type DenoHandler struct{}

func (d DenoHandler) InstallDependencies(projectName string) {
	// Deno does not require any dependencies to be installed as it is a runtime environment and not a framework.
	// This function is a placeholder for future implementations.

	if _, err := exec.Command("deno", "--version").Output(); err != nil {
		if strings.Contains(err.Error(), "deno: command not found") {
			fmt.Println(constants.ANSIBrightRed + "It looks like you don't have Deno 2.0 Runtime installed. \n.Deno 2.0 Runtime can be downloaded from:" + constants.ANSIBrightBlue + " https://docs.deno.com/runtime/#install-deno" + constants.ANSIReset)
			return
		}
	}

	fmt.Printf(constants.ANSIBrightGreen+"As Deno does not require any dependencies to be installed, you can directly run the project using the command:"+constants.ANSIBrightWhite+"\n\n\tcd %s/\n\tdeno task dev\n\n"+constants.ANSIBrightGreen+"Please configure the .env file and run the project. Happy coding! 🦕"+constants.ANSIReset, projectName)
}
