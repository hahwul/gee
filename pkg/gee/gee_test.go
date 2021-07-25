package gee

import (
	"testing"

	model "github.com/hahwul/gee/pkg/model"
)

func TestGee(t *testing.T) {
	type args struct {
		options model.Options
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{
				options: model.Options{
					Append: true,
					Files:  []string{"", ""},
					Debug:  true,
				},
			},
		},
		{
			name: "test",
			args: args{
				options: model.Options{
					Append: true,
					Files:  []string{"", ""},
					Uniq:   true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Gee(tt.args.options)
		})
	}
}
