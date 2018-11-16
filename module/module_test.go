package module

import (
	"os/exec"
	"reflect"
	"testing"
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
