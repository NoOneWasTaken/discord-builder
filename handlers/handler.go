package handlers

import "fmt"

type LanguageHandler interface {
	InstallDependencies(projectName string)
}

func LanguageHandlerFactory(language string, packageManager string) (LanguageHandler, error) {
	switch language {
	case "csharp":
		return CSharpHandler{}, nil
	case "deno":
		return DenoHandler{}, nil
	case "bun":
		return BunHandler{}, nil
	default:
		return nil, fmt.Errorf("unsupported language: %s", language)
	}
}
