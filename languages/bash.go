package languages

import "github.com/dsijakovski98/feit-code-runner/utils"

type BashRunner struct {
	LanguageConfig
}

func (sh BashRunner) GetConfig() LanguageConfig {
	return sh.LanguageConfig
}

func (sh BashRunner) ExtraRunCommands(filePath string, containerId string) (string, error) {
	// Give execution permission
	output, err := utils.ContainerExec(containerId, []string{"chmod", "+x", filePath})

	return output, err
}

func (sh BashRunner) RunCommand(filePath string) []string {
	return []string{filePath}
}

// Unused
func (sh BashRunner) ParserDir() string {
	return ""
}

// Unused
func (sh BashRunner) ParserCommand(filePath string) (string, []string) {
	return "", []string{}
}
