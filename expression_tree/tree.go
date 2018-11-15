package expression_tree

import "fmt"

type Node struct {
	value       rune
	left, right *Node
}

type ExpTree struct {
	root *Node
}

func isOperator(c rune) bool {
	for _, r := range []rune{'+', '-', '*', '/'} {
		if r == c {
			return true
		}
	}
	return false
}

func Express(postfix string) *Node {
	runes := []rune(postfix)
	var stack []*Node

	for _, r := range runes {
		if !isOperator(r) {
			newNode := &Node{value: r}
			stack = append(stack, newNode)
			continue
		}

		// 遇到操作符，创建三角子树
		rNum := stack[len(stack)-1]  // 注意先 pop 的是右操作数
		stack = stack[:len(stack)-1] // stack: [left, right] -> left, op, right	// right 先弹出
		lNum := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		opNode := &Node{
			value: r,
			left:  lNum,
			right: rNum,
		}
		stack = append(stack, opNode)
	}
	return stack[0]
}

// 只有二元运算符，放心地按层遍历
func (tree *ExpTree) bfsTraverse() {
	if tree.root == nil {
		return
	}
	q := []*Node{tree.root}
	for len(q) > 0 {
		cur := q[0]
		q = q[1:]
		fmt.Printf("%c ", cur.value)
		if cur.left != nil {
			q = append(q, cur.left)
		}
		if cur.right != nil {
			q = append(q, cur.right)
		}
	}
}
