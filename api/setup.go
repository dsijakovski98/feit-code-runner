package api

import (
	"context"
	"fmt"
	"io"
	"sync/atomic"

	"github.com/docker/docker/api/types/image"
	"github.com/dsijakovski98/feit-code-runner/languages"
	"github.com/dsijakovski98/feit-code-runner/utils"
)

var ready atomic.Bool

func IsReady() bool {
	return ready.Load()
}

func EnsureImagesInstalled(ctx context.Context) error {
	cli := utils.NewClient()

	// Derive unique images from ProgrammingLanguages
	seen := make(map[string]bool)
	for _, runner := range languages.ProgrammingLanguages {
		img := runner.GetConfig().DockerImage
		if seen[img] {
			continue // e.g. gcc appears twice (C and C++)
		}

		seen[img] = true

		_, _, err := cli.ImageInspectWithRaw(ctx, img)
		if err == nil {
			fmt.Println("Image", img, "is already present", img)
			continue
		}

		fmt.Println("Pulling image", img, "...")

		out, err := cli.ImagePull(ctx, img, image.PullOptions{})
		if err != nil {
			return fmt.Errorf("failed to pull %s: %w", img, err)
		}

		io.Copy(io.Discard, out)
		out.Close()

		fmt.Println("Image", img, "is ready")
	}

	ready.Store(true)

	return nil
}
