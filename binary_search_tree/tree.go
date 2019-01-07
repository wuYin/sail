package bst

import "fmt"

type Node struct {
	value       int
	left, right *Node
}

// Binary Search Tree
type BST struct {
	root *Node
}

// Insert
func (tree *BST) Insert(v int) {
	// 在处理二叉树的操作时候，注意检查 root、node 为空等退化情况
	if tree.root == nil {
		tree.root = &Node{value: v}
		return
	}
	cur := tree.root
	for cur != nil {
		switch {
		case cur.value < v:
			if cur.right == nil {
				cur.right = &Node{value: v}
				return
			}
			cur = cur.right
		case cur.value > v:
			if cur.left == nil {
				cur.left = &Node{value: v}
				return
			}
			cur = cur.left
		default:
			return // 假设元素值互异 // do nothing 或更新
		}
	}
}

// Delete
func (tree *BST) Delete(node *Node, v int) *Node {
	if node == nil {
		return nil
	}

	switch {
	case v > node.value:
		node.right = tree.Delete(node.right, v)
	case v < node.value:
		node.left = tree.Delete(node.left, v)
	default: // 找到了要删除的节点
		// 叶子节点或只有一个 child
		cur := node
		if node.left == nil || node.right == nil {
			if node.left == nil {
				node = node.right
			} else if node.right == nil {
				node = node.left
			}
		} else {
			// 一般情况，删除有 2 个 child 的节点
			// rMinNode := tree.FindMin(node.right) // 把右子树的所有最右子树统一向上挪一位
			rMinNode := tree.FindMax(node.left) // also ok
			node.value = rMinNode.value
			node.right = tree.Delete(node.right, node.value)
		}
		tree.Free(cur)
	}
	return node
}

// Free
func (tree *BST) Free(node *Node) {
	node = nil
}

// Find
func (tree *BST) Find(node *Node, v int) *Node {
	if node == nil {
		return nil
	}

	switch {
	case node.value < v:
		return tree.Find(node.right, v) // 递归查找右子树
	case node.value > v:
		return tree.Find(node.left, v)
	default:
		return node
	}
}

// FindMin
func (tree *BST) FindMin(node *Node) *Node {
	if node == nil {
		return nil
	}
	if node.left == nil {
		return node
	}
	return tree.FindMin(node.left) // 闷头向左走	// bst is good
}

// FindMax
func (tree *BST) FindMax(node *Node) *Node {
	if node == nil {
		return nil
	}
	if node.right == nil {
		return node
	}
	return tree.FindMax(node.right) // 闷头向右走
}

// BFS Traverse
func (tree *BST) BFSTraverse() {
	if tree.root == nil {
		return
	}
	q := []*Node{tree.root}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		fmt.Print(cur.value, " ")
		if cur.left != nil {
			q = append(q, cur.left)
		}
		if cur.right != nil {
			q = append(q, cur.right)
		}
	}
}
