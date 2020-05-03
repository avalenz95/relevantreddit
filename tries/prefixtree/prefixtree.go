package prefixtree

import (
	"rr/tries/prefixnode"
)

//Prefix Tree
type PrefixTree struct {
	root *prefixnode.Node
	size int
}

func New() PrefixTree {
	rootNode := prefixnode.New('0')

	tree := PrefixTree{
		root: &rootNode,
		size: 0,
	}

	return tree
}
