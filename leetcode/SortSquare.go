package leetcode

import "sort"

func SortedSquares(A []int) []int {
	newA := make([]int, len(A))
	for i, v := range A {
		newA[i] = v*v
	}
	sort.Slice(newA, func(i, j int) bool { return newA[i] < newA[j] })
	return newA
}
