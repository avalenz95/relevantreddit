package prefixnode_test

import (
	"rr/tries/prefixnode"
	"testing"
)

//test creation of a node object
func TestNew(t *testing.T) {

	var tests = []struct {
		char rune // input
	}{
		{'a'},
		{'b'},
		{'c'},
		{'d'},
		{'e'},
		{'A'},
		{'B'},
		{'@'},
	}

	for _, test := range tests {

		node := prefixnode.New(test.char)

		//Check Character
		if node.Char != test.char {
			t.Errorf("Expected Character: %c, got %c instead.", test.char, node.Char)
		}

		//Check Terminal
		if node.Terminal != false {
			t.Errorf("Expected Terminal Value: %v, got %v  instead.", false, node.Terminal)
		}

		//Check Children Node length
		if len(node.Children) != 0 {
			t.Errorf("Expected children nodes of length %v, instead got length %v.", 0, len(node.Children))
		}

		//Check count of users
		if len(node.Users) != 0 {
			t.Errorf("Expected users number of users %v, found %v users instead.", 0, len(node.Users))
		}

	}

}

func TestIsTerminal(t *testing.T) {

	node := prefixnode.New('a')

	//Check Terminal
	if node.Terminal != false {
		t.Errorf("Expected Terminal Value: %v, got %v  instead.", false, node.Terminal)
	}

	node.Terminal = true

	//Check Terminal
	if node.Terminal != true {
		t.Errorf("Expected Terminal Value: %v, got %v  instead.", true, node.Terminal)
	}

}

func TestAddChild(t *testing.T) {
	node := prefixnode.New('a')
	//Insert child
	child := prefixnode.New('b')
	result := node.AddChild(&child)
	if result != true {
		t.Errorf("Expected insertion of %v and return value of true.", child)
	}

	//Test duplicate insertion
	child1 := prefixnode.New('b')
	result1 := node.AddChild(&child1)
	if result1 != false {
		t.Errorf("Expected insertion to be invalid for %v", child1)
	}
}
