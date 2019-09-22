package leetcode

func FlipAndInvertImage(A [][]int) [][]int {
	count := len(A[0])

	for _, row := range A {
		for i := 0; i < count/2; i++ {
			tmp := row[i]
			row[i] = row[count-1-i]
			row[count-1-i] = tmp
		}
	}

	for i := 0; i < len(A); i++ {
		for j := 0; j < count; j++ {
			if A[i][j] == 0 {
				A[i][j] = 1
			} else {
				A[i][j] = 0
			}
		}
	}

	return A

}
