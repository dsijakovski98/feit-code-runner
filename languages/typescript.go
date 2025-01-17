package languages

type TypeScriptRunner struct {
	LanguageConfig
}

func (ts TypeScriptRunner) GetConfig() LanguageConfig {
	return ts.LanguageConfig
}

func (js TypeScriptRunner) ExtraRunCommands(filePath string, containerId string) (string, error) {
	// Nothing extra needed
	return "", nil
}

func (ts TypeScriptRunner) RunCommand(filePath string) []string {
	return []string{"bun", "run", filePath}
}

func (ts TypeScriptRunner) ParserDir() string {
	return "js-ts"
}

func (ts TypeScriptRunner) ParserCommand(filePath string) (string, []string) {
	return "bun", []string{"start", "--", "--file", filePath, "--lang", "ts"}
}
