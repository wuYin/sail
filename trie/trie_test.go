package trie

import "testing"

// go test -v -test.run Search
func TestTrieNode_Search(t *testing.T) {
	root := &TrieNode{
		IsWord: false,
		Next:   make(map[rune]*TrieNode),
	}
	for _, s := range []string{"we", "all", "are", "warrior"} {
		root.Insert(s)
	}

	var trieTests = []struct {
		input  string
		expect bool
	}{
		{"we", true},
		{"are", true},
		{"arex", false},
		{"al", false},
	}
	for _, unit := range trieTests {
		res := root.Search(unit.input)
		if res != unit.expect {
			t.Fatalf("search: %s, expect: %t, got:%t", unit.input, unit.expect, res)
		}
	}
}
