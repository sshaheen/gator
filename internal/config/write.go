package config

import (
	"encoding/json"
	"os"
)

func Write(cfg *Config) error {
	config_path, err := getConfigFilePath()
	if err != nil {
		return err
	}

	data, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	os.WriteFile(config_path, data, 0644)

	return nil
}
