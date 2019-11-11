package task

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)
// 实现 LISP 加法表达式的求值
// 例：6, (+45)，(+(+25)8) 的都是表达式，其结果分别为 6,9,15
// 要求：从文件中读入表达式，文件的每一行为一个表达式，每读入一行就输出一个表达式的值

func Evaluate(file string){
	fi,err := os.Open(file)
	if err != nil {
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		l, _, err := br.ReadLine()
		if err == io.EOF{
			break
		}

		fmt.Println(Calculate(string(l)+"#"))
	}
}

func Calculate(s string) int{
	op := map[string]int{
		"(": 0,
		"+": 1,
		")": 2,
		"#": 3,
	}
	var (
		optr []string
		opnd []int
	)
	optr = append(optr, "#")


	i := 0
	for string(s[i]) != "#" || optr[len(optr)-1] != "#" {
		c := string(s[i])
		if _, ok := op[c]; !ok{
			d, _ := strconv.Atoi(c)
			opnd = append(opnd, d)
			i++
			continue
		} else {
			switch (Precede(optr[len(optr)-1], c)) {
			case -1:
				optr = append(optr, c)
				i++
				break
			case 0:
				optr = optr[:len(optr)-1]
				i++
				break
			case 1:
				data1 := opnd[len(opnd)-1]
				opnd = opnd[:len(opnd)-1]
				data2 := opnd[len(opnd)-1]
				opnd = opnd[:len(opnd)-1]
				optr = optr[:len(optr)-1]
				opnd = append(opnd, data1+data2)
				break
			}
		}
	}
	return opnd[len(opnd)-1]
}

func Precede(s1 string, s2 string) int{
	if (s1 == "+" && (s2 == "+" || s2 == ")" || s2 == "#")) || (s1 == ")" && (s2 == "+" || s2 == ")" || s2 == "#")) {
		return 1
	} else if (s1 == "(" && s2 == ")") || (s1 == "#" && s2 == "#"){
		return 0
	}
	return -1
}