package module

import (
	"os"
	"path/filepath"
)

type Resolver interface {
	Resolve() ([]Module, error)
}

type resolver struct {
	FileName string
	Root     string
}

func NewResolver(root string) Resolver {
	return &resolver{
		FileName: ".module.yml",
		Root:     root,
	}
}

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
