package main

import (
	"fmt"
)

func main() {
	str1 := "1223"
	str2 := "1224"

	if str1 > str2 {
		fmt.Println("str1 is greater than str2")
	} else if str1 < str2 {
		fmt.Println("str1 is less than str2")
	} else {
		fmt.Println("str1 is equal to str2")
	}
}
