package languages

type JavaScriptRunner struct {
	LanguageConfig
}

func (js JavaScriptRunner) GetConfig() LanguageConfig {
	return js.LanguageConfig
}

func (js JavaScriptRunner) ExtraRunCommands(filePath string, containerId string) (string, error) {
	// Nothing extra needed
	return "", nil
}

func (js JavaScriptRunner) RunCommand(filePath string) []string {
	return []string{"node", filePath}
}

func (js JavaScriptRunner) ParserDir() string {
	return "js-ts"
}

func (js JavaScriptRunner) ParserCommand(filePath string) (string, []string) {
	return "bun", []string{"start", "--", "--file", filePath, "--lang", "js"}
}
