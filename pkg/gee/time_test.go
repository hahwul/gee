package gee

import "testing"

func TestGetNowTime(t *testing.T) {
	tests := []struct {
		name    string
		notwant int
	}{
		{
			name:    "test",
			notwant: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetNowTime(); len(got) == tt.notwant {
				t.Errorf("GetNowTime() = %v, want %v", len(got), tt.notwant)
			}
		})
	}
}
