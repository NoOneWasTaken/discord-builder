package validators

import (
	"discord-builder/constants"
	"regexp"
	"strings"
	"unicode"
)

func ValidateNpmPackageName(projectName string) (bool, string) {
	var warnings []string
	var errors []string

	if len(projectName) == 0 || projectName == "" {
		errors = append(errors, "project name cannot be empty")
	}

	for _, char := range projectName {
		if unicode.IsDigit(char) {
			errors = append(errors, "project name should not contain any numbers")
			break
		}
	}

	if regexp.MustCompile("^\\.").MatchString(projectName) {
		errors = append(errors, "project name cannot start with a period")
	}

	if regexp.MustCompile("^_").MatchString(projectName) {
		errors = append(errors, "project name cannot start with an underscore")
	}

	if regexp.MustCompile("^\\s+|\\s+$").MatchString(projectName) {
		errors = append(errors, "project name cannot contain any leading/trailing whitespace")
	}

	for _, blacklisted := range constants.NodeJSBlacklistedKeywords {
		if strings.ToLower(projectName) == blacklisted {
			errors = append(errors, "project name is a blacklisted name ("+blacklisted+")")
		}
	}

	// Warnings

	for _, module := range constants.NodeJSDefaultModules {
		if strings.ToLower(projectName) == module {
			warnings = append(warnings, "project name is a default module name ("+module+")")
		}
	}

	if len(projectName) > 214 {
		warnings = append(warnings, "project name is longer than 214 characters")
	}

	if strings.ToLower(projectName) != projectName {
		warnings = append(warnings, "project name should not be in mixed case")
	}

	if regexp.MustCompile("[~\\\\'!()* ]").MatchString(projectName) {
		warnings = append(warnings, "project name contains special characters (~\\'!()*)")
	}

	if len(errors) > 0 || len(warnings) > 0 {
		infoString := constants.ANSIBrightRed + "❌ Failed to create project with name " + constants.ANSIBrightWhite + projectName + constants.ANSIBrightRed + ". This is due to the npm naming restrictions.\nPlease checkout the following for more details:" + constants.ANSIBrightBlue + " https://docs.npmjs.com/cli/configuring-npm/package-json" + constants.ANSIReset
		errorString := constants.ANSIRed + "\n\nerrors:\n\t" + constants.ANSIReset
		warningString := constants.ANSIYellow + "\n\nwarnings:\n\t" + constants.ANSIReset

		for i, err := range errors {
			if i == len(errors)-1 {
				errorString += err
			} else {
				errorString += err + ", "
			}
		}

		for i, warn := range warnings {
			if i == len(warnings)-1 {
				warningString += warn
			} else {
				warningString += warn + ", "
			}
		}

		return false, infoString + errorString + warningString
	}

	return true, ""
}
