package api

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/dsijakovski98/feit-code-runner/languages"
	"github.com/dsijakovski98/feit-code-runner/utils"
)

func isValidLanguage(language string) bool {
	_, exists := languages.ProgrammingLanguages[language]

	return exists
}

func supportsTests(language string) bool {
	langConfig := languages.ProgrammingLanguages[language]

	return langConfig.GetConfig().TestsSupport
}

type FilterErrorConfig struct {
	errOutput     string
	filePath      string
	taskName      string
	langExtension string
}

func filterErrorOutput(config FilterErrorConfig) string {
	var cleanErr = strings.TrimSpace(config.errOutput)
	cleanErr = strings.TrimPrefix(cleanErr, utils.ERROR_PREFIX)

	chunks := strings.Split(cleanErr, "\n")

	// Filter out OS sensitive data
	chunkTail := chunks[len(chunks)-1]
	if strings.Contains(strings.ToLower(chunkTail), strings.ToLower("Linux")) {
		cleanErr = strings.Join(chunks[:len(chunks)-1], "\n")
	}

	cleanFilePath := fmt.Sprintf("%s.%s", config.taskName, config.langExtension)
	cleanErr = strings.Replace(cleanErr, config.filePath, cleanFilePath, -1)

	return cleanErr
}

func filterUnicode(output string) string {
	cleanOutput := strings.Map(func(r rune) rune {
		// Allow letters, digits, symbols, and spaces
		if unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsSpace(r) || unicode.IsPunct(r) || unicode.IsSymbol(r) {
			return r
		}

		// Discard other characters
		return -1
	}, output)

	// Remove consecutive spaces, while keeping new-line characters
	// https://stackoverflow.com/a/71646989
	return strings.ReplaceAll(strings.Join(strings.FieldsFunc(cleanOutput, func(r rune) bool {
		if r == '\n' {
			return false
		}
		return unicode.IsSpace(r)
	}), " "), " \n", "\n")
}
