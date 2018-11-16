package module

import (
	"reflect"
	"testing"
)

func TestResolver_Resolve(t *testing.T) {
	type fields struct {
		FileName string
		Root     string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []Module
		wantErr bool
	}{
		{
			"should return all modules of a valid root",
			fields{
				FileName: ".module.yml",
				Root:     "./testdata/valid",
			},
			[]Module{
				{"module-a", "testdata/valid/module-a"},
				{"module-b", "testdata/valid/module-b"},
				{"module-nested", "testdata/valid/module-nested"},
				{"module-nested-child", "testdata/valid/module-nested/module-nested-child"},
			},
			false,
		},
		{
			"should return error if a module is invalid",
			fields{
				FileName: ".module.yml",
				Root:     "./testdata/invalid",
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Resolver{
				FileName: tt.fields.FileName,
				Root:     tt.fields.Root,
			}
			got, err := s.Resolve()
			if (err != nil) != tt.wantErr {
				t.Errorf("Resolver.Resolve() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Resolver.Resolve() = %v, want %v", got, tt.want)
			}
		})
	}
}
