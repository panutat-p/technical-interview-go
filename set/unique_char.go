package set

func UniqueChar(s string) bool {
	set := make(map[string]struct{})
	for _, v := range s {
		s := string(v)
		if _, ok := set[s]; ok {
			return false
		} else {
			set[s] = struct{}{}
		}
	}
	return true
}
