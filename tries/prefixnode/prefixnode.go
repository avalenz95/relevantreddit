package prefixnode

//Node in the trie
type Node struct {
	char     rune
	terminal bool
	children map[rune]*Node
	users    []string
}

//New creates a n with character and terminal value
func New(char rune) Node {

	n := Node{
		char:     char,
		terminal: false,
		children: make(map[rune]*Node),
		users:    make([]string, 0, 0),
	}

	return n
}

//IsTerminal checks to see if current n is the end of a word
func (n Node) IsTerminal() bool {
	return n.terminal
}

//GetChar associated with n
func (n Node) GetChar() rune {
	return n.char
}

//HasChild if n has child with given rune value
func (n Node) HasChild(char rune) bool {

	_, found := n.children[char]

	if found {
		return true
	}
	return false
}

//AddChild Node to parents children
func (n Node) AddChild(child *Node) bool {

	if n.HasChild(child.char) {
		return false
	}
	//Add Node
	n.children[child.char] = child

	return true
}

//GetChild Returns a pointer to child object
func (n Node) GetChild(char rune) *Node {
	if n.HasChild(char) {
		return n.children[char]
	}

	return nil
}

//GetChildren of a given n
func (n Node) GetChildren() map[rune]*Node {
	return n.children
}

//AddUser at a terminal point
func (n Node) AddUser(name string) {
	n.users = append(n.users, name)
}

//GetUsers for a given n
func (n Node) GetUsers() []string {
	return n.users
}
