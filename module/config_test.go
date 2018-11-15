package module

import (
	"reflect"
	"testing"
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
			Module{
				Name: "module-a",
			},
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := FromConfigFile(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("FromConfigFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FromConfigFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
