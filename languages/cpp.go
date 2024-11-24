package languages

import (
	"github.com/dsijakovski98/feit-code-runner/utils"
)

type CppRunner struct {
	LanguageConfig
}

func (cpp CppRunner) GetConfig() LanguageConfig {
	return cpp.LanguageConfig
}

func (cpp CppRunner) ExtraCommands(filePath string, containerId string) (string, error) {
	outPath := utils.GetOutPath(filePath)

	// Compile cpp file
	output, err := utils.ContainerExec(containerId, []string{"g++", "-o", outPath, filePath})

	return output, err
}

func (cpp CppRunner) CommandChain(filePath string) []string {
	outPath := utils.GetOutPath(filePath)

	return []string{outPath}
}
