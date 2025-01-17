package api

import (
	"bytes"
	"fmt"
	"os/exec"

	"github.com/dsijakovski98/feit-code-runner/languages"
)

type CleanupRequest struct {
	TaskName string `json:"taskName" binding:"required"`
	Code     string `json:"code" binding:"required"`
	Language string `json:"language" binding:"required"`
}

type CleanupResponse struct {
	CleanCode string `json:"cleanCode" binding:"required"`
}

func cleanupCode(req CleanupRequest, filePath string) error {
	lang := languages.ProgrammingLanguages[req.Language]

	testParser := lang.ParserDir()
	cmdName, args := lang.ParserCommand(filePath)

	cmd := exec.Command(cmdName, args...)
	cmd.Dir = fmt.Sprintf("parsers/%s", testParser)

	var errOut bytes.Buffer
	cmd.Stderr = &errOut

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("cleanup failed: %w, stderr: %s", err, errOut.String())
	}

	return nil
}
