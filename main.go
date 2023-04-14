package main

import (
	"fmt"
	"github.com/panutat-p/technical-interview-go/pkg"
)

func main() {
	m := pkg.NewCharMap()
	m.Add("a")
	m.Add("b")
	fmt.Println(m.Has("a"))
	fmt.Println(m.Has("z"))
}
