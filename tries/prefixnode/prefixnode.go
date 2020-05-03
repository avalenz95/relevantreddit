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
func (node PrefixNode) NumChildren() int {
	return len(node.Children)
}

//returns number of users associated with a node
func (node PrefixNode) NumUsers() int {
	return len(node.Users)
}

//find child in parent node
func (node PrefixNode) HasChild(child *PrefixNode) bool {

	_, found := node.Children[child.Char]

	if found {
		return true
	}
	return false
}

//Add node
func (node PrefixNode) AddChild(child *PrefixNode) bool {

	if node.HasChild(child) {
		return false
	}
	//Add Node
	node.Children[child.Char] = child

	return true
}
