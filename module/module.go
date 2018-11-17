package module

import (
	"os/exec"
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

func (s *Module) Command(command string, args ...string) *exec.Cmd {
	cmd := exec.Command(command, args...)
	cmd.Dir = s.Path
	return cmd
}
