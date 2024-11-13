package api

var SUPPORTED_LANGUAGES = map[string]bool{
	"JavaScript": true,
	"TypeScript": true,
	"C":          true,
	"C++":        true,
	"Bash":       true,
	"Go":         true,
	"Python":     true,
	"Rust":       true,
	"PHP":        true,
}

func isValidLanguage(language string) bool {
	_, exists := SUPPORTED_LANGUAGES[language]

	return exists
}
