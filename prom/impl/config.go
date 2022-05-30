package impl

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

const CONFIG_LOCATION = "$HOME/.config/prom.json"

type Config struct {
	Version string `json:"version"`
	StaleDir string `json:"stale_dir"`
	Projects map[string]string `json:"projects"`
}

func LoadConfig() (Config, error) {
	if _, err := os.Stat(os.ExpandEnv(CONFIG_LOCATION)); errors.Is(err, os.ErrNotExist) {
		fmt.Print("Enter directory path for stale projects: ")
		var staleDir string
		fmt.Scanln(&staleDir)

		config := Config{
			Version: "1",
			StaleDir: staleDir,
			Projects: make(map[string]string),
		}

		err := config.Save()

		return config, err
	}

	data, err := os.ReadFile(os.ExpandEnv(CONFIG_LOCATION))
	if err != nil {
		return Config{}, err
	}

	var config Config
	err = json.Unmarshal(data, &config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func (c *Config) Save() error {
	data, err := json.Marshal(c)
	if err != nil {
		return err
	}

	config_location := os.ExpandEnv(CONFIG_LOCATION)

	if _, err := os.Stat(config_location); errors.Is(err, os.ErrNotExist) {
		os.Create(config_location)
	}

	return os.WriteFile(config_location, data, 0644)
}
