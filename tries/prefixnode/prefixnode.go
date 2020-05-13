package prefixnode

//Node in the trie
type Node struct {
	Char     rune
	Children map[rune]*Node
	Users    map[string]struct{}
}

//New creates a n with character and terminal value
func New(char rune) Node {

	n := Node{
		Char:     char,
		Children: make(map[rune]*Node),
		Users:    make(map[string]struct{}),
	}

	return n
}

//GetChar associated with n
func (n Node) GetChar() rune {
	return n.Char
}

//HasChild if n has child with given rune value
func (n Node) HasChild(char rune) bool {

	_, found := n.Children[char]

	if found {
		return true
	}
	return false
}

//AddChild Node to parents children
func (n Node) AddChild(child *Node) bool {

	if n.HasChild(child.Char) {
		return false
	}
	//Add Node
	n.Children[child.Char] = child

	return true
}

//GetChild Returns a pointer to child object
func (n Node) GetChild(char rune) *Node {
	if n.HasChild(char) {
		return n.Children[char]
	}

	return nil
}

//GetChildren of a given n
func (n Node) GetChildren() map[rune]*Node {
	return n.Children
}

//AddUser to usernames
func (n Node) AddUser(name string) {
	n.Users[name] = struct{}{}
}

//GetUsers for a given n
func (n Node) GetUsers() map[string]struct{} {
	return n.Users
}
