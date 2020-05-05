package prefixtree

import (
	"rr/tries/prefixnode"
)

//PrefixTree for nodes
type PrefixTree struct {
	subredditName string
	root          *prefixnode.Node
	size          int
}

//New tree with a default root node
func New() PrefixTree {
	rootNode := prefixnode.New('0')

	tree := PrefixTree{
		root: &rootNode,
		size: 0,
	}

	return tree
}

//TODO: Replace USER STRING WITH USER OBJECT
func (tree PrefixTree) insertKeyword(word string, userName string) {
	node := tree.root

	for index, char := range word {
		//Node does not have child with character
		if !node.HasChild(char) {
			child := prefixnode.New(char)
			node.AddChild(&child)
			tree.size++
		}

		//Child already exists advance pointer
		node = node.GetChild(char)

		//End of word add username to word
		if index == len(word)-1 {
			node.AddUser(userName)
		}

	}

}

//checks if a word is contained and returns list of users associated with word
func (tree PrefixTree) contains(word string) []string {
	node := tree.root

	//Loop until end of word is hit
	for _, char := range word {
		if node.HasChild(char) {
			node = node.GetChild(char)
		}
	}

	return node.GetUsers()
}
