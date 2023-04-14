package pkg

import (
	"fmt"
	"strings"
)

type CharMap map[string]int

func NewCharMap() CharMap {
	return CharMap{}
}

func (m CharMap) Has(str string) bool {
	_, ok := m[str]
	return ok
}

func (m CharMap) Add(str string) {
	m[str] += 1
}

func (m CharMap) Keys() []string {
	var (
		keys []string
	)
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

func (m CharMap) Print() {
	var sb strings.Builder
	for k, v := range m {
		sb.WriteString(fmt.Sprintf("%v:%v ", k, v))
	}
	fmt.Println(sb.String())
}
