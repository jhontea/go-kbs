package config

import (
	"io/ioutil"
	"log"
	"sync"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

var doOnce sync.Once
var singleton *Config

// NewConfig initialize config object
func NewConfig() *Config {
	doOnce.Do(func() {
		cfg, err := readCfg("./config.yaml")
		if err != nil {
			log.Fatalf(err.Error())
		}
		singleton = cfg
	})
	return singleton
}

// GetConfig :nodoc:
func GetConfig() *Config {
	if singleton != nil {
		return singleton
	}

	return &Config{
		App: &App{},
	}
}

// Config :nodoc:
type Config struct {
	App *App `yaml:"app"`
}

// App :nodoc:
type App struct {
	Port         string `yaml:"port"`
	WriteTimeout int    `yaml:"write_timeout"`
	ReadTimeout  int    `yaml:"read_timeout"`
}

func readCfg(fname string) (*Config, error) {
	var cfg *Config

	err := ReadFromYAML(fname, &cfg)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to read file")
	}

	if cfg == nil {
		return nil, errors.New("No config file found")
	}

	return cfg, nil
}

// ReadFromYAML :nodoc:
func ReadFromYAML(path string, target interface{}) error {
	yf, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}
	return yaml.Unmarshal(yf, target)
}
