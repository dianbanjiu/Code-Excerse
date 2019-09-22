package leetcode

func CountPrimes(n int) int {
	var lens int
	if n == 0 || n == 1 {
		return 0
	}
	for i := 2; i < n; i++ {
		if judge(i) {
			lens++
		}
	}
	return lens
}
func judge(num int) bool {
	for i := 2; i*i <= num; i++ {
		if num%i == 0 {
			return false
		}
	}
	return true
}
