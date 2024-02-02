/*
Copyright © 2024 Lakshy Sharma lakshy1106@protonmail.com
Author: Lakshy Sharma
Description:

	This file contains the code responsible for handling the incoming API requests.
	It handles all requests and routes them further to dedicated functions.
*/
package api

import "musician/pkg/db"

func CreateDeployment(deploymentData db.Deployment) {

	// For each supported deployment in the provided data.
	// We must create a new deployment.
	for _, supportedDeployment := range deploymentData.DeploymentsSupported {
		if supportedDeployment == "docker" {
			createDockerDeployment(deploymentData.DockerData)
		}
	}
}
