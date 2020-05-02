package prefixnode

//prefix node in the trie
type prefixNode struct {
	char     string
	terminal bool
	children map[string]*prefixNode
	users []string
}

//New creates a node with character and terminal value
func New(char string, terminal bool) prefixNode {

	node := prefixNode{char, terminal, map[string]*prefixNode{}, []string}

	return node
}

//
func (node prefixNode) isTerminal() {
	return node.terminal
}