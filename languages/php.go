package languages

type PhpRunner struct {
	LanguageConfig
}

func (php PhpRunner) GetConfig() LanguageConfig {
	return php.LanguageConfig
}

func (php PhpRunner) ExtraRunCommands(filePath string, containerId string) (string, error) {
	// Nothing extra needed
	return "", nil
}

func (php PhpRunner) RunCommand(filePath string) []string {
	return []string{"php", filePath}
}

// Unused
func (php PhpRunner) ParserDir() string {
	return ""
}

// Unused
func (php PhpRunner) ParserCommand(filePath string) (string, []string) {
	return "", []string{}
}
