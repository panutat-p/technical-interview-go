package main

import (
	"fmt"
	"github.com/panutat-p/technical-interview-go/pkg"
)

func main() {
	fmt.Println("🟩 want -2, got", pkg.Min(-2, 8))
	fmt.Println("🟩 want 10, got", pkg.Max(-1, 10))
}
