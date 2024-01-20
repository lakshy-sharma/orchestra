package db

import (
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
)

type State int

const (
	Pending State = iota
	Scheduled
	Running
	Completed
	Failed
)

type Task struct {
	ID            uuid.UUID         `json:"id"`
	Name          string            `json:"name"`
	State         State             `json:"state"`
	Image         string            `json:"image"`
	Memory        int               `json:"memory"`
	Disk          int               `json:"disk"`
	ExposedPorts  nat.PortSet       `json:"exposed_ports"`
	PortBindings  map[string]string `json:"port_bindings"`
	RestartPolicy string            `json:"restart_policy"`
	StartTime     time.Time         `json:"start_time"`
	FinishTime    time.Time         `json:"finish_time"`
}

type TaskEvent struct {
	ID        uuid.UUID `json:"id"`
	State     State     `json:"state"`
	Timestamp time.Time `json:"timestamp"`
	Task      Task      `json:"task"`
}

// This function creates a new task in the database.
func (db *Database) CreateTask() {}
