package sqrt

func Get2(num int) int {
	if num == 1 {
		return 1
	}
	ans := 2
	for i := 2; i < num; i += 1 {
		if i*i <= num {
			ans = i
		} else {
			break
		}
	}
	return ans
}
