package leetcode

func FindDuplicate(nums []int) int {
	numMap := make(map[int]bool, len(nums))
	var answer int
	for i := 0; i < len(nums); i++ {
		if numMap[nums[i]] {
			answer = nums[i]
		}
		numMap[nums[i]] = true
	}
	return answer
}