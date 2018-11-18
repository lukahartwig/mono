package client

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"

	"github.com/lukahartwig/mono/module"
)

type Client interface {
	Exec(command string, args ...string) (io.Reader, error)
	List() ([]module.Module, error)
	RunTask(name string) (io.Reader, error)
}

type client struct {
	resolver module.Resolver
}

func New(resolver module.Resolver) Client {
	return &client{
		resolver: resolver,
	}
}

func (s *client) Exec(command string, args ...string) (io.Reader, error) {
	modules, err := s.resolver.Resolve()
	if err != nil {
		return nil, err
	}
	return s.execSync(modules, command, args...), nil
}

func (s *client) List() ([]module.Module, error) {
	return s.resolver.Resolve()
}

// RunTask runs the task with the given name in every module
func (s *client) RunTask(name string) (io.Reader, error) {
	modules, err := s.resolver.Resolve()
	if err != nil {
		return nil, err
	}
	return s.runTaskSync(modules, name), nil
}

func (s *client) execSync(modules []module.Module, command string, args ...string) io.Reader {
	cmds := make([]*exec.Cmd, 0)
	for _, m := range modules {
		cmds = append(cmds, m.Command(command, args...))
	}
	return s.execCommandSync(cmds...)
}

func (s *client) runTaskSync(modules []module.Module, task string) io.Reader {
	cmds := make([]*exec.Cmd, 0)
	for _, m := range modules {
		cmd, err := m.GetTask(task)
		if err != nil {
			fmt.Printf("skipping task: %s\n", err)
			continue
		}
		cmds = append(cmds, cmd)
	}
	return s.execCommandSync(cmds...)
}

func (s *client) execCommandSync(cmds ...*exec.Cmd) io.Reader {
	pr, pw := io.Pipe()
	go func() {
		for _, cmd := range cmds {
			out, err := cmd.CombinedOutput()
			if err != nil {
				_ = pw.CloseWithError(err)
				return
			}
			if _, err := io.Copy(pw, bytes.NewReader(out)); err != nil {
				_ = pw.CloseWithError(err)
				return
			}
		}
		_ = pw.Close()
	}()
	return pr
}
