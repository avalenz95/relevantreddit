package prefixnode

//PrefixNode in the trie
type PrefixNode struct {
	Char     rune
	Terminal bool
	Children map[rune]*PrefixNode
	Users    []string
}

//New creates a node with character and terminal value
func New(char rune) PrefixNode {

	node := PrefixNode{char, false, make(map[rune]*PrefixNode), make([]string, 0, 0)}

	return node
}

//checks to see if  current node is terminal
func (node PrefixNode) isTerminal() bool {
	return node.Terminal
}

//returns number of child nodes
func (node PrefixNode) numChildren() int {
	return len(node.Children)
}

//returns number of users associated with a node
func (node PrefixNode) numUsers() int {
	return len(node.Users)
}
