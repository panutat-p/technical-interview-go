package main

import "fmt"

func main() {
	sl := []int{100}
	fmt.Println("sl:", sl)
	r := sl[1:]
	fmt.Println("r:", r)
}
