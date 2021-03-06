package client

import (
	"io/ioutil"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/lukahartwig/mono/module"
)

var (
	moduleA = module.Module{
		Name: "module-a",
		Path: ".",
		Tasks: map[string]module.Task{
			"build": {
				Command: "echo",
				Args:    []string{"running", "build", "task"},
			},
			"invalid-task": {
				Command: "this-is-not-a-command",
				Args:    []string{},
			},
		},
	}
	moduleB = module.Module{
		Name: "module-b",
		Path: ".",
		Tasks: map[string]module.Task{
			"build": {
				Command: "echo",
				Args:    []string{"running", "build", "in", "module-b"},
			},
		},
	}
)

type mockResolver struct {
	modules []module.Module
}

func (s *mockResolver) Resolve() ([]module.Module, error) {
	return s.modules, nil
}

func Test_client_Exec(t *testing.T) {
	type fields struct {
		resolver module.Resolver
	}
	type args struct {
		command string
		args    []string
		opts    *ExecOptions
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			"should execute a command in every module",
			fields{
				&mockResolver{
					[]module.Module{moduleA},
				},
			},
			args{
				"echo",
				[]string{"hello", "world"},
				&ExecOptions{},
			},
			"hello world\n",
			false,
		},
		{
			"should error when a command fails",
			fields{
				&mockResolver{
					[]module.Module{moduleA},
				},
			},
			args{
				"this-is-not-a-command",
				[]string{},
				&ExecOptions{},
			},
			"",
			true,
		},
		{
			"should have no output when no modules were found",
			fields{
				&mockResolver{},
			},
			args{
				"pwd",
				[]string{},
				&ExecOptions{},
			},
			"",
			false,
		},
		{
			"should only run command for included modules",
			fields{
				&mockResolver{
					[]module.Module{moduleA, moduleB},
				},
			},
			args{
				"echo",
				[]string{"hello", "world"},
				&ExecOptions{
					[]string{moduleB.Name},
				},
			},
			"hello world\n",
			false,
		},
		{
			"should run no command when included modules do not exist",
			fields{
				&mockResolver{
					[]module.Module{moduleA, moduleB},
				},
			},
			args{
				"echo",
				[]string{"hello", "world"},
				&ExecOptions{
					Included: []string{"module-that-does-not-exist"},
				},
			},
			"",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New(tt.fields.resolver)
			out, _ := s.Exec(tt.args.command, tt.args.args, tt.args.opts)
			got, err := ioutil.ReadAll(out)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.EqualValues(t, tt.want, string(got))
		})
	}
}

func Test_client_List(t *testing.T) {
	type fields struct {
		resolver module.Resolver
	}
	tests := []struct {
		name    string
		fields  fields
		want    []module.Module
		wantErr bool
	}{
		{
			"should return all modules",
			fields{&mockResolver{[]module.Module{moduleA}}},
			[]module.Module{moduleA},
			false,
		},
		{
			"should return an empty list when no modules are found",
			fields{&mockResolver{}},
			nil,
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &client{
				resolver: tt.fields.resolver,
			}
			got, err := s.List()
			if (err != nil) != tt.wantErr {
				t.Errorf("client.List() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("client.List() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_Run(t *testing.T) {
	type fields struct {
		resolver module.Resolver
	}
	type args struct {
		name string
		opts *RunOptions
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			"should run task in every module",
			fields{
				&mockResolver{
					[]module.Module{moduleA},
				},
			},
			args{
				"build",
				&RunOptions{},
			},
			"running build task\n",
			false,
		},
		{
			"should skip modules that do not have the task",
			fields{
				&mockResolver{
					[]module.Module{moduleA},
				},
			},
			args{
				"task-that-does-not-exist",
				&RunOptions{},
			},
			"",
			false,
		},
		{
			"should return an error when the task is invalid",
			fields{
				&mockResolver{
					[]module.Module{moduleA},
				},
			},
			args{
				"invalid-task",
				&RunOptions{},
			},
			"",
			true,
		},
		{
			"should only run tasks for included modules",
			fields{
				&mockResolver{
					[]module.Module{moduleA, moduleB},
				},
			},
			args{
				"build",
				&RunOptions{
					Included: []string{moduleB.Name},
				},
			},
			"running build in module-b\n",
			false,
		},
		{
			"should run no tasks when included modules do not exist",
			fields{
				&mockResolver{
					[]module.Module{moduleA, moduleB},
				},
			},
			args{
				"build",
				&RunOptions{
					Included: []string{"module-that-does-not-exist"},
				},
			},
			"",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := New(tt.fields.resolver)
			out, _ := s.Run(tt.args.name, tt.args.opts)
			got, err := ioutil.ReadAll(out)
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.EqualValues(t, tt.want, string(got))
		})
	}
}
