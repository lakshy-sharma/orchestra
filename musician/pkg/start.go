/*
Copyright © 2024 Lakshy Sharma lakshy1106@protonmail.com
Author: Lakshy Sharma
Description:

	The main entrypoint for the code. This file contains the core controller loop for the musician.
*/
package pkg

import (
	"fmt"
	"musician/pkg/api"
	"musician/pkg/db"
	"time"

	"github.com/docker/docker/api/types/container"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// This function is used for starting the gRPC servers which will handle the requests from clients.
// It only expects a logger and some predefined configurations.
func StartMusician(musicianConfigs db.MusicianConfig, logger *zap.Logger) {
	// Connect with conductors and get authentication details for database.
	// Setup a connection with the database.
	// Setup a main control loop which performs the following steps on each iteration.
	// 1. Check the database for any messages from the conductors.
	// 2. Perform the actions requested by the conductors.
	// 3. Collect metrics and upload them to the database.
	// 4. Monitor ongoing tasks and update their status on the database.

	// Establish a connection with database server.
	historianConnection, err := db.NewDatabase("127.0.0.1:6379", "", logger)
	if err != nil {
		logger.Fatal("Failed to connect to the database. Shutting Down Musician.")
	}

	stopMusician := false
	for !stopMusician {
		// Fetch data from historian.
		deploymentData, err := historianConnection.Client.LPop("docker_deployments").Result()
		if err != nil {
			logFields := zapcore.Field{Key: "error", String: err.Error()}
			logger.Error("Failed to fetch docker deployments queue from historian", logFields)
		}
		// Unmarshal deployment data into a deployment json.
		fmt.Print(deploymentData)
		// If application hasnt found any deployments then skip the creation steps.

		// Create a dummy deployment.
		dockerDeployment := db.DockerDeployment{ContainerName: "TestingDockerDeployment", ContainerConfig: &container.Config{Image: "alpine:latest", Cmd: []string{"echo", "Hello Lakshy"}, Tty: false}, Deployment: db.Deployment{Name: "Test", RestartPolicy: "always"}}
		// Spawn the deployment using a goroutine.
		go api.CreateDockerDeployment(dockerDeployment)

		time.Sleep(time.Duration(SleepDuration) * time.Second)
		stopMusician = true
	}
}
