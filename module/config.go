package module

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type config struct {
	Name string `yaml:"name"`
}

func FromConfigFile(path string) (Module, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		return Module{}, err
	}

	var config config
	if err := yaml.Unmarshal(buf, &config); err != nil {
		return Module{}, err
	}

	return Module{
		Name: config.Name,
		Path: filepath.Dir(path),
	}, nil
}
