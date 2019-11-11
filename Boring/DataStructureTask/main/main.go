package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main(){
	a :=  []int{1,2,3,4,5,6,7}
	k := 3
	rotate(a,k)
}

func rotate(nums []int, k int)  {
	k = k%len(nums)
	nums = append(nums[len(nums)-k:],nums[:len(nums)-k]...)
	fmt.Println(nums)
}