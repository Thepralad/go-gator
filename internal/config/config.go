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

func Read() (Config, error) {
	filePath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(filePath)
	if err != nil {
		return Config{}, err
	}

	defer file.Close()

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

