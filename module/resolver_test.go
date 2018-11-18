package module

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResolver_Resolve(t *testing.T) {
	type fields struct {
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
				Root:     "./testdata/valid",
			},
			[]Module{
				moduleA,
				moduleB,
				moduleNested,
				moduleNestedChild,
			},
			false,
		},
		{
			"should return error if a module is invalid",
			fields{
				Root:     "./testdata/invalid",
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewResolver(tt.fields.Root).Resolve()
			if tt.wantErr {
				assert.Error(t, err)
				return
			}
			assert.EqualValues(t, tt.want, got)
		})
	}
}
