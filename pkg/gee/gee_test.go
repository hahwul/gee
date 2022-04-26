package gee

import (
	"os"
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
			name: "test debug",
			args: args{
				options: model.Options{
					Append: true,
					Files:  []string{"", ""},
					Debug:  true,
				},
			},
		},
		{
			name: "test uniq",
			args: args{
				options: model.Options{
					Append: true,
					Files:  []string{"", ""},
					Uniq:   true,
				},
			},
		},
		{
			name: "test regex",
			args: args{
				options: model.Options{
					Append: false,
					Files:  []string{"", ""},
					Regex:  "HAHWUL",
				},
			},
		},
		{
			name: "test regexv",
			args: args{
				options: model.Options{
					Append: false,
					Files:  []string{"", ""},
					RegexV: "HAHWUL",
				},
			},
		},
		{
			name: "test chunkedLine",
			args: args{
				options: model.Options{
					Append:      false,
					Files:       []string{"", ""},
					ChunkedLine: 1,
				},
			},
		},
		{
			name: "test distribute",
			args: args{
				options: model.Options{
					Append:     false,
					Files:      []string{"", ""},
					Distribute: true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := []byte("Alice")
			r, w, err := os.Pipe()
			if err != nil {
				t.Fatal(err)
			}

			_, err = w.Write(input)
			if err != nil {
				t.Error(err)
			}
			w.Close()
			stdin := os.Stdin
			// Restore stdin right after the test.
			defer func() { os.Stdin = stdin }()
			os.Stdin = r
			Gee(tt.args.options)
		})
	}
}
