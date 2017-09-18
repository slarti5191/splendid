package configuration

import (
	"sync"
)

// ConfigLoader - Orchestrates the configuration loading.
// Config - The plain config type.
// ConfigFile - Loads configuration data from a file.
// ConfigFlags - Loads configuration data from a flag.

// Singleton pattern ensures a single config across concurrent threads.
var instance *Config
var once sync.Once

// GetConfig is concurrency safe loading and retrieving of the config data.
func GetConfig() *Config {
	once.Do(func() {
		instance = loadConfig()
	})
	return instance
}

// loadConfig fetches the various sources of configuration data, and returns the fully prepared config.
func loadConfig() *Config {
	config := new(Config)

	// Fetch defaults
	config.mergeConfig(getConfigDefaults())

	// Fetch file
	// Load INI
	configPath := configFlagsGetConfigPath()
	mergeConfigFile(config, configPath)
	//config.mergeConfig(getFileConfig(configPath))

	// Fetch flags
	//config.mergeConfig(getConfigFlags())
	mergeConfigFlags(config)

	return config
}
