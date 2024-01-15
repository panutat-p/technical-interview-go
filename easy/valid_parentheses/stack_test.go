package valid_parentheses

import "testing"

func TestIsValid(t *testing.T) {
	tests := []struct {
		name  string
		given string
		want  bool
	}{
		{
			name:  "parentheses",
			given: "()",
			want:  true,
		},
		{
			name:  "parentheses, brackets, braces",
			given: "()[]{}",
			want:  true,
		},
		{
			name:  "valid nested braces",
			given: "[{()}]",
			want:  true,
		},
		{
			name:  "not match",
			given: "(]",
			want:  false,
		},
		{
			name:  "no left braces",
			given: "))",
			want:  false,
		},
		{
			name:  "imbalance brackets",
			given: "[[[]])]",
			want:  false,
		},
		{
			name:  "nested and mixed",
			given: "[{{{(())}}}]((()))",
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := IsValid(tt.given)
			if got != tt.want {
				t.Errorf("\nWant %v\nGot %v", tt.want, got)
			}
		})
	}
}
