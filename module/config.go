package module

import (
	"io/ioutil"
	"path/filepath"

	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
)

type commandConfig []string

type config struct {
	Name  string                   `yaml:"name"`
	Tasks map[string]commandConfig `yaml:"tasks"`
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

	tasks := make(map[string]Task)
	for name, command := range config.Tasks {
		if len(command) < 1 {
			return Module{}, errors.Errorf("module %s has empty task: %s", config.Name, name)
		}

		tasks[name] = Task{
			Command: command[0],
			Args:    command[1:],
		}
	}

	return Module{
		Name:  config.Name,
		Path:  filepath.Dir(path),
		Tasks: tasks,
	}, nil
}
