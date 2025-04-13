//Package for accessing, opening and interacting with the config.json file in /home
package config

import (
	"encoding/json"
	"log"
	"os"
)

type Config struct {
	Db_url            string `json:"db_url"`
	Current_user_name string `json:"current_user_name"`
}

// Reads the json file and returns instance of Config
func Read() (Config, error) {
	//Getting the json file path
	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	//Opening the json file
	file, err := os.Open(filePath)
	if err != nil {
		return Config{}, err
	}

	defer file.Close()

	//Unmarshaling and appending it to a Config instance and returning it
	decoder := json.NewDecoder(file)
	cfg := Config{}
	err = decoder.Decode(&cfg)
	if err != nil {
		return Config{}, err
	}
	return cfg, nil

}

func (cfg *Config) SetUser(current_user_name string) error {
	cfg.Current_user_name = current_user_name
	err := writer(*cfg)
	if err != nil {
		return err
	}
	return nil
}

// Helper
func getConfigFilePath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		log.Print(err)
		return "", err
	}
	return homePath + "/.gatorconfig.json", nil

}

// Helper
//Overwriting the contents of json file with the new cfg(from the params)
func writer(cfg Config) error {
	path, err := getConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil
}

