package expression_tree

import "testing"

func TestExpress(t *testing.T) {
	postfix := "ab+cde+**"
	expTree := &ExpTree{}
	expTree.root = Express(postfix) // * + * a b c + d e
	expTree.bfsTraverse()
}
