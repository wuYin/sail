package heap

import (
	"fmt"
	"testing"
)

func TestHeapify(t *testing.T) {
	h := &Heap{3, 1, 2}
	h.init()
	fmt.Println(h)
}
