package client

import (
	"github.com/lukahartwig/mono/module"
)

type Client interface {
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

func (s *client) List() ([]module.Module, error) {
	modules, err := s.resolver.Resolve()
	if err != nil {
		return nil, err
	}
	return modules, nil
}
