package heap

type Heap []int // 堆中的元素值即优先队列中的优先度

func (h *Heap) Init() {
	n := len(*h)
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i, n)
	}
}

func (h *Heap) Push(v int) {
	*h = append(*h, v)
	h.up(len(*h) - 1)
}

func (h *Heap) Pop() int {
	hs := *h
	n := len(hs) - 1
	min := hs[0]
	hs[0], hs[n] = hs[n], hs[0]
	h.down(0, n)
	*h = hs[:n]
	return min
}

// 从位置 i 将元素上滤
func (h *Heap) up(i int) {
	for {
		parent := (i - 1) / 2
		if parent == i || h.isLess(parent, i) {
			break
		}
		h.swap(parent, i) // 保证插入后依旧是最小堆
		i = parent
	}
}

// 将位置 mid 下滤交换
func (h *Heap) down(mid, n int) {
	for {
		l := 2*mid + 1
		if l >= n || l < 0 { // l < 0 防止 int 上溢
			break
		}

		min := l
		if r := l + 1; r < n && h.isLess(r, l) {
			min = r
		}

		if !h.isLess(min, mid) { // 可提前退出：从 n/2-1 开始 down，能保证从最后一棵子树开始就是最小堆
			break
		}

		h.swap(mid, min)
		mid = min
	}
}

func (h *Heap) isLess(x, y int) bool {
	return (*h)[x] < (*h)[y]
}

func (h *Heap) swap(i, j int) {
	(*h)[i], (*h)[j] = (*h)[j], (*h)[i]
}
