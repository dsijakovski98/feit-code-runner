package languages

type PythonRunner struct {
	LanguageConfig
}

func (py PythonRunner) GetConfig() LanguageConfig {
	return py.LanguageConfig
}

func (py PythonRunner) ExtraCommands(filePath string, containerId string) (string, error) {
	// Nothing extra needed
	return "", nil
}

func (py PythonRunner) CommandChain(filePath string) []string {
	return []string{"python3", filePath}
}
