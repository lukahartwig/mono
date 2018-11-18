package module

import (
	"os"
	"path/filepath"
)

// Resolver resolves all the modules
type Resolver interface {
	Resolve() ([]Module, error)
}

type resolver struct {
	FileName string
	Root     string
}

// NewResolver returns a new Resolver instance
func NewResolver(root string) Resolver {
	return &resolver{
		FileName: ".module.yml",
		Root:     root,
	}
}

// Resolve returns all the modules that a below the root directory. A module is
// defined by a .module.yml config file in the directory.
func (s *resolver) Resolve() ([]Module, error) {
	modulePaths := resolvePaths(s.Root, s.FileName)

	modules := make([]Module, len(modulePaths))
	for i, path := range modulePaths {
		module, err := FromConfigFile(path)
		if err != nil {
			return nil, err
		}
		modules[i] = module
	}

	return modules, nil
}

func resolvePaths(root string, fileName string) []string {
	paths := make([]string, 0)
	_ = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if info.Name() == fileName {
			paths = append(paths, path)
		}
		return nil
	})
	return paths
}
