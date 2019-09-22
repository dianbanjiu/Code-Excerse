package leetcode

func ContainsDuplicate(nums []int) bool {
	count := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := count[nums[i]]; !ok {
			count[nums[i]] = 1
		} else {
			return true
		}
	}
	return false
}
