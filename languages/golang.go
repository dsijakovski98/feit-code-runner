package languages

type GolangRunner struct {
	LanguageConfig
}

func (golang GolangRunner) GetConfig() LanguageConfig {
	return golang.LanguageConfig
}

func (golang GolangRunner) ExtraCommands(filePath string, containerId string) (string, error) {
	// Nothing extra needed
	return "", nil
}

func (golang GolangRunner) CommandChain(filePath string) []string {
	return []string{"go", "run", filePath}
}
