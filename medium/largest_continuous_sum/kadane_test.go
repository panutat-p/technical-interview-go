package largest_continuous_sum

import "testing"

func TestLargestSumKadane(t *testing.T) {
	tests := []struct {
		name  string
		given []int
		want  int
	}{
		{
			name:  "one positive",
			given: []int{3},
			want:  3,
		},
		{
			name:  "one negative",
			given: []int{-10},
			want:  -10,
		},
		{
			name:  "filter minus on the left",
			given: []int{-50, 3, 6, 4, 2},
			want:  15,
		},
		{
			name:  "big wall",
			given: []int{-1, -2, -3, 50, -4, -5},
			want:  50,
		},
		{
			name:  "filter minus on the right",
			given: []int{3, 6, 4, 2, -50},
			want:  15,
		},
		{
			name:  "ignore minus portion",
			given: []int{1, 2, -1, 3, 10, 10, -10 - 1},
			want:  25,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := LargestSumKadane(tt.given)
			if got != tt.want {
				t.Errorf("\nWant %v\nGot %v", tt.want, got)
			}
		})
	}
}
