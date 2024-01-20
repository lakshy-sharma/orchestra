/*
Copyright © 2024 Lakshy Sharma lakshy1106@protonmail.com
Author: Lakshy Sharma
Description:

	This file contains the code which has the structures used for controlling deployments.
*/
package db

import (
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
)

// This is the base structure used for making the deployments.
// Other deployment structures should embed this base structure as these properties are minimum requirements for any deployment.
type Deployment struct {
	ID            uuid.UUID `json:"id"`
	Name          string    `json:"name"`
	State         State     `json:"state"`
	RestartPolicy string    `json:"restart_policy"`
	StartTime     time.Time `json:"start_time"`
	FinishTime    time.Time `json:"finish_time"`
}

// This structure represents any event that occurs on a deployment.
// It is a base structure and shoulud be embedded as required by other event structures for deployments.
type DeploymentEvent struct {
	ID         uuid.UUID  `json:"id"`
	State      State      `json:"state"`
	Timestamp  time.Time  `json:"timestamp"`
	Deployment Deployment `json:"deployment"`
}

// The docker deployment is a structure which acts as the main structure for docker based deployments.
type DockerDeployment struct {
	Deployment
	Image        string            `json:"image"`
	Memory       int               `json:"memory"`
	Disk         int               `json:"disk"`
	ExposedPorts nat.PortSet       `json:"exposed_ports"`
	PortBindings map[string]string `json:"port_bindings"`
}

// The docker deployment event structure is used for creating events related to docker deployments.
type DockerDeploymentEvent struct {
	DeploymentEvent
}

// This function registers a new deployment in the database.
func (db *Database) RegisterDeployment() {}
