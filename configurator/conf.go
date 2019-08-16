package configurator

import (
	"github.com/dev2choiz/f7k/interfaces"
)

type Config struct {
	Port          int    `yaml:"port"`
	ImportPath    string `yaml:"importPath"`
	ControllerDir string `yaml:"controllerDir"`

	View       	  interfaces.ConfViewer
}

func (c *Config) GetPort() int {
	return c.Port
}

func (c *Config) GetImportPath() string {
	return c.ImportPath
}

func (c *Config) GetControllerDir() string {
	return c.ControllerDir
}

func (c *Config) PostConfig() {
}

func (c *Config) Data() interfaces.ConfigInterface {
	return c
}
