package config // Defines the package name as "config"

import (
	"flag"      // Used for command-line flag parsing
	"log"       // Provides logging capabilities
	"os"        // Allows interaction with the operating system (e.g., reading environment variables, checking file existence)

	"github.com/ilyakaznacheev/cleanenv" // Library to read and parse YAML configuration files into structs
)

// HTTPServer configuration struct
type HTTPServer struct {
	Addr string `yaml:"addr"` // Maps "addr" field in YAML to this struct field
}

// Config struct for the application
type Config struct {
	Env         string     `yaml:"env" env:"ENV" env-required:"true"`         // Reads "env" from YAML and allows overriding with ENV variable
	StoragePath string     `yaml:"storage_path" env-required:"true"`          // Reads "storage_path" from YAML, required
	HTTPServer  HTTPServer `yaml:"http_server"`                               // Reads "http_server" block from YAML
}

// MustLoad loads configuration from a file or environment variables
func MustLoad() *Config {
	var configPath string // Stores the path to the configuration file

	// First, check if CONFIG_PATH is set in the environment variables
	configPath = os.Getenv("CONFIG_PATH")

	// If CONFIG_PATH is not set, check for a command-line flag
	if configPath == "" {
		flagConfig := flag.String("config", "", "path to the configuration file") // Define a flag for setting the config path
		flag.Parse() // Parse command-line flags

		configPath = *flagConfig // Assign the value of the "config" flag to configPath
	}

	// If config path is still empty after checking both environment variable and command-line flag, exit with an error
	if configPath == "" {
		log.Fatal("config path is not set!") // Log the error and terminate the program
	}

	// Check if the config file exists at the specified path
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("config file does not exist: %s", configPath) // Log error and exit if file is missing
	}

	// Create an instance of the Config struct
	var cfg Config

	// Load configuration from the YAML file into the Config struct
	err := cleanenv.ReadConfig(configPath, &cfg)
	if err != nil {
		log.Fatalf("cannot read config file: %s", err.Error()) // Log error and exit if loading fails
	}

	return &cfg // Return the loaded configuration
}
