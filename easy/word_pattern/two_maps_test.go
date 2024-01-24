package word_pattern

import "testing"

func TestTwoMaps(t *testing.T) {
	tests := []struct {
		name    string
		pattern string
		input   string
		want    bool
	}{
		{
			name:    "two words",
			pattern: "abba",
			input:   "dog cat cat dog",
			want:    true,
		},
		{
			name:    "three words",
			pattern: "abba",
			input:   "dog cat cat fish",
			want:    false,
		},
		{
			name:    "all a",
			pattern: "aaaa",
			input:   "dog cat cat dog",
			want:    false,
		},
		{
			name:    "all dog",
			pattern: "abba",
			input:   "dog dog dog dog",
			want:    false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TwoMaps(tt.pattern, tt.input)
			if got != tt.want {
				t.Errorf("\nWant %v\nGot %v", tt.want, got)
			}
		})
	}
}
