/*
Copyright © 2024 Lakshy Sharma lakshy1106@protonmail.com
Author: Lakshy Sharma
Description:

	This file contains the code which is responsible for hanlding incoming API commands related to Docker.
*/
package api

import (
	"musician/pkg/db"
	"musician/pkg/pods"

	"github.com/docker/docker/api/types"
)

func createDockerDeployment(dockerData db.DockerData) error {
	dockerInstance, err := pods.NewDockerClient()
	if err != nil {
		panic(err)
	}
	// Pull the image.
	dockerInstance.PullDockerImage(dockerData.ContainerConfig.Image, types.ImagePullOptions{})
	// Create the container using configurations.
	creationResponse, err := dockerInstance.CreateDockerContainer(dockerData.ContainerConfig, nil, nil, nil, dockerData.ContainerName)
	if err != nil {
		return err
	}
	// Start the container.
	dockerInstance.StartDockerContainer(creationResponse.ID)
	// Fetch the container logs.
	dockerInstance.GetDockerContainerLogs(creationResponse.ID)
	// Stop the container.
	dockerInstance.StopDockerContainer(creationResponse.ID, false)
	return nil
}

// Get logs pertaining to a deployment.
func GetDockerDeploymentLogs() {

}

// Stop the deployment.
func StopDockerDeployment() {

}
