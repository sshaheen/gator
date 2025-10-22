package config

import (
	"fmt"
	"os"
)

const conf = ".gatorconfig.json"

func getConfigFilePath() (string, error) {
	home_dir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	config_path := fmt.Sprintf("%s/%s", home_dir, conf)
	return config_path, nil
}
