package skip_list

import (
	"math/rand"
	"time"
)

type Node struct {
	key   int
	val   interface{}
	nexts []*Node // 向后的节点值
}

type SkipList struct {
	head  *Node      // 头结点
	tail  *Node      // 尾结点（始终是 nil）
	level int        // 层数
	size  int        // 节点总数
	rand  *rand.Rand // 生成随机层数的生成器
}

func NewSkipList(level int) *SkipList {
	if level <= 0 {
		level = 1
	}

	l := &SkipList{
		head: &Node{nexts: make([]*Node, level)},
		tail: &Node{},
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
	for i := range l.head.nexts {
		l.head.nexts[i] = l.tail
	}
	return l
}

// 参考 https://en.wikipedia.org/wiki/File:Skip_list_add_element-en.gif
// 插入节点
func (l *SkipList) insert(k int, v interface{}) {
	level := l.randLevel()               // 将新节点随机插入到某一层，能使跳跃表大体上分布均匀
	prevs := make([]*Node, level, level) // 新节点的 prev 数组
	cur := l.head
	for i := level - 1; i >= 0; i-- { // 逐层向下（其实存储上是从后向前）搜索
		for {
			next := cur.nexts[i] // cur->next
			if next.key == k {
				next.val = v
				return
			}
			if next == l.tail || k < next.key { // 向后找第一个比 k 大的节点 next，取其前驱节点 cur 使 cur->newNode
				prevs[i] = cur
				break
			}
			cur = next // 本层继续向后找
		}
	}

	newNode := &Node{key: k, val: v, nexts: make([]*Node, level, level)}
	for i, prev := range prevs {
		newNode.nexts[i] = prev.nexts[i] // 完成每一层的节点插入
		prev.nexts[i] = newNode
	}
	l.size++
}

// 删除节点
func (l *SkipList) remove(k int) bool {
	cur := l.head
	var target *Node
	rmNodes := make([]*Node, l.level, l.level)

	for level := len(cur.nexts) - 1; level >= 0; level-- {

		for {
			next := cur.nexts[level]
			if next == l.tail || k < next.key { // 向后的节点值更大，无需向后再查找
				break
			}

			if next.key == k { // bingo
				rmNodes[level] = next // 当前层要删除的节点
				target = next
				break
			}

			cur = next
		}
	}

	// 删除每一层的所有该节点
	if target != nil {
		for level, node := range rmNodes {
			if node != nil { // 执行删除
				node.nexts[level] = target.nexts[level]
			}
		}
		l.size--
		return true
	}
	return false
}

// 搜索节点
func (l *SkipList) search(k int) interface{} {
	cur := l.head

	for level := len(l.head.nexts) - 1; level >= 0; level-- { // 从最高层向下搜索

		for {
			next := cur.nexts[level]
			if next == l.tail || k < next.key { // k 肯定不在本层，跳过进入下一层
				break
			}

			if next.key == k {
				return next.val
			}

			cur = next // 本层继续向后搜索
		}
	}

	return nil
}

func (l *SkipList) length() int {
	return l.size
}

func (l *SkipList) randLevel() int {
	i := 1
	for i < l.level && l.rand.Int()%2 == 1 {
		i++
	}
	return i // 插入元素选择 50% 几率分布的随机 level
}
