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

//TODO: Replace USER STRING WIHT USER OBJECT
func (tree PrefixTree) insertKeyword(word string, userName string) {
	node := tree.root

	for index, char := range word {
		//Node does not have child with character
		if !node.HasChild(char) {
			child := prefixnode.New(char)
			node.AddChild(&child)
		}

		//Child already exists advance pointer
		node = node.GetChild(char)

		//End of word add username to word
		if index == len(word)-1 {
			node.SetTerminal()
			node.AddUser(userName)
		}
	}

}
