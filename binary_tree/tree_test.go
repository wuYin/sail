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

func TestBreadthFirstTraverse(t *testing.T) {
	tree.BreadthFirstTraverse(f) // 6 4 8 2 5 7 9 1 3 10
}

func TestDepth(t *testing.T) {
	fmt.Println(tree.Depth(tree.root)) // 4
	tree.Insert(11)
	fmt.Println(tree.Depth(tree.root)) // 5
}

func TestCopy(t *testing.T) {
	var newTree BinaryTree
	newTree.root = tree.Copy(tree.root)
	newTree.BreadthFirstTraverse(f) // 6 4 8 2 5 7 9 1 3 10
}

func TestMirror(t *testing.T) {
	var newTree BinaryTree
	newTree.root = tree.Mirror(tree.root)
	newTree.BreadthFirstTraverse(f) // 6 8 4 7 9 2 5 10 1 3
}

func TestNodeCount(t *testing.T) {
	fmt.Println(tree.NodeCount(tree.root)) // 10
}

func TestLeafCount(t *testing.T) {
	fmt.Println(tree.LeafCount(tree.root)) // 5
}

func TestSame(t *testing.T) {
	var newTree BinaryTree
	newTree.root = tree.Copy(tree.root)
	fmt.Println(tree.Same(tree.root, newTree.root)) // true
	newTree.root.right.value = 100
	fmt.Println(tree.Same(tree.root, newTree.root)) // false
}

func TestFree(t *testing.T) {
	tree.Free()
	fmt.Printf("%+v\n", tree.root) // <nil>
}

func TestAllPath(t *testing.T) {
	tree.AllPath(tree.root, nil, f)
}

func TestCurSearch(t *testing.T) {
	fmt.Printf("%+v\n", tree.CurSearch(tree.root, 2))  // &{value:2 left:0xc42000a140 right:0xc42000a160}
	fmt.Printf("%+v\n", tree.CurSearch(tree.root, 12)) // <nil>
}
