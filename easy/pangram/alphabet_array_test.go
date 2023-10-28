package pangram

import "testing"

func TestAlphabetArray(t *testing.T) {
	t.Run("a to z", func(t *testing.T) {
		want := true
		got := AlphabetArray("abcdefghijklmnopqrstuvwxyz")
		if got != want {
			t.Error("\nwant", want, "\ngot ", got)
		}
	})
	t.Run("thequickbrownfoxjumpsoverthelazydog", func(t *testing.T) {
		want := true
		got := AlphabetArray("thequickbrownfoxjumpsoverthelazydog")
		if got != want {
			t.Error("\nwant", want, "\ngot ", got)
		}
	})
	t.Run("leetcode", func(t *testing.T) {
		want := false
		got := AlphabetArray("leetcode")
		if got != want {
			t.Error("\nwant", want, "\ngot ", got)
		}
	})
}
