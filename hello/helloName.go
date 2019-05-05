package main

import (
	"fmt"
	"github.com/dianbanjiu/CodeExcerse/Boring"
)

func main() {
	var count int
	fmt.Println(" Please input the count number(>2): ")
	fmt.Scanf("%d", &count)

	fmt.Println("Print fibonacci list now: ")
	for _, item := range Boring.Fibonacci(count){
		fmt.Println(item)
	}

}
