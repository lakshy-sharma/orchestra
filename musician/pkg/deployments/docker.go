/*
Copyright © 2024 Lakshy Sharma lakshy1106@protonmail.com
Author: Lakshy Sharma
Description:

	This file contains the code which is responsible for hanlding incoming API commands related to Docker.
*/
package deployments

import (
	"fmt"
	"musician/pkg/pods"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
)

func CreateDockerDeployment(imageName string, containerConfig *container.Config) error {
	dockerInstance, err := pods.NewDockerClient()
	if err != nil {
		panic(err)
	}

	// Pull the image.
	fmt.Println("Pulling the docker image.")
	dockerInstance.PullDockerImage(imageName, types.ImagePullOptions{})

	// Define a container and then create it.
	fmt.Println("Starting a new container.")
	//containerDefinition := &container.Config{
	//  Image: imageName,
	//	Cmd:   []string{"echo", "hello world"},
	//	Tty:   false,
	//}
	containerName := ""
	creationResponse, err := dockerInstance.CreateDockerContainer(containerConfig, nil, nil, nil, containerName)
	if err != nil {
		return err
	}

	// Start the container.
	fmt.Println("Starting the container.")
	dockerInstance.StartDockerContainer(creationResponse.ID)

	// Fetch the container logs.
	fmt.Println("Fetching container logs.")
	dockerInstance.GetDockerContainerLogs(creationResponse.ID)

	// Stop the container.
	fmt.Println("Stopping the container.")
	dockerInstance.StopDockerContainer(creationResponse.ID, false)
	return nil
}

// Get logs pertaining to a deployment.
func GetDockerDeploymentLogs() {

}

// Stop the deployment.
func StopDockerDeployment() {

}
