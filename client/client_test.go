package client

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/lukahartwig/mono/module"
)

var (
	moduleA = module.Module{
		Name: "module-a",
		Path: ".",
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
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			"should execute a command and write the output to stdout",
			fields{&mockResolver{
				[]module.Module{moduleA},
			}},
			args{"echo", []string{"hello", "world"}},
			"hello world\n",
			false,
		},
		{
			"should error when a command fails",
			fields{&mockResolver{
				[]module.Module{moduleA},
			}},
			args{"this-is-not-a-command", []string{}},
			"",
			true,
		},
		{
			"should have no output when no modules were found",
			fields{&mockResolver{}},
			args{"pwd", []string{}},
			"",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			out := new(bytes.Buffer)
			s := New(
				tt.fields.resolver,
				&Options{Stdout: out},
			)
			if err := s.Exec(tt.args.command, tt.args.args...); (err != nil) != tt.wantErr {
				t.Errorf("client.Exec() error = %v, wantErr %v", err, tt.wantErr)
			}
			if out.String() != tt.want {
				t.Errorf("client.Exec() stdout = %v, want %v", out.String(), tt.want)
			}
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
