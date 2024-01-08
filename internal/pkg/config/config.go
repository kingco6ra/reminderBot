package config

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

var Cfg *Config

func init() {
	err := ReadConfigYML("configs/config.yml")
	if err != nil {
		panic(err)
	}
}

// Project - contains all project config.
type Project struct {
	debug bool
}

// Database - contains all database config.
type Database struct {
	Host     string `yaml:"host"`
	Port     uint32 `yaml:"port"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

func (d *Database) GetPostgresDialector() string {
	return fmt.Sprintf(
		"user=%s password=%s dbname=%s host=%s port=%d",
		d.User, d.Password, d.Name, d.Host, d.Port,
	)
}

// Telegram - contains all telegram config.
type Telegram struct {
	Token string `yaml:"token"`
	Debug bool   `yaml:"debug"`
}

// Metrics - contains all metrics config.
type Metrics struct {
	Host string `yaml:"host"`
	Port uint32 `yaml:"port"`
}

// Config - contains all configuration parameters in config package.
type Config struct {
	Project  Project  `yaml:"project"`
	Database Database `yaml:"database"`
	Telegram Telegram `yaml:"telegram"`
	Metrics  Metrics  `yaml:"metrics"`
}

// ReadConfigYML - read configurations from file and init instance Config.
func ReadConfigYML(filePath string) error {
	file, err := os.Open(filepath.Clean(filePath))
	if err != nil {
		return err
	}

	decoder := yaml.NewDecoder(file)
	if decodeErr := decoder.Decode(&Cfg); decodeErr != nil {
		return decodeErr
	}

	err = file.Close()
	if err != nil {
		panic(err.Error())
	}

	return nil
}
