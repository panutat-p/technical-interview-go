package pangram

import "testing"

func TestCheckByMap(t *testing.T) {
	t.Run("a to z", func(t *testing.T) {
		want := true
		got := CheckByMap("abcdefghijklmnopqrstuvwxyz")
		if got != want {
			t.Error("\nwant", want, "\ngot ", got)
		}
	})
	t.Run("thequickbrownfoxjumpsoverthelazydog", func(t *testing.T) {
		want := true
		got := CheckByMap("thequickbrownfoxjumpsoverthelazydog")
		if got != want {
			t.Error("\nwant", want, "\ngot ", got)
		}
	})
	t.Run("leetcode", func(t *testing.T) {
		want := false
		got := CheckByMap("leetcode")
		if got != want {
			t.Error("\nwant", want, "\ngot ", got)
		}
	})
}
