package test

import (
	"fmt"
	"github.com/dianbanjiu/Code-Excerse/Boring"
	"testing"
)

func TestNew(t *testing.T){
	in := &Boring.Node{
		Val:  0,
		Next: nil,
	}

	p := in.Next
	if p == nil {
		fmt.Println("got it")
	}
}
