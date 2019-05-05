package main

import "fmt"

func main(){
	a := 1
	b := returnTest(a)
	fmt.Println(b)
}

func returnTest(a int) int {
	if a == 0 {
		return a
	} else {
		return 1
	}
}