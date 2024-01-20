/*
Copyright © 2024 Lakshy Sharma lakshy1106@protonmail.com
Author: Lakshy Sharma
Description:

	The main entrypoint for the code. This file contains the core controller loop for the musician.
*/
package pkg

import (
	"musician/pkg/db"

	"go.uber.org/zap"
)

// This function is used for starting the gRPC servers which will handle the requests from clients.
// It only expects a logger and some predefined configurations.
func StartMusician(musicianConfigs db.MusicianConfig, logger *zap.Logger) {
	// Setup a connection with the database.
	// Setup the servers to communicate in gRPC.
	// Setup a main control loop which performs the following steps on each iteration.
	// 1. Check the database for any messages from the conductors.
	// 2. Perform the actions requested by the conductors.
	// 3. Collect metrics and upload them to the database.
	// 4. Monitor ongoing tasks and update their status on the database.

	//stopMusician := false
	//for !stopMusician {}
}
