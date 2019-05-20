package trie

import (
	"strings"
)

// Result result of a search
type Result struct {
	IsPrefix bool

	IsWord bool

	SearchString string
}

// Tree structure
type Tree struct {
	root *Node
}

// NewTree create new tree
func NewTree() *Tree {
	return &Tree{
		root: NewNode(),
	}
}

// Insert word
func (tree *Tree) Insert(word string) {
	wordUpper := strings.ToUpper(word)
	curr := tree.root
	for _, runeValue := range wordUpper {
		char := string(runeValue)
		node := curr.getNode(char)

		if node == nil {
			curr.addNewChildNode(char)
		}
		curr = curr.getNode(char)
	}
	curr.markAsEndOfWord()
}

func (tree *Tree) find(word string) *Node {
	wordUpper := strings.ToUpper(word)
	var curr *Node
	curr = tree.root
	for _, runeValue := range wordUpper {
		char := string(runeValue)
		curr = curr.getNode(char)
		if curr == nil {
			return &Node{}
		}
	}
	return curr
}

// Search for word
func (tree *Tree) Search(word string) Result {
	node := tree.find(word)
	return Result{IsPrefix: (node != nil) && node.hasChildren(),
		IsWord:       (node != nil) && node.isEndOfWord(),
		SearchString: word}
}
