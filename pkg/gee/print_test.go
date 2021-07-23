package gee

import (
	"testing"

	model "github.com/hahwul/gee/pkg/model"
)

func TestStdPrint(t *testing.T) {
	type args struct {
		line    string
		options model.Options
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test",
			args: args{
				line: "",
				options: model.Options{
					RemoveNewLine: true,
				},
			},
		},
		{
			name: "test",
			args: args{
				line: "",
				options: model.Options{
					RemoveNewLine: false,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			StdPrint(tt.args.line, tt.args.options)
		})
	}
}
