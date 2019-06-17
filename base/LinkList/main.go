package main

// 创建链表节点
type Node struct {
	Val  int
	Next *Node
}

//初始化链表
func (linkList *Node)Init(){
	linkList.Val = 0
	linkList.Next = nil
}
//单链表的读取
func GetElem(linkList *Node, i int) (e *int,err bool){
	var p *Node
	p = linkList.Next

	j := 1
	for p != nil && j < i {
		p = p.Next
		j++
	}
	err = true
	if p == nil || j > i {
		err = false
	} 
	e = &p.Val
	return

}

//单链表的插入
func ListInsert(linkList *Node, i, ele int) (err bool){
	var p *Node
	p = linkList.Next
	j := 1

	for p != nil && j < i {
		p = p.Next
		j++
	}
	
	err = true
	if p == nil || j > i{
		err = false
	}

	var s *Node
	s.Init();
	s.Val = ele
	s.Next = p.Next
	p.Next = s
	return 
}

//删除


func main(){

}