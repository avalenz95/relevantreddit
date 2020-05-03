package prefixnode

//PrefixNode in the trie
type PrefixNode struct {
	char     rune
	terminal bool
	children map[rune]*PrefixNode
	users    []string
}

//New creates a node with character and terminal value
func New(char rune) PrefixNode {

	node := PrefixNode{
		char,
		false,
		make(map[rune]*PrefixNode),
		make([]string, 0, 0),
	}

	return node
}

//IsTerminal checks to see if current node is the end of a word
func (node PrefixNode) IsTerminal() bool {
	return node.terminal
}

//GetChar associated with node
func (node PrefixNode) GetChar() rune {
	return node.char
}

//HasChild if node has child with given rune value
func (node PrefixNode) HasChild(char rune) bool {

	_, found := node.children[char]

	if found {
		return true
	}
	return false
}

//AddChild Node to parents children
func (node PrefixNode) AddChild(child *PrefixNode) bool {

	if node.HasChild(child.char) {
		return false
	}
	//Add Node
	node.children[child.char] = child

	return true
}

//GetChild Returns a pointer to child object
func (node PrefixNode) GetChild(char rune) *PrefixNode {
	if node.HasChild(char) {
		return node.children[char]
	}

	return nil
}

//GetChildren of a given node
func (node PrefixNode) GetChildren() map[rune]*PrefixNode {
	return node.children
}

//AddUser at a terminal point
func (node PrefixNode) AddUser(name string) {
	node.users = append(node.users, name)
}

//GetUsers for a given node
func (node PrefixNode) GetUsers() []string {
	return node.users
}
