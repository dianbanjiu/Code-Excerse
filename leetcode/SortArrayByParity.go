package leetcode

func SortArrayByParity(A []int) []int {
	//奇偶校验。
	i := 0
	j := len(A)-1

	for i < j {
		if A[i]%2 > A[j]%2 {
			var temp int
			temp = A[i]
			A[i] = A[j]
			A[j] = temp
		}
		if A[i]%2 == 0 {
			i++
		}
		if A[j]%2 == 1{
			j--
		}
	}

	return A
}