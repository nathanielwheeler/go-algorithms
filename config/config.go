package config

import (
	"os"

	"gopkg.in/yaml.v2"
)

const (
	configFile = "algorithms.yml"
)

// Config holds configuration variables
type Config struct {
	Algs map[string]AlgorithmConfig `yaml:"algorithms"`
}

// LoadConfig will load production or development configuration files.
func LoadConfig() Config {
	f, err := os.Open(configFile)
	if err != nil {
		panic("No configuration file detected!")
	}
	defer f.Close()
	// Decode file and return Config struct
	var c Config
	d := yaml.NewDecoder(f)
	if err := d.Decode(&c); err != nil {
		panic(err)
	}
	return c
}

// AlgorithmConfig holds metadata and configuration data for an algorithm
type AlgorithmConfig struct {
	FuncMap   map[string]interface{} `yaml:"funcmap"`
	ShortDesc string                 `yaml:"shortdesc"`
	FullDesc  string                 `yaml:"fulldesc"`
	Arguments []ArgumentConfig       `yaml:"arguments"`
	// TODO add flag support
}

// ArgumentConfig holds metadata and configuration data an argument to an algorithm
type ArgumentConfig struct {
	Name    string `yaml:"name"`
	Type    string `yaml:"type"`
	Desc    string `yaml:"desc"`
	Default string `yaml:"default"`
}
