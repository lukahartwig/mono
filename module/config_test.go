package module

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	moduleA = Module{
		Name: "module-a",
		Path: "testdata/valid/module-a",
		Tasks: map[string]Task{
			"build": {
				Command: "yarn",
				Args:    []string{"run", "build"},
			},
			"test": {
				Command: "yarn",
				Args:    []string{"run", "test"},
			},
		},
	}
	moduleB = Module{
		Name: "module-b",
		Path: "testdata/valid/module-b",
		Tasks: map[string]Task{
			"build": {
				Command: "make",
				Args: []string{},
			},
		},
	}
	moduleNested = Module{
		Name: "module-nested",
		Path: "testdata/valid/module-nested",
		Tasks: make(map[string]Task),
	}
	moduleNestedChild = Module{
		Name: "module-nested-child",
		Path: "testdata/valid/module-nested/module-nested-child",
		Tasks: make(map[string]Task),
	}
)

func TestFromConfigFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    Module
		wantErr bool
	}{
		{
			"should return a module from the config file",
			args{"./testdata/valid/module-a/.module.yml"},
			moduleA,
			false,
		},
		{
			"should return a module when task has no arguments",
			args{"./testdata/valid/module-b/.module.yml"},
			moduleB,
			false,
		},
		{
			"should return error when the config file does not exist",
			args{"./testdata/valid/no-module/.module.yml"},
			Module{},
			true,
		},
		{
			"should return error when the config file is invalid",
			args{"./testdata/invalid/invalid-module/.module.yml"},
			Module{},
			true,
		},
		{
			"should return error when a task has no command",
			args{"./testdata/invalid/task-without-cmd/.module.yml"},
			Module{},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromConfigFile(tt.args.path)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.EqualValues(t, tt.want, got)
		})
	}
}
