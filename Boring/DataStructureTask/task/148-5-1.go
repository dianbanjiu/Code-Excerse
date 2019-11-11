package task

import (
	"fmt"
	"strconv"
	"strings"
)

type optr []string
type opnd []string

func (tr optr)Top() string{
	if len(tr) == 0{
		return ""
	}
	return tr[len(tr)-1]
}

func (nd opnd)Top() string{
	if len(nd) == 0{
		return ""
	}
	return nd[len(nd)-1]
}


func (tr *optr)Push(e string){
	*tr = append(*tr, e)
}

func (nd *opnd)Push(e string){
	*nd = append(*nd, e)
}

func (tr *optr) Pop()string{
	if len(*tr)==0{
		return ""
	}
	t := *tr
	r := t[len(t)-1]
	t = t[:len(t)-1]
	*tr = t
	return r
}

func (nd *opnd) Pop()string{
	if len(*nd)==0{
		return ""
	}
	n := *nd
	r := n[len(n)-1]
	n = n[:len(n)-1]
	*nd = n
	return r
}

func EvaluateExpre(s string){
	s = TurntoTF(s)
	var(
		optr optr
		opnd opnd
	)
	op := "&|~()#"
	optr.Push("#")
	i := 0
	c := string(s[i])
	for c!="#"||optr.Top()!="#"{
		c := string(s[i])
		if !strings.Contains(op, c){
			opnd.Push(c)
			i++
			continue
		}
		switch (Precedeexp(optr.Top(), c)) {
		case -1:
			optr.Push(c)
			i++
			break
		case 0:
			optr.Pop()
			i++
			break
		case 1:
			theta := optr.Pop()
			a := opnd.Pop()
			var b string
			if theta=="~"{
				b = ""
			} else {
				b = opnd.Pop()
			}
			opnd.Push(Operateexp(a, theta, b))
			break
		}
		switch opnd.Top() {
		case "1":
			fmt.Println("True Forever")
			break
		case "0":
			fmt.Println("False forever")
			break
		default:
			fmt.Println("Satisfactible")
			break
		}
	}
}

func TurntoTF(s string) string {
	var r string
	d := make(map[string]int)
	op := "&|~()#"
	println("为 "+s+" 的变元赋值，真为 1， 假为 0")
	for i := 0; i < len(s); i++ {
		c := string(s[i])
		if _, ok := d[c]; !ok && !strings.Contains(op, c) {
			var t int
			fmt.Println("为变元 "+ c+ " 赋值")
			fmt.Scanln(&t)
		}
	}
	for i := 0; i < len(s); i++ {
		c := string(s[i])
		if !strings.Contains(op, c) {
			c = strconv.Itoa(d[c])
		}
		r = r+c
	}
	return r
}

func Operateexp(data1, theta, data2 string) string {
	var r , d1, d2 bool
	d1 = Judge(data1)
	d2 = Judge(data2)
	switch theta {
	case "&":
		r = d1&&d2
		break
	case "|":
		r = d1||d2
		break
	case "~":
		r = !d1
	}
	if r {
		return "1"
	}
	return "0"
}

func Judge(s string) bool {
	if s == "1" {
		return true
	}
	return false
}

func Precedeexp(top string, s string) int{
	l1 := "~("
	l3 := "&|~("
	t1 := "&|"
	t2 := "(#"
	if (strings.Contains(t1, top)&&strings.Contains(l1,s))||(top=="~"&&s=="(")||(strings.Contains(t2, top)&&strings.Contains(l3, s)){
		return -1
	} else if (top == "(" && s == ")") || (top == "#" && s == "#") {
		return 0
	}
	return 1
}