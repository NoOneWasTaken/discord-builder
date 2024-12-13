package handlers

import (
	"discord-builder/constants"
	"fmt"
	"os/exec"
	"strings"
)

type CSharpHandler struct{}

func (c CSharpHandler) InstallDependencies(projectName string) {
	fmt.Print("Proceeding with auto-installation of dependencies for C#...\n\n")
	if _, err := exec.Command("dotnet", "--version").Output(); err != nil {
		if strings.Contains(err.Error(), "executable file not found") {
			fmt.Println(constants.ANSIBrightRed + "It looks like you don't have .net core framework installed. Please install it and manually install the packages.\n.NET Core Framework can be downloaded from:" + constants.ANSIBrightBlue + " https://dotnet.microsoft.com/download" + constants.ANSIReset)
			return
		}
	}

	cmd := exec.Command("dotnet", "restore")
	cmd.Dir = projectName

	LogStdoutStream(cmd)

	fmt.Println(constants.ANSIBrightGreen + "\nDependencies have been successfully installed for C# project. Please configure the .env file and run the project. Happy coding! 🚀" + constants.ANSIReset)
}
