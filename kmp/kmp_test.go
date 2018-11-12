package kmp

import (
	"fmt"
	"testing"
)

var (
	source  string
	pattern string
)

func init() {
	n := 2
	s := "abxabxabcaby"
	//			______
	for n > 0 {
		source += s
		n--
	}
	pattern = "abcaby"
}

// go test -v -test.run TestKMP
func TestKMP(t *testing.T) {
	fmt.Println(KMP(source, pattern))
	fmt.Println(ForceTraverse(source, pattern))
}

// go test -v -bench=.
func BenchmarkForceTraverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ForceTraverse(source, pattern)
	}
}

func BenchmarkKMP(b *testing.B) {
	for i := 0; i < b.N; i++ {
		KMP(source, pattern)
	}
}
