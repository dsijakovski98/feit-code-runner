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
		Tty: true,
	})
	if err != nil {
		return "", err
	}

	defer exec.Close()

	// Read the response and convert it to a string
	var outputBuffer bytes.Buffer
	_, err = io.Copy(&outputBuffer, exec.Reader)
	if err != nil {
		panic(err)
	}

	// Convert byte slice to string
	output := outputBuffer.String()

	// Output stored in the buffer
	return output, nil

}
