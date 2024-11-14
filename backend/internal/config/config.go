package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io"
	"log"
	"os"
)

type Config struct {
	App     AppConfig     `yaml:"app"`
	Logging LoggingConfig `yaml:"logging"`
}

type AppConfig struct {
	Name        string `yaml:"name"`
	Environment string `yaml:"environment"`
	Port        int    `yaml:"port"`
}

type LoggingConfig struct {
	Level  string `yaml:"level"`
	Format string `yaml:"format"`
	File   string `yaml:"file"`
}

func LoadConfig(path string) (*Config, error) {
	file, err := os.Open(path)
	defer file.Close()

	if err != nil {
		fmt.Printf("The configuration file could not be opened: %v\n", err)
		return nil, err
	}

	var cfg Config
	decoder := yaml.NewDecoder(file)
	if err := decoder.Decode(&cfg); err != nil {
		fmt.Printf("Configuration file could not be decoded: %v\n", err)
		return nil, err
	}

	if err := setupLogging(cfg.Logging); err != nil {
		fmt.Printf("Failed to set up logging: %v\n", err)
		return nil, err
	}

	log.Println("Configuration loaded")
	return &cfg, nil
}

func setupLogging(cfg LoggingConfig) error {
	logFile, err := os.OpenFile(cfg.File, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return fmt.Errorf("could not open log file: %v", err)
	}

	multiWriter := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(multiWriter)
	return nil
}
