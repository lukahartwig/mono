package module

import (
	"os"
	"os/exec"
)

type Module struct {
	Name string
	Path string
}

func (s *Module) Exec(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Dir = s.Path
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
