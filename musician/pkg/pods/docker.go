/*
Copyright © 2024 Lakshy Sharma lakshy1106@protonmail.com
Author: Lakshy Sharma
Description:

	This file implements an interface for common docker operations.
	The functions contained here are used for providing access to the underlying docker API.
*/
package pods

import (
	"context"
	"io"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
	v1 "github.com/opencontainers/image-spec/specs-go/v1"
)

// This structure contains the bare minimum details required to perform docker operations.
// It can be initialized using NewDockerClient() function defined below.
type Docker struct {
	Context context.Context
	Client  *client.Client
}

// A function to initialize a new Docker client.
// Once initialized you can provide the generated client to other functions which implement the methods.
func NewDockerClient() (*Docker, error) {
	dockerContext := context.Background()
	dockerClient, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		defer dockerClient.Close()
		// Return a docker instance which can be used by other functions.
		return &Docker{
			Context: dockerContext,
			Client:  dockerClient,
		}, nil
	}

	// Return a docker instance which can be used by other functions.
	return &Docker{
		Context: dockerContext,
		Client:  dockerClient,
	}, err
}

// A function to pull docker images.
// Inputs are image name and pull options as defined by Docker SDK.
func (docker *Docker) PullDockerImage(imageName string, pullOptions types.ImagePullOptions) error {
	// Read the image from container registry.
	imageReader, err := docker.Client.ImagePull(docker.Context, imageName, pullOptions)
	if err != nil {
		return err
	}
	defer imageReader.Close()

	// The ImagePull is asynchronous this means we need to completely read the reader before we proceed.
	// You can use os.Stdout to track the pull.
	io.Copy(io.Discard, imageReader)
	return nil
}

// A function to create Docker container.
// The response provides you with the ID and Warnings that were faced during initializing process.
func (docker *Docker) CreateDockerContainer(containerConfig *container.Config, hostConfig *container.HostConfig, networkConfig *network.NetworkingConfig, platformConfig *v1.Platform, containerName string) (container.CreateResponse, error) {
	// Issue the call to create a new docker container.
	response, err := docker.Client.ContainerCreate(docker.Context, containerConfig, hostConfig, networkConfig, platformConfig, containerName)
	return response, err
}

// A function that takes a Container ID as an input and starts a docker container.
func (docker *Docker) StartDockerContainer(containerID string) error {
	if err := docker.Client.ContainerStart(docker.Context, containerID, types.ContainerStartOptions{}); err != nil {
		return err
	}
	return nil
}

// A function to fetch docker container logs and print them to stdout.
func (docker *Docker) GetDockerContainerLogs(containerID string) error {
	out, err := docker.Client.ContainerLogs(docker.Context, containerID, types.ContainerLogsOptions{ShowStdout: true})
	if err != nil {
		return err
	}
	// Simple print the logs to the stdout or stderr.
	stdcopy.StdCopy(os.Stdout, os.Stderr, out)

	// Return no error if not required.
	return nil
}

// A function to inspect a running docker container.
func (docker *Docker) GetDockerContainerStatus(containerID string) (types.ContainerJSON, error) {
	output, err := docker.Client.ContainerInspect(docker.Context, containerID)
	return output, err
}

// A function to stop running docker containers.
func (docker *Docker) StopDockerContainer(containerID string, forceStop bool) error {
	var noWaitTimeout = 10
	// If you want to force quit set timeout to 0.
	if forceStop {
		noWaitTimeout = 0
	}
	if err := docker.Client.ContainerStop(docker.Context, containerID, container.StopOptions{Timeout: &noWaitTimeout}); err != nil {
		return err
	}
	return nil
}
