package leetcode

func MajorityElement(nums []int) int {
	//make a map,key store the nums, value store the key's count
	count := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if _, ok := count[nums[i]]; !ok{
			count[nums[i]] = 1
		} else {
			count[nums[i]]++
		}
	}
	max := nums[0]
	for i, v := range count  {
		if v > count[max] {
			max = i
		}
	}
	return max
}
