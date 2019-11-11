package leetcode

func IsValid(s string) bool{
	if s == ""{
		return true
	}
	var optr []string
	optr = append(optr, "#")
	var c string
	for i:=0;i<len(s)||c!= ""; i++{
		c = string(s[i])
		r := optr[len(optr)-1]+c
		if r == "()"|| r=="[]"|| r == "{}" {
			optr = optr[:len(optr)-1]
			continue
		}
		optr = append(optr, c)
	}
	if len(optr) != 1 {
		return false
	}
	return true
}
