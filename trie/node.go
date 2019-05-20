package trie

// Node to a char and its children
type Node struct {
	isLeaf bool
	nodes  map[string]*Node
}

func (node *Node) getNode(char string) *Node {
	return node.nodes[char]
}

// NewNode create new node
func NewNode() *Node {
	return &Node{isLeaf: false, nodes: make(map[string]*Node)}
}

func (node *Node) addNewChildNode(value string) {
	node.nodes[value] = NewNode()
}

func (node *Node) isEndOfWord() bool {
	return node.isLeaf
}

func (node *Node) markAsEndOfWord() {
	node.isLeaf = true
}

func (node *Node) hasChildren() bool {
	return len(node.nodes) > 0
}
