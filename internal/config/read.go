package config

import (
	"encoding/json"
	"os"
)

func Read() (Config, error) {
	config_path, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.ReadFile(config_path)
	if err != nil {
		return Config{}, err
	}

	var data Config
	json.Unmarshal(file, &data)

	return data, nil
}
