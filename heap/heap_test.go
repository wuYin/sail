package heap

import (
	"fmt"
	"testing"
)

func TestHeapify(t *testing.T) {
	h := &Heap{8, 5, 7, 4, 1, 6, 3, 2}
	h.Init()
	fmt.Println(h) // [1 2 3 4 5 6 7 8] // max heap 转为了 min heap
}

func TestOperation(t *testing.T) {
	h := &Heap{7, 5, 2, 3, 6, 4, 1}
	h.Init()
	fmt.Println(*h)      // [1 3 2 5 6 4 7] // min heap
	fmt.Println(h.Pop()) // 1
	fmt.Println(*h)      // [2 3 4 5 6 7] 	// still min heap
}
