package gee

import "testing"

func TestRegex(t *testing.T) {
	type args struct {
		pattern string
		str     string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test - true",
			args: args{
				pattern: "abcd",
				str:     "123abcd444",
			},
			want: true,
		},
		{
			name: "test - false",
			args: args{
				pattern: "abcd",
				str:     "123d444",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Regex(tt.args.pattern, tt.args.str); got != tt.want {
				t.Errorf("Regex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRegexV(t *testing.T) {
	type args struct {
		pattern string
		str     string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "test - true",
			args: args{
				pattern: "abcd",
				str:     "123d444",
			},
			want: true,
		},
		{
			name: "test - false",
			args: args{
				pattern: "abcd",
				str:     "123abcd444",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RegexV(tt.args.pattern, tt.args.str); got != tt.want {
				t.Errorf("RegexV() = %v, want %v", got, tt.want)
			}
		})
	}
}
