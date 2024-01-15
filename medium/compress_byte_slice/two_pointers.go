package compress_byte_slice

import (
	"strconv"
)

// Compress
// 443. String Compression
// https://leetcode.com/problems/string-compression
//
// Run-length encoding (RLE) algorithm
func Compress(input []byte) int {
	// loop all character
	var (
		i   int
		ret []byte
	)
	for i < len(input) {
		// add first character
		count := 1
		ret = append(ret, input[i])
		// loop same characters
		for i < len(input)-1 && input[i] == input[i+1] {
			i += 1
			count += 1
		}
		// add count
		if count > 1 {
			ret = append(ret, Bytes(count)...)
		}
		i += 1
	}
	// modify array
	input = input[:0]
	input = append(input[:0], ret...)
	return len(ret)
}

func Bytes(n int) []byte {
	s := strconv.Itoa(n)
	return []byte(s)
}
