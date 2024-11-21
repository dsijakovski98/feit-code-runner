package api

import (
	"context"
	"fmt"
	"os"

	"github.com/docker/docker/api/types/container"
	"github.com/dsijakovski98/feit-code-runner/languages"
	"github.com/dsijakovski98/feit-code-runner/utils"
)

func runCode(req RunRequest, userId string) (string, error) {
	// Create temp file
	dir, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("failed to find working directory: " + err.Error())
	}

	runner := languages.ProgrammingLanguages[req.Language]
	langConfig := runner.GetConfig()

	runName := fmt.Sprintf("%s_%s", userId, req.Name)
	filename := fmt.Sprintf("%s.%s", runName, langConfig.Extension)
	filepath := fmt.Sprintf("%s/_tmp/%s", dir, filename)

	f, err := os.Create(filepath)
	if err != nil {
		return "", fmt.Errorf("failed to create code file: " + err.Error())
	}
	defer f.Close()

	if err := os.WriteFile(filepath, []byte(req.Code), 0644); err != nil {
		return "", fmt.Errorf("failed to write to code file: " + err.Error())
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
		return "", fmt.Errorf("failed to create container: " + err.Error())
	}

	// Start the container
	if err := docker.ContainerStart(ctx, runContainer.ID, container.StartOptions{}); err != nil {
		return "", fmt.Errorf("failed to start container: " + err.Error())
	}

	codeDir := "/user_code"
	if _, err := utils.ContainerExec(runContainer.ID, []string{"mkdir", "-p", codeDir}); err != nil {
		return "", fmt.Errorf("failed to create user_code dir: " + err.Error())
	}

	tgzPath, err := utils.CreateTgz(filepath)
	if err != nil {
		return "", fmt.Errorf("failed to create tarball: " + err.Error())
	}

	tgzReader, err := os.Open(tgzPath)
	if err != nil {
		return "", fmt.Errorf("failed to open tarball" + err.Error())
	}
	defer tgzReader.Close()

	if err := docker.CopyToContainer(ctx, runContainer.ID, codeDir, tgzReader, container.CopyToContainerOptions{
		AllowOverwriteDirWithFile: true,
	}); err != nil {
		return "", fmt.Errorf("failed to copy tarball to container: " + err.Error())
	}

	if err := runner.ExtraCommands(runContainer.ID); err != nil {
		return "", fmt.Errorf("failed to run extra commands: " + err.Error())
	}

	codeFilePath := fmt.Sprintf("%s/%s", codeDir, filename)
	output, err := utils.ContainerExec(runContainer.ID, runner.CommandChain(codeFilePath))
	if err != nil {
		return "", fmt.Errorf("failed to execute code: " + err.Error())
	}

	// Cleanup in background using goroutine
	go func() {
		defer os.Remove(filepath)
		defer os.Remove(tgzPath)

		if err := docker.ContainerStop(ctx, runContainer.ID, container.StopOptions{}); err != nil {
			fmt.Printf("Error stopping container: %v\n", err) // Log error but don't block return
		}

		if err := docker.ContainerRemove(ctx, runContainer.ID, container.RemoveOptions{}); err != nil {
			fmt.Printf("Error removing container: %v\n", err) // Log error but don't block return
		}
	}()

	return output, nil
}
