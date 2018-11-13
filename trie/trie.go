package trie

type TrieNode struct {
	IsWord bool               // 标记到达该节点是否构成单词
	Next   map[rune]*TrieNode // 最多为 26 个的子树分支
}

// 构建 s 的子树
func (root *TrieNode) Insert(s string) {
	runes := []rune(s)
	cur := root
	for _, r := range runes {
		if next, ok := cur.Next[r]; ok {
			cur = next // 已存在子节点
		} else {
			node := TrieNode{IsWord: false}
			if cur.Next == nil {
				cur.Next = make(map[rune]*TrieNode) // 不存在则创建新节点 // 不判断会覆盖map
			}
			cur.Next[r] = &node
			cur = &node
		}
	}
	cur.IsWord = true
}

// 按字符顺序向下搜索
func (root *TrieNode) Search(s string) bool {
	cur := root
	for i := 0; i < len(s) && cur != nil; i++ {
		if next, ok := cur.Next[rune(s[i])]; !ok {
			return false // 不存在的
		} else {
			cur = next
		}
	}
	return cur != nil && cur.IsWord // 找到节点且被标记为单词
}
