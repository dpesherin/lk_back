package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path/filepath"
)

type ServerConfig struct {
	Port string `yaml:"port" env-default:"8080"`
	Host string `yaml:"host" env-default:"localhost"`
}
type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}
type Config struct {
	Server   ServerConfig   `yaml:"server"`
	Database DatabaseConfig `yaml:"database"`
}

func LoadConfig(configPath string) (*Config, error) {
	filename, _ := filepath.Abs(configPath)
	yamlFile, err := os.ReadFile(filename)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
		return nil, err
	}

	var cfg Config
	err = yaml.Unmarshal(yamlFile, &cfg)
	if err != nil {
		return nil, err
	}
	return &cfg, nil
}
