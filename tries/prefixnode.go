package prefixnode

//prefix node in the trie
type prefixNode struct {
	char     string
	terminal bool
	children map[string]*prefixNode
}
