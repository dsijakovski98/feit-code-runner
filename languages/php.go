package languages

type PhpRunner struct {
	LanguageConfig
}

func (php PhpRunner) GetConfig() LanguageConfig {
	return php.LanguageConfig
}

func (php PhpRunner) ExtraCommands(filePath string, containerId string) (string, error) {
	// Nothing extra needed
	return "", nil
}

func (php PhpRunner) CommandChain(filePath string) []string {
	return []string{"php", filePath}
}
