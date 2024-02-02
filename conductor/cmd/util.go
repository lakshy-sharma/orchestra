/*
Copyright © 2024 Lakshy Sharma lakshy1106@protonmail.com
*/
package cmd

import (
	"conductor/pkg/db"
	"fmt"
	"log"
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func readConfigurations() db.ConductorConfig {
	viper.SetConfigName("conductor")
	viper.SetConfigType("toml")
	viper.AddConfigPath("/etc/orchestra")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error while handling the config file: %w", err))
	}
	musicianConfig := &db.ConductorConfig{}
	return *musicianConfig
}

func setupLogger() *zap.Logger {
	loggerConfig := zap.NewProductionConfig()
	loggerConfig.EncoderConfig.TimeKey = "timestamp"
	loggerConfig.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)

	logger, err := loggerConfig.Build()
	if err != nil {
		log.Fatal(err)
	}
	return logger
}
