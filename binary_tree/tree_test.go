package binary_tree

import (
	"fmt"
	"testing"
)

var tree BinaryTree

func init() {
	tree.root = &Node{value: 6}
	for _, n := range []int{4, 8, 2, 5, 7, 9, 1, 3, 10} {
		tree.Insert(n)
	}
}

var f = func(v int) {
	fmt.Printf("%d ", v)
}

// go test -v .
func TestPreOrderTraverse(t *testing.T) {
	tree.PreOrderTraverse(tree.root, f) // 6 4 2 1 3 8 5 7 9 10
}

func TestPostOrderTraverse(t *testing.T) {
	tree.PostOrderTraverse(tree.root, f) // 1 3 2 5 4 7 10 9 8 6
}

func TestInOrderTraverse(t *testing.T) {
	tree.InOrderTraverse(tree.root, f) // 1 2 3 4 5 6 7 8 9 10
}

func TestSearch(t *testing.T) {
	fmt.Printf("%+v\n", tree.Search(9))  // &{value:9 left:<nil> right:0xc42000a1c0}
	fmt.Printf("%+v\n", tree.Search(12)) // <nil>
}
