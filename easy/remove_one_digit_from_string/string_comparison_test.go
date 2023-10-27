package remove_one_digit_from_string

import (
	"testing"
)

func TestRemoveDigitStringComparison_back(t *testing.T) {
	want := "12"
	got := RemoveDigitStringComparison("123", '3')
	if got != want {
		t.Error("\nwant", want, "\ngot ", got)
	}
}

func TestRemoveDigitStringComparison_front(t *testing.T) {
	want := "231"
	got := RemoveDigitStringComparison("1231", '1')
	if got != want {
		t.Error("\nwant", want, "\ngot ", got)
	}
}

func TestRemoveDigitStringComparison_duplicated(t *testing.T) {
	want := "51"
	got := RemoveDigitStringComparison("551", '5')
	if got != want {
		t.Error("\nwant", want, "\ngot ", got)
	}
}

func TestRemoveDigitStringComparison_middle(t *testing.T) {
	want := "779547853679443616467964135298543163376223791274561861738666981419251859535331546947347395531332878"
	got := RemoveDigitStringComparison("7795478535679443616467964135298543163376223791274561861738666981419251859535331546947347395531332878", '5')
	if got != want {
		t.Error("\nwant", want, "\ngot ", got)
	}
}
