package skip_list

import (
	"fmt"
	"testing"
)

func TestInsert(t *testing.T) {
	l := NewSkipList(4)
	l.insert(1, "one")
	l.insert(2, "two")
	l.insert(3, "three")
	fmt.Println(l.length())  // 3
	fmt.Println(l.search(2)) // two
	fmt.Println(l.search(4)) // nil
}
