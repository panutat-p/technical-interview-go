package pkg

func Pop[T comparable](sl []T, idx int) []T {
	var (
		ret = make([]T, len(sl)-1, cap(sl)-1)
	)
	copy(ret, sl[:idx])         // copy first half
	copy(ret[idx:], sl[idx+1:]) // copy after the index to the end
	return ret
}

type Slice []int

func (s *Slice) Append(e int) {
	*s = append(*s, e)
}

func (s *Slice) Prepend(e int) {
	*s = append([]int{e}, *s...)
}

func (s *Slice) Insert(e, idx int) {
	left := (*s)[:idx]
	right := (*s)[idx:]
	*s = append(left, e)
	*s = append(*s, right...)
}

func (s *Slice) Pop(idx int) int {
	ret := (*s)[idx]
	*s = append((*s)[:idx], (*s)[idx+1:]...)
	return ret // same backing longest_parking_roof
}
