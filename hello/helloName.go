package main

import (
	"fmt"

	"github.com/dianbanjiu/leetcode"
)

func main() {
	num := []int{1, 2, 2, 3, 4, 5, 2}
	val := 2

	fmt.Println(leetcode.RemoveElement(num, val))
}
