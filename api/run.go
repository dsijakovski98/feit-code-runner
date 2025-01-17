package api

import (
	"context"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/dsijakovski98/feit-code-runner/languages"
	"github.com/dsijakovski98/feit-code-runner/utils"
)

type RunRequest struct {
	Name     string `json:"name" binding:"required"`
	Code     string `json:"code" binding:"required"`
	Language string `json:"language" binding:"required"`
}

type RunOutput struct {
	Output string `json:"output"`
}

func runCode(req RunRequest, userId string) (string, error) {
	// Create temp file
	dir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("%s", "failed to find working directory: "+err.Error())
	}

	runner := languages.ProgrammingLanguages[req.Language]
	langConfig := runner.GetConfig()

	runName := strings.ReplaceAll(fmt.Sprintf("%s_%s_%s", userId, req.Name, time.Now().Format("Jan_02_15_04_05")), " ", "_")
	filename := fmt.Sprintf("%s.%s", runName, langConfig.Extension)
	filepath := fmt.Sprintf("%s/%s/%s", dir, utils.TMP_RUN_FILE, filename)

	f, err := os.Create(filepath)
	if err != nil {
		return "", fmt.Errorf("%s", "failed to create code file: "+err.Error())
	}
	defer f.Close()

	if err := os.WriteFile(filepath, []byte(req.Code), 0644); err != nil {
		return "", fmt.Errorf("%s", "failed to write to code file: "+err.Error())
	}

	ctx := context.Background()
	docker := utils.NewClient()

	config := container.Config{
		Image: langConfig.DockerImage,
		Cmd:   []string{"sh"}, // Use a shell to execute commands,
		Tty:   true,
	}

	// Create the container
	runContainer, err := docker.ContainerCreate(ctx, &config, nil, nil, nil, runName)
	if err != nil {
		return "", fmt.Errorf("%s", "failed to create container: "+err.Error())
	}

	// Start the container
	if err := docker.ContainerStart(ctx, runContainer.ID, container.StartOptions{}); err != nil {
		return "", fmt.Errorf("%s", "failed to start container: "+err.Error())
	}

	codeDir := "/user_code"
	if _, err := utils.ContainerExec(runContainer.ID, []string{"mkdir", "-p", codeDir}); err != nil {
		return "", fmt.Errorf("%s", "failed to create user_code dir: "+err.Error())
	}

	tgzPath, err := utils.CreateTgz(filepath)
	if err != nil {
		return "", fmt.Errorf("%s", "failed to create tarball: "+err.Error())
	}

	tgzReader, err := os.Open(tgzPath)
	if err != nil {
		return "", fmt.Errorf("%s", "failed to open tarball"+err.Error())
	}
	defer tgzReader.Close()

	if err := docker.CopyToContainer(ctx, runContainer.ID, codeDir, tgzReader, container.CopyToContainerOptions{
		AllowOverwriteDirWithFile: true,
	}); err != nil {
		return "", fmt.Errorf("%s", "failed to copy tarball to container: "+err.Error())
	}

	codeFilePath := fmt.Sprintf("%s/%s", codeDir, filename)

	// Cleanup in background
	cleanup := func() {
		defer os.Remove(filepath)
		defer os.Remove(tgzPath)

		if err := docker.ContainerRemove(ctx, runContainer.ID, container.RemoveOptions{Force: true}); err != nil {
			fmt.Printf("Error removing container: %v\n", err) // Log error but don't block return
		}
	}

	preOutput, err := runner.ExtraRunCommands(codeFilePath, runContainer.ID)
	if err != nil {
		return "", fmt.Errorf("%s", "failed to run extra commands: "+err.Error())
	}

	if utils.IsErrorOutput(preOutput) {
		filteredErr := filterErrorOutput(FilterErrorConfig{
			errOutput:     preOutput,
			filePath:      codeFilePath,
			taskName:      req.Name,
			langExtension: langConfig.Extension,
		})

		go cleanup()
		return filterUnicode(filteredErr), nil
	}

	output, err := utils.ContainerExec(runContainer.ID, runner.RunCommand(codeFilePath))
	if err != nil {
		return "", fmt.Errorf("%s", "failed to execute code: "+err.Error())
	}

	go cleanup()

	if utils.IsErrorOutput(output) {
		filteredErr := filterErrorOutput(FilterErrorConfig{
			errOutput:     output,
			filePath:      codeFilePath,
			taskName:      req.Name,
			langExtension: langConfig.Extension,
		})

		return filterUnicode(filteredErr), nil
	}

	return filterUnicode(output), nil
}
