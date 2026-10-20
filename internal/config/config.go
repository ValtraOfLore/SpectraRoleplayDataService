package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Configuration struct {
	Port                     string `json:"PORT"`
	DatabaseConnectionString string `json:"MONGODB_CONNECTION_STRING"`
}

func Load(env string) (*Configuration, error) {
	var configFilePath string
	var config *Configuration

	if env == "local" {
		configFilePath = ".env"

		if err := godotenv.Load(configFilePath); err != nil {
			fmt.Println("Unable to load from local config .env file. Make sure a .env file is available at the top level of the project.")
			return nil, errors.New("Unable to read config")
		}

		config = &Configuration{
			Port:                     os.Getenv("PORT"),
			DatabaseConnectionString: os.Getenv("MONGODB_CONNECTION_STRING"),
		}
	} else {
		var jsonData []byte
		var readErr error
		configFilePath = "configs/configuration-" + env + ".json"

		if _, err := os.Stat(configFilePath); err != nil {
			fmt.Println("Config file not found at path: " + configFilePath)
			return nil, errors.New("Config file does not exist")
		}

		if jsonData, readErr = os.ReadFile(configFilePath); readErr != nil {
			fmt.Println("Something went wrong reading configuration file for environemnt: " + env)
			return nil, errors.New("Unable to read config file")
		}

		if parseErr := json.Unmarshal(jsonData, &config); parseErr != nil {
			fmt.Println("Something went wrong parsing configuration file for environment: " + env)
			return nil, errors.New("Unable to parse config")
		}
	}

	fmt.Println("Loaded configuration successful from " + configFilePath)
	return config, nil
}
