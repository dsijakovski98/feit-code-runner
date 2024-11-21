package api

import (
	"github.com/dsijakovski98/feit-code-runner/languages"
)

func isValidLanguage(language string) bool {
	_, exists := languages.ProgrammingLanguages[language]

	return exists
}
