package helper

import (
	"encoding/json"
	"log"
	"os"
)

// Configuration model
type ConfigurationFile struct {
	ConfigurationType string        `json:"config"`
	Development       Configuration `json:"dev"`
	Production        Configuration `json:"prod"`
}

// Configuration model
type Configuration struct {
	Port             int    `json:"port"`
	DatabaseHost     string `json:"databaseHost"`
	DatabasePort     int    `json:"databasePort"`
	DatabaseName     string `json:"databaseName"`
	DatabaseUsername string `json:"databaseUsername"`
	DatabasePassword string `json:"databasePassword"`
	JWTSecret        string `json:"jwtSecret"`
	JWTAdminSecret   string `json:"jwtAdminSecret"`
}

func ReadConfigurationFile() Configuration {
	body, err := os.ReadFile(".config.json")

	if err != nil {
		log.Fatalf("unable to read configuration file: %v", err)
	}

	var configFile ConfigurationFile

	jsonErr := json.Unmarshal(body, &configFile)

	if jsonErr != nil {
		log.Fatalf("unable to read decode json: %v", err)
	}

	var config Configuration

	config.Port = configFile.Production.Port
	config.DatabaseHost = configFile.Production.DatabaseHost
	config.DatabasePort = configFile.Production.DatabasePort
	config.DatabaseName = configFile.Production.DatabaseName
	config.DatabaseUsername = configFile.Production.DatabaseUsername
	config.DatabasePassword = configFile.Production.DatabasePassword
	config.JWTSecret = configFile.Production.JWTSecret

	if configFile.ConfigurationType == "dev" {
		config.Port = configFile.Development.Port
		config.DatabaseHost = configFile.Development.DatabaseHost
		config.DatabasePort = configFile.Development.DatabasePort
		config.DatabaseName = configFile.Development.DatabaseName
		config.DatabaseUsername = configFile.Development.DatabaseUsername
		config.DatabasePassword = configFile.Development.DatabasePassword
		config.JWTSecret = configFile.Development.JWTSecret
	}

	return config
}
