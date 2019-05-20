package test

import (
	"fmt"
	"testing"

	"github.com/frankiennamdi/go-trie/trie"
	"github.com/stretchr/testify/assert"
)

func TestCanFindWord(t *testing.T) {
	trie := trie.NewTree()
	trie.Insert("hello")
	result := trie.Search("hello")
	fmt.Printf("%+v\n", result)
	assert.Equal(t, result.IsPrefix, false, "false expected")
	assert.Equal(t, result.IsWord, true, "true expected")
}

func TestCanFindPrefix(t *testing.T) {
	trie := trie.NewTree()
	trie.Insert("hello")
	result := trie.Search("he")
	fmt.Printf("%+v\n", result)
	assert.Equal(t, result.IsPrefix, true, "true expected")
	assert.Equal(t, result.IsWord, false, "false expected")
}

func TestCannotFindNonExistingWord(t *testing.T) {
	trie := trie.NewTree()
	trie.Insert("hello")
	result := trie.Search("xyz")
	fmt.Printf("%+v\n", result)
	assert.Equal(t, result.IsPrefix, false, "false expected")
	assert.Equal(t, result.IsWord, false, "false expected")
}
