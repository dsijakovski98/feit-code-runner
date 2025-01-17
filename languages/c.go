package languages

import (
	"github.com/dsijakovski98/feit-code-runner/utils"
)

type CRunner struct {
	LanguageConfig
}

func (c CRunner) GetConfig() LanguageConfig {
	return c.LanguageConfig
}

func (c CRunner) ExtraRunCommands(filePath string, containerId string) (string, error) {
	outPath := utils.GetOutPath(filePath)

	// Compile c file
	output, err := utils.ContainerExec(containerId, []string{"gcc", "-o", outPath, filePath})

	return output, err
}

func (c CRunner) RunCommand(filePath string) []string {
	outPath := utils.GetOutPath(filePath)

	return []string{outPath}
}

// Unused
func (c CRunner) ParserDir() string {
	return ""
}

// Unused
func (c CRunner) ParserCommand(filePath string) (string, []string) {
	return "", []string{}
}
