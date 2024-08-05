package model

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ScreenshotTimeout int64    `yaml:"screenshot_timeout"`
	Mirrors           []string `yaml:"mirrors"`
	Blacklist         []string `yaml:"blacklist"`
}

var ConfigFile = "config.yml"

func readConf(filename string) (*Config, error) {
	buf, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	c := &Config{}
	err = yaml.Unmarshal(buf, c)
	if err != nil {
		return nil, fmt.Errorf("in file %q: %w", filename, err)
	}

	return c, err
}

func LoadConfig(file string) (*Config, error) {
	ConfigFile = file
	if cfg, err := readConf(file); err != nil {
		return nil, err
	} else {
		return cfg, nil
	}
}

func ReloadConfig() (*Config, error) {
	return LoadConfig(ConfigFile)
}

func MustConfig(cfg *Config, err error) *Config {
	if err != nil {
		panic(err)
	}
	return cfg
}
