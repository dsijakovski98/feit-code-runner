package languages

import "github.com/dsijakovski98/feit-code-runner/utils"

type BashRunner struct {
	LanguageConfig
}

func (sh BashRunner) GetConfig() LanguageConfig {
	return sh.LanguageConfig
}

func (sh BashRunner) ExtraCommands(filePath string, containerId string) (string, error) {
	// Nothing extra needed
	output, err := utils.ContainerExec(containerId, []string{"chmod", "+x", filePath})

	return output, err
}

func (sh BashRunner) CommandChain(filePath string) []string {
	return []string{filePath}
}
