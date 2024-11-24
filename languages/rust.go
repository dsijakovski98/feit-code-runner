package languages

import (
	"github.com/dsijakovski98/feit-code-runner/utils"
)

type RustRunner struct {
	LanguageConfig
}

func (rs RustRunner) GetConfig() LanguageConfig {
	return rs.LanguageConfig
}

func (rs RustRunner) ExtraCommands(filePath string, containerId string) (string, error) {
	outPath := utils.GetOutPath(filePath)

	// Compile rust file
	output, err := utils.ContainerExec(containerId, []string{"rustc", "-o", outPath, filePath})

	return output, err
}

func (rs RustRunner) CommandChain(filePath string) []string {
	outPath := utils.GetOutPath(filePath)

	return []string{outPath}
}
