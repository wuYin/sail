package bst

import (
	"fmt"
	"testing"
)

var tree BST

func TestInsert(t *testing.T) {
	tree.Insert(2)
	tree.Insert(3)
	tree.Insert(1)
	fmt.Printf("%d %d %d\n", tree.root.value, tree.root.left.value, tree.root.right.value) // 2 1 3
}

func TestFind(t *testing.T) {
	for _, v := range []int{4, 1, 2, 3, 5, 0} {
		tree.Insert(v)
	}
	fmt.Printf("%+v\n", tree.Find(tree.root, 2)) // &{value:2 left:<nil> right:0xc42000a0e0}
	fmt.Printf("%+v\n", tree.FindMin(tree.root)) // &{value:0 left:<nil> right:<nil>}
	fmt.Printf("%+v\n", tree.FindMax(tree.root)) // &{value:5 left:<nil> right:<nil>}
}

func TestDelete(t *testing.T) {
	for _, v := range []int{4, 1, 2, 3, 5, 0} {
		tree.Insert(v)
	}
	tree.BFSTraverse() // 4 1 5 0 2 3
	tree.Delete(tree.root, 5)
	fmt.Println()
	tree.BFSTraverse() // 4 1 0 2 3
}
