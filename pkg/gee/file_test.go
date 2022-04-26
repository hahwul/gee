package gee

import (
	"os"
	"testing"

	model "github.com/hahwul/gee/pkg/model"
)

func TestWriteFile(t *testing.T) {
	type args struct {
		line    string
		options model.Options
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test removeNewLine",
			args: args{
				line: "1111",
				options: model.Options{
					RemoveNewLine: true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			k, _ := os.OpenFile("gee_testcode_testfile.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			WriteFile(k, tt.args.line, tt.args.options)
			os.Remove("gee_testcode_testfile.txt")
		})
	}
}
