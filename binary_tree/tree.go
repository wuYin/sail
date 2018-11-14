package binary_tree

type Node struct {
	value       int
	left, right *Node
}

// 定义完全二叉树
type BinaryTree struct {
	root *Node
}

// 有序插入
func (tree *BinaryTree) Insert(v int) {
	if tree == nil {
		tree.root = &Node{value: v}
		return
	}
	cur := tree.root
	for {
		switch {
		case v < cur.value: // 在左子树上
			if cur.left == nil {
				cur.left = &Node{value: v}
				return
			}
			cur = cur.left
		case v > cur.value: // 在右子树上
			if cur.right == nil {
				cur.right = &Node{value: v}
				return
			}
			cur = cur.right
		default: // 待定
			return
		}
	}
}

// 先序遍历
// parent 总在 child 前被遍历
// 时间空间复杂度均为 O(N)
func (tree *BinaryTree) PreOrderTraverse(node *Node, f func(int)) {
	if node == nil {
		return
	}
	f(node.value) // 处理父节点的值
	if node.left != nil {
		tree.PreOrderTraverse(node.left, f) // 一股脑地遍历左子树
	}
	if node.right != nil {
		tree.PreOrderTraverse(node.right, f) // 回溯时遍历右子树
	}
}

// 后序遍历
// child 总在 parent 前被遍历，left 先于 right，即深度优先
func (tree *BinaryTree) PostOrderTraverse(node *Node, f func(int)) {
	if node == nil {
		return
	}
	if node.left != nil {
		tree.PostOrderTraverse(node.left, f) // 找到最深左子节点
	}
	if node.right != nil {
		tree.PostOrderTraverse(node.right, f) // 回溯时找最深右子节点
	}
	f(node.value)
}

// 中序遍历
// 最顺其自然的遍历方式，left->parent->right
func (tree *BinaryTree) InOrderTraverse(node *Node, f func(int)) {
	if node.left != nil {
		tree.InOrderTraverse(node.left, f)
	}
	f(node.value)
	if node.right != nil {
		tree.InOrderTraverse(node.right, f)
	}
}
