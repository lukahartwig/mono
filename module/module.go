package module

import (
	"os/exec"
)

type Module struct {
	Name string
	Path string
}

func (s *Module) Command(command string, args ...string) *exec.Cmd {
	cmd := exec.Command(command, args...)
	cmd.Dir = s.Path
	return cmd
}
