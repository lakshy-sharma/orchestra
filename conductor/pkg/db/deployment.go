/*
Copyright © 2024 Lakshy Sharma lakshy1106@protonmail.com
Author: Lakshy Sharma
Description:

	This file contains the code which has the structures used for controlling deployments.
*/
package db

import (
	"time"

	"github.com/docker/docker/api/types/container"
	"github.com/google/uuid"
)

type Deployment struct {
	ID                   uuid.UUID  `json:"id"`
	Name                 string     `json:"name"`
	State                State      `json:"state"`
	RestartPolicy        string     `json:"restart_policy"`
	StartTime            time.Time  `json:"start_time"`
	FinishTime           time.Time  `json:"finish_time"`
	DeploymentsSupported []string   `json:"deployments_supported"`
	DockerData           DockerData `json:"docker_data"`
	KvmData              KvmData    `json:"kvm_data"`
}

type DeploymentEvent struct {
	ID         uuid.UUID  `json:"id"`
	State      State      `json:"state"`
	Timestamp  time.Time  `json:"timestamp"`
	Deployment Deployment `json:"deployment"`
}

type DockerData struct {
	ContainerName   string            `json:"container_name"`
	Memory          int               `json:"memory"`
	Disk            int               `json:"disk"`
	PortBindings    map[string]string `json:"port_bindings"`
	ContainerConfig *container.Config `json:"container_config"`
}

type KvmData struct {
}
