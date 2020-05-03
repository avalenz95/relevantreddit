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

//HasChild if node has child with given rune value
func (node PrefixNode) HasChild(char rune) bool {

	_, found := node.Children[char]

	if found {
		return true
	}
	return false
}

//AddChild Node to parents children
func (node PrefixNode) AddChild(child *PrefixNode) bool {

	if node.HasChild(child.Char) {
		return false
	}
	//Add Node
	node.Children[child.Char] = child

	return true
}

//GetChild Returns a pointer to child object
func (node PrefixNode) GetChild(char rune) *PrefixNode {
	if node.HasChild(char) {
		return node.Children[char]
	}

	return nil
}
