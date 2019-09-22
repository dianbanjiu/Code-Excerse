package leetcode

func MoveZeroes(nums []int) {
	//0's count in nums
	zero := 0
	j := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]==0 {
			zero++
		}else {
				nums[j] = nums[i]
				j++
		}
	}

	for i := 0; i < zero; i++ {
		nums[len(nums)-i-1] = 0
	}
}