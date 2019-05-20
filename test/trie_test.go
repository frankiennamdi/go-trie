package test

import (
	"fmt"
	"testing"

	"github.com/frankiennamdi/go-trie/trie"
	"github.com/stretchr/testify/assert"
)

var words = []struct {
	word     string
	isPrefix bool
	isWord   bool
}{
	{"hello", false, true},
	{"hey", false, true},
	{"john", false, true},
	{"he", true, false},
	{"xyz", false, false},
}

func TestCanFindWord(t *testing.T) {
	trie := trie.NewTree()
	trie.Insert("hello")
	trie.Insert("hey")
	trie.Insert("john")

	for _, tc := range words {
		t.Run(tc.word, func(t *testing.T) {
			result := trie.Search(tc.word)
			fmt.Printf("%+v\n", result)
			assert.Equal(t, result.IsPrefix, tc.isPrefix)
			assert.Equal(t, result.IsWord, tc.isWord)
		})
	}
}
