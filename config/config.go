package config

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

const (
	configDir       = "config"
	configExtension = "yaml"
)

type DBConfig struct {
	Dialect   string `yaml:"dialect"`
	Username  string `yaml:"username"`
	Password  string `yaml:"password"`
	Host      string `yaml:"host"`
	Port      int    `yaml:"port"`
	Name      string `yaml:"name"`
	Charset   string `yaml:"charset"`
	ParseTime bool   `yaml:"parseTime"`
}

type AppConfig struct {
	App struct {
		Name string `yaml:"name"`
		Port int    `yaml:"port"`
		Env  string `yaml:"env"`
	} `yaml:"app"`
	Database DBConfig `yaml:"database"`
}

func (a AppConfig) Port() string {
	return fmt.Sprintf(":%d", a.App.Port)
}

func (a AppConfig) Env() string {
	return a.App.Env
}

func LoadConfig(env string) (*AppConfig, error) {
	fileName := fmt.Sprintf("%s/%s.%s", configDir, env, configExtension)
	buf, err := os.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	var conf AppConfig
	if err = yaml.Unmarshal(buf, &conf); err != nil {
		return nil, err
	}
	return &conf, nil
}
