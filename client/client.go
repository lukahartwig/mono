package client

import (
	"io"
	"os"

	"github.com/lukahartwig/mono/module"
)

type Client interface {
	Exec(command string, args ...string) error
	List() ([]module.Module, error)
}

type Options struct {
	Stdout io.Writer
	Stderr io.Writer
}

type client struct {
	resolver module.Resolver

	stdout io.Writer
	stderr io.Writer
}

func New(resolver module.Resolver, opts *Options) Client {
	if opts.Stdout == nil {
		opts.Stdout = os.Stdout
	}
	if opts.Stderr == nil {
		opts.Stderr = os.Stderr
	}

	return &client{
		resolver: resolver,

		stdout: opts.Stdout,
		stderr: opts.Stderr,
	}
}

func (s *client) Exec(command string, args ...string) error {
	modules, err := s.resolver.Resolve()
	if err != nil {
		return err
	}
	return s.execSync(modules, command, args...)
}

func (s *client) List() ([]module.Module, error) {
	return s.resolver.Resolve()
}

func (s *client) execSync(modules []module.Module, command string, args ...string) error {
	for _, m := range modules {
		cmd := m.Command(command, args...)
		cmd.Stdout = s.stdout
		cmd.Stderr = s.stderr
		if err := cmd.Run(); err != nil {
			return err
		}
	}
	return nil
}
