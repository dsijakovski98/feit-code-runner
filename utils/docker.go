package utils

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"strings"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
)

func NewClient() *client.Client {
	cli, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic(err)
	}

	return cli
}

func CheckClient() error {
	cli := NewClient()

	if _, err := cli.Info(context.Background()); err != nil {
		return err
	}

	return nil
}

func ContainerExec(containerId string, command []string) (string, error) {
	cli := NewClient()
	defer cli.Close()

	ctx := context.Background()

	fmt.Println("Running command: ", strings.Join(command, " "))

	// Create an exec configuration
	execConfig := container.ExecOptions{
		AttachStdout: true,
		AttachStderr: true,
		Cmd:          command,
	}

	// Create an exec session
	resp, err := cli.ContainerExecCreate(ctx, containerId, execConfig)
	if err != nil {
		return "", err
	}

	// Start the exec session
	exec, err := cli.ContainerExecAttach(ctx, resp.ID, container.ExecAttachOptions{
		Tty: false,
	})
	if err != nil {
		return "", err
	}

	defer exec.Close()

	// Read the response and convert it to a string
	var outputBuffer bytes.Buffer
	var errBuffer bytes.Buffer

	mux := io.MultiWriter(&outputBuffer, &errBuffer)

	_, err = io.Copy(mux, exec.Reader)
	if err != nil {
		panic(err)
	}

	inspect, err := cli.ContainerExecInspect(ctx, resp.ID)
	if err != nil {
		panic(err)
	}

	if inspect.ExitCode != 0 {
		errOutput := fmt.Sprintf("Error: %s", errBuffer.String())

		return errOutput, nil
	}

	// Convert byte slice to string
	return outputBuffer.String(), nil
}
