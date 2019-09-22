package leetcode

func Rotate(nums []int, k int){
	k %= len(nums)
	if k != 0 {
		nums = append(nums[k+1:], nums[:k+1]...)
	}
}