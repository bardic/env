package main

import (
	"context"
	"dagger/env/internal/dagger"
	"fmt"
	"strings"
)

type Env struct{}

func (m *Env) Load(ctx context.Context, f *dagger.File, c *dagger.Container) (*dagger.Container, error) {
	envs, err := f.Contents(ctx)

	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	envPair := strings.Split(envs, "\n")

	for _, v := range envPair {
		envVals := strings.SplitN(v, "=", 2)

		isSecret := strings.Contains(envVals[0], "S|")
		if isSecret {
			fmt.Println("Secret found")
			secretName := strings.TrimPrefix(envVals[0], "S|")
			c.WithSecretVariable(secretName, dag.SetSecret(secretName, envVals[1]))

		} else {
			fmt.Println("Env found")
			c = c.WithEnvVariable(envVals[0], envVals[1])
		}
	}

	return c, nil
}
