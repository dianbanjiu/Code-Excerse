package Boring

import "fmt"

type Node struct {
	Val int
	Next *Node
}

func New(v int) *Node{
	return &Node{
		Val:  v,
		Next: nil,
	}
}

func (node *Node)Append(v int) *Node{
	if node == nil {
		node.Val = v
		return node
	}

	p := node
	for p.Next != nil {
		p = p.Next
	}
	p.Next = &Node{
		Val:  v,
		Next: nil,
	}
	return p.Next
}

func PrintList(node *Node){
	p := node
	for p != nil {
		fmt.Print(p.Val, " -> ")
		p = p.Next
	}
}

