package gee

import (
	"testing"

	"github.com/hahwul/gee/pkg/model"
)

func TestStringProc(t *testing.T) {
	type args struct {
		l       string
		stdLine int
		options model.Options
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{
				l:       "a1234",
				stdLine: 0,
				options: model.Options{
					Find:     "a",
					Replace:  "b",
					Format:   "",
					WithLine: true,
				},
			},
			want: "[0] b1234",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, _ = StringProc(tt.args.l, tt.args.stdLine, tt.args.options)
		})
	}
}

func Test_toFormat(t *testing.T) {
	type args struct {
		str     string
		options model.Options
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toFormat(tt.args.str, tt.args.options); got != tt.want {
				t.Errorf("toFormat() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setFix(t *testing.T) {
	type args struct {
		str     string
		options model.Options
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := setFix(tt.args.str, tt.args.options); got != tt.want {
				t.Errorf("setFix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setPrefix(t *testing.T) {
	type args struct {
		str     string
		options model.Options
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := setPrefix(tt.args.str, tt.args.options); got != tt.want {
				t.Errorf("setPrefix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setSuffix(t *testing.T) {
	type args struct {
		str     string
		options model.Options
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := setSuffix(tt.args.str, tt.args.options); got != tt.want {
				t.Errorf("setSuffix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_setReverse(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test",
			args: args{
				s: "abcd",
			},
			want: "dcba",
		},
		{
			name: "test",
			args: args{
				s: "1234",
			},
			want: "4321",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := setReverse(tt.args.s); got != tt.want {
				t.Errorf("setReverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
