package main

import (
	"file-path-validator/internal/processor"
	"file-path-validator/internal/validator"
	"flag"
	"fmt"
	"os"
	"strings"
)

const (
	ExpectedArgumentsCount = 2
)

func main() {
	if len(os.Args) < ExpectedArgumentsCount {
		fmt.Println("Usage: file_path_validator")
		fmt.Println("    --naming-convention  snake_case|PascalCase")
		fmt.Println("    --path-to-validate  <file_path>")
		flag.Usage()

		return
	}

	namingConvention, pathToValidate := parseFlags()

	validateFlags(namingConvention, pathToValidate)

	validateFunction := getValidateFunction(namingConvention)

	validatePath(pathToValidate, validateFunction)
}

func parseFlags() (string, string) {
	namingConventionFlag := flag.String("naming-convention", "", "Naming convention to use (snake_case|PascalCase)")
	pathToValidateFlag := flag.String("path-to-validate", "", "A file path to validate")

	flag.Parse()

	return *namingConventionFlag, *pathToValidateFlag
}

func validateFlags(namingConvention, pathToValidate string) {
	if namingConvention != "snake_case" && namingConvention != "PascalCase" {
		fmt.Println("Error: --naming-convention must be either 'snake_case' or 'PascalCase'")
		flag.Usage()
		os.Exit(1)
	}

	if pathToValidate == "" {
		fmt.Println("Error: --path-to-validate is required")
		flag.Usage()
		os.Exit(1)
	}
}

func getValidateFunction(namingConvention string) func(string) bool {
	if namingConvention == "snake_case" {
		return validator.IsSnakeCase
	}

	return validator.IsPascalCase
}

func validatePath(path string, validateFunction func(string) bool) {
	if processor.IsHiddenFile(path) {
		return
	}

	fileName := processor.GetFileNameWithoutExtension(path)

	if !validateFunction(fileName) {
		fmt.Printf("Invalid file name: %s\nPath: %s\n", fileName, path)
		os.Exit(1)
	}

	folders := processor.GetFoldersFromPath(path)
	isAnyFolderInvalid := false

	for _, folderName := range folders {
		// NOTE: remove hidden prefix from folder name to validate it correctly
		folderName = strings.TrimPrefix(folderName, ".")

		if !validateFunction(folderName) {
			fmt.Printf("Invalid folder name: %s\nPath: %s\n", folderName, path)

			isAnyFolderInvalid = true
		}
	}

	if isAnyFolderInvalid {
		os.Exit(1)
	}
}
