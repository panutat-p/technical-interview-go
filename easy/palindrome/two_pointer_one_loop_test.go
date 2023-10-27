package palindrome

import "testing"

func TestIsPalindrome_comma(t *testing.T) {
	yes := IsPalindrome("A man, a plan, a canal, Panama!")
	if !yes {
		t.Error("want true")
	}
}
