package main

import (
	"discord-builder/constants"
	unzipModule "discord-builder/unzip"
	"discord-builder/validators"
	"fmt"
	"github.com/NoOneWasTaken/prompt"
	"os"
	"path"
	"strings"
)

type ProjectConfig struct {
	ProjectName    string
	Language       string
	PackageManager string
}

func InputCycle() (*ProjectConfig, error) {
	projectName, projectNameError := prompt.New().Ask("What would you like to name your project?").Input("my-bot")
	if projectNameError != nil {
		return nil, projectNameError
	}

	language, languageError := prompt.New().Ask("What language would you like to use?").Choose([]string{"CSharp", "Deno", "Bun"})
	if languageError != nil {
		return nil, languageError
	}

	projectConfig := &ProjectConfig{
		ProjectName: projectName,
		Language:    strings.ToLower(language),
	}

	if strings.ToLower(language) == "javascript" || strings.ToLower(language) == "typescript" {
		packageManager, packageManagerError := prompt.New().Ask("What package manager would you like to use?").Choose([]string{"npm", "yarn", "pnpm"})
		if packageManagerError != nil {
			return nil, packageManagerError
		}

		projectConfig.PackageManager = strings.ToLower(packageManager)
	}

	return projectConfig, nil
}

func main() {
	//InputCycle()
	projectConfig, inputErr := InputCycle()

	if inputErr != nil {
		if inputErr.Error() != "user quit prompt" {
			fmt.Println(constants.ANSIBrightRed + inputErr.Error() + constants.ANSIReset)
		}

		os.Exit(1)
	}

	if (projectConfig.Language == "javascript" || projectConfig.Language == "typescript") && projectConfig.PackageManager != "" {
		isValidName, message := validators.ValidateNpmPackageName(projectConfig.ProjectName)

		if !isValidName {
			fmt.Println("\n\n" + message + "\n\n")
			os.Exit(1)
		}
	}

	currentWorkDir, _ := os.Getwd()
	err := unzipModule.Extraction(path.Join(currentWorkDir, projectConfig.ProjectName), projectConfig.Language, projectConfig.PackageManager)

	if err != nil {
		fmt.Println(constants.ANSIBrightRed + err.Error() + constants.ANSIReset)
		os.Exit(1)
	}
}
