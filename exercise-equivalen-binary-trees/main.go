package main

import (
	"fmt"
	"sort"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	cl := make(chan int)
	cr := make(chan int)

	if t.Left != nil {
		go Walk(t.Left, cl)

		for x := range cl {
			ch <- x
		}
	}

	if t.Right != nil {
		go Walk(t.Right, cr)
		for x := range cr {
			ch <- x
		}
	}

	ch <- t.Value

	close(ch)
}

func values(t *tree.Tree) []int {
	ch := make(chan int)
	var s []int

	go Walk(t, ch)
	for x := range ch {
		s = append(s, x)
	}

	sort.Sort(sort.IntSlice(s))

	return s
}

// // Same determines whether the trees
// // t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	v1, v2 := values(t1), values(t2)

	if len(v1) != len(v2) {
		return false
	}

	for i := range v1 {
		if v1[i] != v2[i] {
			return false
		}

	}

	return true
}

func main() {
	fmt.Println("Two treas of 1 are equal: ", Same(tree.New(1), tree.New(1)))
	fmt.Println("Two treas of 2 are equal: ", Same(tree.New(2), tree.New(2)))
	fmt.Println("Tree of 1 and tree of 2 are NOT equal: ", !Same(tree.New(1), tree.New(2)))
}
