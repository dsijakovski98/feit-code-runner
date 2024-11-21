package languages

type JavaScriptRunner struct {
	LanguageConfig
}

func (js JavaScriptRunner) GetConfig() LanguageConfig {
	return js.LanguageConfig
}

func (js JavaScriptRunner) ExtraCommands(containerId string) error {
	// Nothing extra needed
	return nil
}

func (js JavaScriptRunner) CommandChain(filePath string) []string {
	return []string{"node", filePath}
}
