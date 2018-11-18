package module

import (
	"os/exec"

	"github.com/pkg/errors"
)

type Task struct {
	Command string
	Args    []string
}

type Module struct {
	Name  string
	Path  string
	Tasks map[string]Task
}

func (s *Module) GetTask(name string) (*exec.Cmd, error) {
	task, ok := s.Tasks[name]
	if !ok {
		return nil, errors.Errorf("module %s has no task %s", s.Name, name)
	}
	return s.Command(task.Command, task.Args...), nil
}

func (s *Module) Command(command string, args ...string) *exec.Cmd {
	cmd := exec.Command(command, args...)
	cmd.Dir = s.Path
	return cmd
}
