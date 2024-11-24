package api

import (
	"fmt"
	"strings"
	"unicode"

	"github.com/dsijakovski98/feit-code-runner/languages"
)

const ERROR_PREFIX = "Error: "

func isValidLanguage(language string) bool {
	_, exists := languages.ProgrammingLanguages[language]

	return exists
}

func isErrorOutput(output string) bool {
	return strings.HasPrefix(output, ERROR_PREFIX)
}

type FilterErrorConfig struct {
	errOutput     string
	filePath      string
	taskName      string
	langExtension string
}

func filterErrorOutput(config FilterErrorConfig) string {
	var cleanErr = strings.TrimSpace(config.errOutput)
	cleanErr = strings.TrimPrefix(cleanErr, ERROR_PREFIX)

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
	return strings.TrimFunc(output, func(r rune) bool {
		return !unicode.IsGraphic(r)
	})
}
