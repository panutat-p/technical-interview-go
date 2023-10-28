package remove_one_digit_from_string

import "testing"

func TestSignificantDigit(t *testing.T) {
	t.Run("remove at back", func(t *testing.T) {
		want := "12"
		got := SignificantDigit("123", '3')
		if got != want {
			t.Error("\nwant", want, "\ngot ", got)
		}
	})
	t.Run("remove at front", func(t *testing.T) {
		want := "231"
		got := SignificantDigit("1231", '1')
		if got != want {
			t.Error("\nwant", want, "\ngot ", got)
		}
	})
	t.Run("duplicated", func(t *testing.T) {
		want := "51"
		got := SignificantDigit("551", '5')
		if got != want {
			t.Error("\nwant", want, "\ngot ", got)
		}
	})
	t.Run("remove at mid", func(t *testing.T) {
		want := "779547853679443616467964135298543163376223791274561861738666981419251859535331546947347395531332878"
		got := SignificantDigit("7795478535679443616467964135298543163376223791274561861738666981419251859535331546947347395531332878", '5')
		if got != want {
			t.Error("\nwant", want, "\ngot ", got)
		}
	})
}
