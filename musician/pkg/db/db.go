/*
Copyright © 2024 Lakshy Sharma lakshy1106@protonmail.com
Author: Lakshy Sharma
Description:

	This file contains the code which is used for handling connections with database and initiating the database.
*/
package db

import (
	"context"
	"errors"

	"github.com/go-redis/redis"
)

// Pre-defined list of states.
type State int

const (
	Pending State = iota // 0
	Scheduled
	Running
	Completed
	Failed
)

// A struct to connect with the database.
type Database struct {
	Client *redis.Client
}

var (
	ErrNil = errors.New("no matching record found in redis database")
	Ctx    = context.TODO()
)

type MusicianConfig struct {
}

// This function creates a new database connection and returns it to you.
func NewDatabase(address string) (*Database, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	})
	if err := client.Ping().Err(); err != nil {
		return nil, err
	}
	return &Database{
		Client: client,
	}, nil
}
