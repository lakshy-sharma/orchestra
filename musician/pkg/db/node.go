/*
Copyright © 2024 Lakshy Sharma lakshy1106@protonmail.com
Author: Lakshy Sharma
Description:

	This file contains the code which has the structures used for controlling nodes.
*/
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
