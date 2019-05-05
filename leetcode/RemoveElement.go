package leetcode

func RemoveElement(num []int, val int) int {

	// 使用 i 来对非 val 值进行计数
	// 如果需要返回的是去掉 val 值的 num，可以使用 num[:i]
	var i int
	for j := 0; j < len(num); j++ {
		if num[j] != val {
			num[i] = num[j]
			i++
		}
	}
	return i
}
