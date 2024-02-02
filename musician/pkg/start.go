/*
Copyright © 2024 Lakshy Sharma lakshy1106@protonmail.com
Author: Lakshy Sharma
Description:

	The main entrypoint for the code. This file contains the core controller loop for the musician.
*/
package pkg

import (
	"encoding/json"
	"musician/pkg/api"
	"musician/pkg/db"
	"time"

	"go.uber.org/zap"
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
		logger.Fatal("Failed to connect to with the database. Shutting Down Musician.")
	}

	stopMusician := false
	for !stopMusician {
		// Create a simple variable to store the incoming deployment data.
		// Fetch data from historian and unmarshal it into your variable.
		var DeploymentData db.Deployment
		deploymentString, err := historianConnection.Client.LPop("deployments").Result()
		if err != nil {
			logger.Error("Failed to fetch deployments queue from historian")
		}
		err = json.Unmarshal([]byte(deploymentString), &DeploymentData)

		// If unmarshaling succeeds, proceed with deployment else throw an error.
		if err == nil {
			// Use goroutines to spawn deployments.
			go api.CreateDeployment(DeploymentData)
		} else {
			logger.Sugar().Errorf("Failed while unmarshaling docker deployment data. Error: %s", err)
		}

		// Sleep for predetermined time and exit the program.
		time.Sleep(time.Duration(SleepDuration) * time.Second)
		stopMusician = true

	}
}
