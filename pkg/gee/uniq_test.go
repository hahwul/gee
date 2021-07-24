package gee

import (
	"reflect"
	"testing"
)

func TestUniq(t *testing.T) {
	type args struct {
		line   string
		dtable []string
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want1 []string
	}{
		{
			name: "test true",
			args: args{
				line:   "3",
				dtable: []string{"1", "2", "3"},
			},
			want:  true,
			want1: []string{"1", "2", "3"},
		},
		{
			name: "test false",
			args: args{
				line:   "a",
				dtable: []string{"1", "2", "3"},
			},
			want:  false,
			want1: []string{"1", "2", "3", "a"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := Uniq(tt.args.line, tt.args.dtable)
			if got != tt.want {
				t.Errorf("Uniq() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("Uniq() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
