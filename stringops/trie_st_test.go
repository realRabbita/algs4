package stringops

import (
	"testing"
)

func TestTrieST_Put_Delete(t *testing.T) {
	var tests = []struct {
		k string
		v int
	}{
		{"one", 1},
		{"two", 2},
		{"three", 3},
		{"four", 4},
		{"five", 5},
		{"six", 6},
		{"seven", 7},
		{"eight", 8},
		{"nine", 9},
		{"oneee", 111},
		{"twooo", 222},
		{"threeee", 333},
	}

	trie := NewTrie()
	n := len(tests)
	for i := 0; i < n; i++ {
		trie.Put(tests[i].k, tests[i].v)
	}
	for _, test := range tests {
		if got := trie.Get(test.k); got.(int) != test.v {
			t.Errorf("trie.Get(%v) = %v, should be %v", test.k, got, test.v)
		}
	}
	N := trie.Size()
	for i := n - 1; i >= 0; i-- {
		trie.Delete(tests[i].k)
		N--
		if trie.Size() != N {
			t.Errorf("Size = %v, Expected = %v", trie.Size(), N)
		}
	}
}
