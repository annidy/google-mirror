package model

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Mirrors []string `yaml:"mirrors"`
}

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
	if cfg, err := readConf(file); err != nil {
		return nil, err
	} else {
		return cfg, nil
	}
}
