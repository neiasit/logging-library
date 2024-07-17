package internal

import (
	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	MODE string `json:"mode" env:"MODE" env-default:"prod"`
}

func LoadConfig() (*Config, error) {
	var cfg struct {
		Config Config `json:"logger" env-prefix:"LOGGER_"`
	}
	err := cleanenv.ReadConfig("config.json", &cfg)
	if err != nil {
		err := cleanenv.ReadEnv(&cfg)
		if err != nil {
			return nil, err
		}
	}
	return &cfg.Config, nil
}
