package cmd

import (
	"fmt"

	"github.com/spf13/viper"
)

func ReadConfigurations() {
	viper.SetConfigName("conductor")
	viper.SetConfigType("toml")
	viper.AddConfigPath("/etc/orchestra")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("fatal error config file: %w", err))
	}
}
