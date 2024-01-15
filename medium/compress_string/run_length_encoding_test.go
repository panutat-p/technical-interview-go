package compress_string

import (
	"testing"
)

func TestCompressString(t *testing.T) {
	tests := []struct {
		name  string
		given string
		want  string
	}{
		{
			name:  "single character",
			given: "a",
			want:  "a",
		},
		{
			name:  "multiple characters",
			given: "aabbccc",
			want:  "a2b2c3",
		},
		{
			name:  "AB",
			given: "AAAABBB",
			want:  "A4B3",
		},
		{
			name:  "ABCDE",
			given: "AAAABBBBCCCCCDDEEEE",
			want:  "A4B4C5D2E4",
		},
		{
			name:  "duplicated character",
			given: "AAABBAAAA",
			want:  "A3B2A4",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := Compress(tt.given)
			if got != tt.want {
				t.Errorf("\nWant %v\nGot %v", tt.want, got)
			}
		})
	}
}
