package client

import (
	"fmt"

	"github.com/lukahartwig/mono/module"
)

type Client interface {
	Exec(command string, args ...string) error
	List() ([]module.Module, error)
}

type Options struct {
	Root string
}

type client struct {
	resolver *module.Resolver
}

func New(opts *Options) Client {
	resolver := &module.Resolver{
		FileName: ".module.yml",
		Root:     opts.Root,
	}

	return &client{
		resolver: resolver,
	}
}

func (s *client) Exec(command string, args ...string) error {
	modules, err := s.resolver.Resolve()
	if err != nil {
		return err
	}
	for _, m := range modules {
		if err := m.Exec(command, args...); err != nil {
			fmt.Printf("%s: command failed: %s\n", m.Name, err)
		}
	}
	return nil
}

func (s *client) List() ([]module.Module, error) {
	modules, err := s.resolver.Resolve()
	if err != nil {
		return nil, err
	}
	return modules, nil
}
