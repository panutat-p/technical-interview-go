package array

import "fmt"

// IsAnagram
// sum of rune value
func IsAnagram(s1 string, s2 string) bool {
	var total1 int32
	var total2 int32
	for _, i32 := range s1 {
		if i32 == 32 { // single space char
			continue
		} else {
			total1 += i32
		}
	}

	for _, i32 := range s2 {
		fmt.Printf("type is %T, value is %v, string is %v\n", i32, i32, string(i32))
		if i32 == 32 {
			continue
		} else {
			total2 += i32
		}
	}
	return total1 == total2
}
