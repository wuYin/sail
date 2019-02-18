package single_list

import (
	"fmt"
	"testing"
)

func TestList(t *testing.T) {
	l := NewList()
	l.Append(1)
	l.Append(2)
	l.Prepend(3)
	fmt.Println(l) // 3->1->2->nil

	l.Remove(2)
	fmt.Println(l) // 3->1->nil
}
