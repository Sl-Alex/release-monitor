package config

import (
	"encoding/json"
	"os"

	"release-monitor/model"
)

type Config struct {
	Apps []model.AppConfig `json:"apps"`
}

func LoadConfig(path string) (Config, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Config{}, err
	}

	var cfg Config
	if err := json.Unmarshal(data, &cfg); err != nil {
		return Config{}, err
	}

	if err := validate(cfg); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
