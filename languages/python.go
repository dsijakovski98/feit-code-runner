package languages

type PythonRunner struct {
	LanguageConfig
}

func (py PythonRunner) GetConfig() LanguageConfig {
	return py.LanguageConfig
}

func (py PythonRunner) ExtraRunCommands(filePath string, containerId string) (string, error) {
	// Nothing extra needed
	return "", nil
}

func (py PythonRunner) RunCommand(filePath string) []string {
	return []string{"python3", filePath}
}

// Unused
func (py PythonRunner) ParserDir() string {
	return ""
}

// Unused
func (py PythonRunner) ParserCommand(filePath string) (string, []string) {
	return "", []string{}
}
