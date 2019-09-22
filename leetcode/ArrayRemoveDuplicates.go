package leetcode

// 给定一个排好序的数组，去掉其中重复出现的元素，使得每个元素在数组中仅出现一次，返回新数组的长度
func ArrayRemoveDuplicates(nums []int) int {

	/*
		for i := 0; i < len(nums)-1; i++ {
			for j := i + 1; j < len(nums); j++ {
				if nums[j] == nums[i] {
					nums = append(nums[:j], nums[j+1:]...)
					j = i
				}

			}
		}
	*/
	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] == nums[i] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}

	return len(nums)
}
