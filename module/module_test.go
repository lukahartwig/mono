package module

import (
	"os/exec"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestModule_Command(t *testing.T) {
	type fields struct {
		Name string
		Path string
	}
	type args struct {
		command string
		args    []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *exec.Cmd
	}{
		{
			"should create command with module as context",
			fields{"module-a", "testdata/valid/module-a"},
			args{
				command: "pwd",
				args:    []string{},
			},
			func() *exec.Cmd {
				cmd := exec.Command("pwd")
				cmd.Dir = "testdata/valid/module-a"
				return cmd
			}(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Module{
				Name: tt.fields.Name,
				Path: tt.fields.Path,
			}
			if got := s.Command(tt.args.command, tt.args.args...); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Module.Command() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestModule_Task(t *testing.T) {
	type fields struct {
		Name  string
		Path  string
		Tasks map[string]Task
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *exec.Cmd
		wantErr bool
	}{
		{
			"should return command from task with arguments",
			fields{
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
			},
			args{
				name: "build",
			},
			func() *exec.Cmd {
				cmd := exec.Command("yarn", "run", "build")
				cmd.Dir = "testdata/valid/module-a"
				return cmd
			}(),
			false,
		},
		{
			"should return command from task without arguments",
			fields{
				Name: "module-a",
				Path: "testdata/valid/module-b",
				Tasks: map[string]Task{
					"build": {
						Command: "make",
						Args:    []string{},
					},
				},
			},
			args{
				name: "build",
			},
			func() *exec.Cmd {
				cmd := exec.Command("make")
				cmd.Dir = "testdata/valid/module-b"
				return cmd
			}(),
			false,
		},
		{
			"should return error when the task does not exist",
			fields{
				Name: "module-a",
				Path: "testdata/valid/module-b",
				Tasks: map[string]Task{
					"build": {
						Command: "make",
						Args:    []string{},
					},
				},
			},
			args{
				name: "task-that-does-not-exist",
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Module{
				Name:  tt.fields.Name,
				Path:  tt.fields.Path,
				Tasks: tt.fields.Tasks,
			}
			got, err := s.Task(tt.args.name)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.EqualValues(t, tt.want, got)
		})
	}
}
