package prefixtree

import (
	"rr/tries/prefixnode"
)

//PrefixTree for nodes
type PrefixTree struct {
	name string
	root *prefixnode.Node
	size int
}

//New tree with a default root node
func New(name string) PrefixTree {
	rootNode := prefixnode.New('0')

	tree := PrefixTree{
		name: name,
		root: &rootNode,
		size: 0,
	}

	return tree
}

//InsertKeyword TODO: Replace USER STRING WITH USER OBJECT
func (tree PrefixTree) InsertKeyword(word string, userName string) {
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

//Contains returns list of users associated with word
func (tree PrefixTree) Contains(word string) []string {
	node := tree.root

	//Loop until end of word is hit
	for _, char := range word {
		if node.HasChild(char) {
			node = node.GetChild(char)
		} else {
			return nil
		}
	}

	return node.GetUsers()
}

//Getters and setters

//GetName returns the name of the tree
func (tree PrefixTree) GetName() string {
	return tree.name
}

//GetSize returns size of current tree
func (tree PrefixTree) GetSize() int {
	return tree.size
}

//GetRoot returns rootnode in a tree
func (tree PrefixTree) GetRoot() *prefixnode.Node {
	return tree.root
}
