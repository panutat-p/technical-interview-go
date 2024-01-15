package compression

import (
	"testing"
)

func TestCompress(t *testing.T) {
	tests := []struct {
		name  string
		given []byte
		want1 int
		want2 []byte
	}{
		{
			name:  "single character",
			given: []byte("a"),
			want1: 1,
			want2: []byte("a"),
		},
		{
			name:  "multiple characters",
			given: []byte("aabbccc"),
			want1: 6,
			want2: []byte("a2b2c3"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got1 := Compress(tt.given)
			if got1 != tt.want1 {
				t.Errorf("\nWant %v\nGot %v", tt.want1, got1)
			}
			got2 := tt.given[:tt.want1]
			if string(got2) != string(tt.want2) {
				t.Errorf("\nWant %v\nGot %v", string(tt.want2), string(got2))
			}
		})
	}
}
