package main

import (
	"github.com/BurntSushi/toml"
)

type Config struct {
	Secrets Secrets `toml:"secrets"`
	Server  Server  `toml:"server"`
}

type Secrets struct {
	GiteaWebhookSecret  string `toml:"GITEA_WEBHOOK_SECRET"`
	GitlabWebhookSecret string `toml:"GITLAB_WEBHOOK_SECRET"`
}

type Server struct {
	ServerIP   string `toml:"WEBHOOK_HOST"`
	ServerPort int    `toml:"WEBHOOK_PORT"`
}

func LoadConfig(filepath string) (Config, error) {
	var cfg Config
	_, err := toml.DecodeFile(filepath, &cfg)
	if err != nil {
		return cfg, err
	}
	return cfg, nil
}
