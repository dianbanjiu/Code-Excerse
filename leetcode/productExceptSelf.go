package leetcode

func ProductExceptSelf(nums []int) []int {
	if len(nums) == 1 {
		return nums
	}
	l := make([]int, len(nums))
	r := make([]int, len(nums))
	answer := make([]int, len(nums))

	l[0] = 1
	for i := 1; i < len(nums); i++ {
		l[i] = l[i-1]*nums[i-1]
	}
	r[len(nums)-1] = 1
	for j := len(nums) - 2; j >= 0; j-- {
		r[j] = r[j+1]*nums[j+1]
	}

	for i := 0; i < len(nums); i++ {
		answer[i] = r[i] * l[i]
	}
	return answer
}
