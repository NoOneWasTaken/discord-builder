package unzip

import (
	"archive/zip"
	"bytes"
	"discord-builder/constants"
	"discord-builder/handlers"
	"embed"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
)

//go:embed templates/*
var templates embed.FS

func Extraction(destination string, language string, packageManager string) error {
	if _, err := os.Stat(destination); err == nil {
		files, err := os.ReadDir(destination)

		if err != nil {
			return fmt.Errorf("error encountered while validating the destination folder:\n\n%v", err)
		}

		if len(files) > 0 {
			return fmt.Errorf("it seems like the destination folder is not empty. please make sure that the folder is empty before proceeding")
		}
	}

	if err := os.MkdirAll(destination, os.ModePerm); err != nil {
		return fmt.Errorf("error encountered while creating the destination folder:\n\n%v", err)
	}

	bundle, err := templates.ReadFile("templates/" + language + ".zip")
	if err != nil {
		return fmt.Errorf("error encountered while reading the bundled files:\n\n%v", err)
	}

	zipReader, err := zip.NewReader(bytes.NewReader(bundle), int64(len(bundle)))
	if err != nil {
		return fmt.Errorf("error encountered while extracting necessary files:\n\n%v", err)
	}

	for _, zipContent := range zipReader.File {
		extractionPath := path.Join(destination, zipContent.Name)

		if zipContent.FileInfo().IsDir() {
			if err := os.MkdirAll(extractionPath, os.ModePerm); err != nil {
				return err
			}
			continue
		}

		fileReader, err := zipContent.Open()
		if err != nil {
			return fmt.Errorf("error encountered while reading the files:\n\n%v", err)
		}

		destFile, err := os.Create(extractionPath)
		if err != nil {
			return fmt.Errorf("error encountered while creating the destination file:\n\n%v", err)
		}

		if _, err := io.Copy(destFile, fileReader); err != nil {
			return fmt.Errorf("error encountered while copying the files to destination source:\n\n%v", err)
		}

		if err := fileReader.Close(); err != nil {
			return fmt.Errorf("error encountered while closing the file reader:\n\n%v", err)
		}

		if err := destFile.Close(); err != nil {
			return fmt.Errorf("error encountered while closing the destination file:\n\n%v", err)
		}
	}

	fmt.Println(constants.ANSIBrightGreen + "\n\nFiles have been successfully extracted to the destination folder." + constants.ANSIReset)

	langHandler, err := handlers.LanguageHandlerFactory(language, packageManager)
	if err != nil {
		return fmt.Errorf(constants.ANSIBrightRed+"error encountered while auto-installing the packages for %s:\n\n%v"+constants.ANSIReset, language, err)
	}

	defer langHandler.InstallDependencies(filepath.Base(destination))

	return nil
}
