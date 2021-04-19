package main

import (
	"context"

	"github.com/docker/docker/client"
)

func getImageArch(c *client.Client, image string) ([]string, error) {
	var architectures []string
	var architecture string

	ctx := context.Background()

	resp, err := c.DistributionInspect(ctx, image, "")
	for _, platform := range resp.Platforms {
		architecture = platform.Architecture
		if platform.Variant != "" {
			architecture = architecture + platform.Variant
		}
		architectures = append(architectures, architecture)
	}

	return architectures, err
}

func architectureIsSupported(targetArch string, allArchitectures []string) bool {
	for _, a := range allArchitectures {
		if string(a) == targetArch {
			return true
		}
	}

	return false
}