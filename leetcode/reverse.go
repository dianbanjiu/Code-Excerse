package leetcode

func Reverse(x int) int {
	var (
		r int
		i int
	)
	if x < 0 {
		i = 0
		x = -x
	} else {
		i = 1
	}
	var s []int
	for {
		e := x % 10
		x = x / 10
		s = append(s, e)
		if x == 0 {
			break
		}
	}

	for j := 0; j < len(s); j++ {
		r = r*10 + s[i]
	}

	if i == 0 {
		return -r
	}

	return r
}
