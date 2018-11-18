package module

import (
	"os/exec"

	"github.com/pkg/errors"
)

// Task can be run in the module
type Task struct {
	Command string
	Args    []string
}

// Module represents a module defined by a .module.yml file under the root
// directory.
type Module struct {
	Name  string
	Path  string
	Tasks map[string]Task
}

// GetTask returns the command for the module that is associated to the task
// with the given name.
func (s *Module) GetTask(name string) (*exec.Cmd, error) {
	task, ok := s.Tasks[name]
	if !ok {
		return nil, errors.Errorf("module %s has no task %s", s.Name, name)
	}
	return s.Command(task.Command, task.Args...), nil
}

// Command returns a command that will be executed in the context of the
// module.
func (s *Module) Command(command string, args ...string) *exec.Cmd {
	cmd := exec.Command(command, args...)
	cmd.Dir = s.Path
	return cmd
}
