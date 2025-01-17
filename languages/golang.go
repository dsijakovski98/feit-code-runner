package languages

type GolangRunner struct {
	LanguageConfig
}

func (golang GolangRunner) GetConfig() LanguageConfig {
	return golang.LanguageConfig
}

func (golang GolangRunner) ExtraRunCommands(filePath string, containerId string) (string, error) {
	// Nothing extra needed
	return "", nil
}

func (golang GolangRunner) RunCommand(filePath string) []string {
	return []string{"go", "run", filePath}
}

func (golang GolangRunner) ParserDir() string {
	return "go"
}

func (golang GolangRunner) ParserCommand(filePath string) (string, []string) {
	return "go", []string{"run", "main.go", "--file", filePath}
}
