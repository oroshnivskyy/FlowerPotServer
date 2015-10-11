package config

import (
	"flag"
)

type CommandLineConfiguration struct {
	ConfigFilePath string
}

// GetCommandLineConfiguration parses and loads command-line configuration
func GetCommandLineConfiguration() (config *CommandLineConfiguration, err error) {
	config = new(CommandLineConfiguration)

	flag.StringVar(&(config.ConfigFilePath), "c", "./cfg/config.toml", "Path to configuration")
	flag.Parse()

	return
}