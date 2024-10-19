package uploader

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/viper"
)

type Server struct {
	Host     string `json:"host"`
	Username string `json:"username"`
	Password string `json:"password"`
	Port     int    `json:"port"`
}

type Config struct {
	Source      string   `json:"source"`
	Destination string   `json:"destination"`
	Server      []Server `json:"server"`
}

// LoadConfig reads the config file and return a Config object
func LoadConfig() (Config, error) {
	viper.SetConfigName(CONFIG_FILENAME)
	viper.SetConfigType(CONFIG_FILETYPE)
	viper.AddConfigPath(CONFIG_FILEPATH)

	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
		return Config{}, err
	}

	var config Config
	if err := viper.Unmarshal(&config); err != nil {
		fmt.Printf("Unable to decode into struct, %v", err)
		return Config{}, err
	}

	return config, nil
}

// Default configuration values
var defaultConfig = Config{
	Server: []Server{
		{
			Host:     "localhost",
			Username: "admin",
			Password: "secret",
			Port:     8080,
		},
	},
	Source:      "source",
	Destination: "/dest/",
}

// Function to create the default config file
func CreateDefaultConfig(filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	// Encode the default config to JSON
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ") // For pretty printing
	if err := encoder.Encode(defaultConfig); err != nil {
		return err
	}

	return nil
}

// ConfigFileExist checks if the config file exists.
// If not, it will create a default config file at the specified location.
// If creation fails, it will return false.
// If the file exists or creation succeeds, it will return true.
func ConfigFileExist() bool {
	configFilename := fmt.Sprintf("%s.%s", CONFIG_FILENAME, CONFIG_FILETYPE)

	if !FileExist(configFilename) {
		fmt.Printf("Config file does not exist. Creating default config at `%s`...\n", configFilename)

		if err := CreateDefaultConfig(configFilename); err != nil {
			fmt.Printf("Error creating config file: %v\n", err)
			return false
		}

		fmt.Println("Default config file created successfully.\nPlease update `config.json` content and re-run again.")
	}

	return true
}
