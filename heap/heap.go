package heap

type Heap []int

func (h *Heap) init() {
	n := len(*h)
	for i := n/2 - 1; i >= 0; i-- {
		h.down(i, n)
	}
}

func (h *Heap) down(mid, n int) {
	for {
		l := 2*mid + 1
		if l >= n || l < 0 {
			break
		}

		min := l
		if r := l + 1; r < n && h.isLess(r, l) {
			min = r
		}

		if !h.isLess(min, mid) {
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
