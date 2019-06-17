package main

import "fmt"

func main() {
	//复数
	var c complex64 = 3 + 2i
	fmt.Println(c)

	fmt.Println("----字符串----")
	//字符串本身并不可以修改，不过可以通过转换为 []byte 来进行
	s := "hello"
	b := []byte(s)
	b[2] = 'x' // byte 类型的字符需要用单引号，而不是双引号
	s2 := string(b)
	fmt.Printf("Now hello change to %v\n", s2)

	//声明多行字符串，使用反引号来进行
	multiString := `hello
	world`
	fmt.Println(multiString)

	// 创建 map
	numbers := make(map[string]int64)
	numbers["one"] = 1
	numbers["two"] = 2
	numbers["three"] = 3
	fmt.Println(numbers)
	fmt.Println("删除map[1]")
	delete(numbers, "one")
	fmt.Println(numbers)

}
