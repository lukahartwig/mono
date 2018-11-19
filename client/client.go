package client

import (
	"bytes"
	"fmt"
	"io"
	"os/exec"

	"github.com/lukahartwig/mono/module"
)

type ExecOptions struct {
	Included []string
}

type RunOptions struct {
	Included []string
}

// Client is the programmatically API for the modules handling and task
// executions.
type Client interface {
	Exec(command string, args []string, opts *ExecOptions) (io.Reader, error)
	List() ([]module.Module, error)
	Run(name string, opts *RunOptions) (io.Reader, error)
}

type client struct {
	resolver module.Resolver
}

// New returns a new client instance
func New(resolver module.Resolver) Client {
	return &client{
		resolver: resolver,
	}
}

// Exec runs a command in every module
func (s *client) Exec(command string, args []string, opts *ExecOptions) (io.Reader, error) {
	modules, err := s.resolve(opts.Included)
	if err != nil {
		return nil, err
	}
	return s.execSync(modules, command, args...), nil
}

// List returns a list of all modules
func (s *client) List() ([]module.Module, error) {
	return s.resolver.Resolve()
}

// Run runs the task with the given name in every module
func (s *client) Run(name string, opts *RunOptions) (io.Reader, error) {
	modules, err := s.resolve(opts.Included)
	if err != nil {
		return nil, err
	}
	return s.runTaskSync(modules, name), nil
}

func (s *client) resolve(included []string) ([]module.Module, error) {
	resolved, err := s.resolver.Resolve()
	if err != nil {
		return nil, err
	}

	var modules []module.Module
	if included != nil {
		modules = applyIncluded(resolved, included)
	} else {
		modules = resolved
	}

	return modules, nil
}

func applyIncluded(resolved []module.Module, included []string) []module.Module {
	modules := make([]module.Module, 0)
	for _, m := range resolved {
		for _, name := range included {
			if name == m.Name {
				modules = append(modules, m)
				continue
			}
		}
	}
	return modules
}

func (s *client) execSync(modules []module.Module, command string, args ...string) io.Reader {
	cmds := make([]*exec.Cmd, 0)
	for _, m := range modules {
		cmds = append(cmds, m.Command(command, args...))
	}
	return execCommandSync(cmds...)
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
	return execCommandSync(cmds...)
}

func execCommandSync(cmds ...*exec.Cmd) io.Reader {
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
