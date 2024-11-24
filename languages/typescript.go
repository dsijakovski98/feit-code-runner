package languages

type TypeScriptRunner struct {
	LanguageConfig
}

func (ts TypeScriptRunner) GetConfig() LanguageConfig {
	return ts.LanguageConfig
}

func (js TypeScriptRunner) ExtraCommands(filePath string, containerId string) (string, error) {
	// Nothing extra needed
	return "", nil
}

func (ts TypeScriptRunner) CommandChain(filePath string) []string {
	return []string{"bun", "run", filePath}
}
