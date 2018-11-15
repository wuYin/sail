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
// 使用递归实现 DFS 深度优先遍历
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
	if node == nil {
		return
	}
	if node.left != nil {
		tree.InOrderTraverse(node.left, f)
	}
	f(node.value)
	if node.right != nil {
		tree.InOrderTraverse(node.right, f)
	}
}

// BFS
// 使用队列实现 BFS 广度(横向)优先遍历
func (tree *BinaryTree) BreadthFirstTraverse(f func(int)) {
	if tree.root == nil {
		return
	}

	q := make([]*Node, 0)
	q = append(q, tree.root) // root 入节点队列
	for len(q) > 0 {
		node := q[0]
		q = q[1:]
		f(node.value)
		if node.left != nil {
			q = append(q, node.left)
		}
		if node.right != nil {
			q = append(q, node.right)
		}
	}
}

// 搜索
// 有种有序数组二分查找元素的感觉
func (tree *BinaryTree) Search(target int) (*Node) {
	if tree.root == nil {
		return nil
	}
	cur := tree.root
	for cur != nil {
		switch {
		case cur.value > target:
			cur = cur.left
		case cur.value < target:
			cur = cur.right
		default:
			return cur
		}
	}
	return nil
}

// 计算树的深度
// Node 的深度是 root 到 Node 的唯一路径长
func (tree *BinaryTree) Depth(node *Node) int {
	if node == nil {
		return 0
	}

	lDepth := tree.Depth(node.left)
	rDepth := tree.Depth(node.right)

	if lDepth > rDepth { // 回溯比较左右子树的深度
		return lDepth + 1 // 算上 node 本身
	}
	return rDepth + 1
}

// Node 节点数
func (tree *BinaryTree) NodeCount(node *Node) int {
	if node == nil {
		return 0
	}
	return 1 + tree.NodeCount(node.left) + tree.NodeCount(node.right)
}

// Leaf 节点数
func (tree *BinaryTree) LeafCount(node *Node) int {
	if node == nil {
		return 0
	}
	if node.left == nil && node.right == nil {
		return 1
	}
	return tree.LeafCount(node.left) + tree.LeafCount(node.right)
}

// 原样复制
func (tree *BinaryTree) Copy(node *Node) *Node {
	if node == nil {
		return nil
	}
	newNode := &Node{value: node.value}
	newNode.left = tree.Copy(node.left)
	newNode.right = tree.Copy(node.right)
	return newNode
}

// 镜像复制
func (tree *BinaryTree) Mirror(node *Node) *Node {
	if node == nil {
		return nil
	}
	newNode := &Node{value: node.value}
	newNode.left = node.right
	newNode.right = node.left
	return newNode
}

// 判等
// 所有子节点的位置和值均相等的两棵树
func (tree *BinaryTree) Same(node1 *Node, node2 *Node) bool {
	if node1 == nil && node2 == nil {
		return true
	}
	if node1 == nil || node2 == nil {
		return false
	}
	return node1.value == node2.value && tree.Same(node1.left, node2.left) && tree.Same(node1.right, node2.right)
}

// 删除整棵树
func (tree *BinaryTree) Free() {
	tree.root = nil // GC 会处理
}
