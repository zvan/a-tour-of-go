package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	cl := make(chan int)
	cr := make(chan int)

	if t.Left != nil {
		go Walk(t.Left, cl)
		
		for x:=range cl{
			ch<-x
		}
	}

	if t.Right != nil {
		go Walk(t.Right, cr)
		for x:=range cr{
			ch<-x
		}
	}
	
	ch <- t.Value

	close(ch)
}

// // Same determines whether the trees
// // t1 and t2 contain the same values.
// func Same(t1, t2 *tree.Tree) bool

func main() {
	ch := make(chan int)

	t:=tree.New(1)

	go Walk(t, ch)
	for i:= range ch{
		fmt.Println(i);
	}
}
