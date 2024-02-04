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
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
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

// This function creates a new database connection and returns it to you.
func NewDatabase(address string, password string, logger *zap.Logger) (*Database, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0,
	})
	if err := client.Ping().Err(); err != nil {
		logger.Error("Faced an error while pinging database", zapcore.Field{Key: "error", Type: zapcore.StringType, String: err.Error()})
		return nil, err
	}
	return &Database{
		Client: client,
	}, nil
}
