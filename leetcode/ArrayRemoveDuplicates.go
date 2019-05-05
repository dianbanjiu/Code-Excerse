package leetcode

// 给定一个排好序的数组，去掉其中重复出现的元素，使得每个元素在数组中仅出现一次，返回新数组的长度
func ArrayRemoveDuplicates(nums []int) int{
	i := len(nums)
	for j := 0; j < i; j++ {
		if nums[j] == nums[j+1] {
			i--
		}
	}
	return i
}