package anagram

import (
	"unicode"
)

// Check2
// https://stackoverflow.com/a/13215807
// https://stackoverflow.com/a/4237875
// https://stackoverflow.com/a/26070946
func Check2(s1 string, s2 string) bool {
	var (
		p1 int64 = 1
		p2 int64 = 1
	)

	for _, r := range s1 {
		if unicode.IsLetter(r) {
			k := unicode.ToLower(r)
			p1 *= prime(k)
		}
	}

	for _, r := range s2 {
		if unicode.IsLetter(r) {
			k := unicode.ToLower(r)
			p2 *= prime(k)
		}
	}

	return p1 == p2
}

func prime(input rune) int64 {
	switch input {
	case 97:
		return 2
	case 98:
		return 3
	case 99:
		return 5
	case 100:
		return 7
	case 101:
		return 11
	case 102:
		return 13
	case 103:
		return 17
	case 104:
		return 19
	case 105:
		return 23
	case 106:
		return 29
	case 107:
		return 31
	case 108:
		return 37
	case 109:
		return 41
	case 110:
		return 43
	case 111:
		return 47
	case 112:
		return 53
	case 113:
		return 59
	case 114:
		return 61
	case 115:
		return 67
	case 116:
		return 71
	case 117:
		return 73
	case 118:
		return 79
	case 119:
		return 83
	case 120:
		return 89
	case 121:
		return 97
	case 122:
		return 101
	default:
		panic("unexpected input")
	}
	return -1
}
