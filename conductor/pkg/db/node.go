package db

type Node struct {
	Name            string `json:"node"`
	Ip              string `json:"ip"`
	Cores           int    `json:"cores"`
	Memory          int    `json:"memory"`
	MemoryAllocated int    `json:"memory_allocated"`
	Disk            int    `json:"disk"`
	DiskAllocated   int    `json:"disk_allocated"`
	Role            string `json:"role"`
	TaskCount       int    `json:"int"`
}

// This function creates a node.
func (db *Database) CreateNode() {}
