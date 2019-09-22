package leetcode
func MissingNumber(nums []int) int {
	lens := len(nums)
	expectedSum := lens*(lens+1)/2
	actualSum := 0
	for _, v := range nums{
		actualSum += v
	}
	return expectedSum-actualSum
}