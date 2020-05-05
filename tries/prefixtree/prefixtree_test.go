package prefixtree_test

import (
	"rr/tries/prefixtree"
	"testing"
)

func TestNew(t *testing.T) {
	tree := prefixtree.New("Test")

	//Check Name
	if tree.GetName() != "Test" {
		t.Errorf("Expected Name: Test: %v instead.", tree.GetName())
	}

	//Check size
	if tree.GetSize() != 0 {
		t.Errorf("Expected Size: 0, instead had size %v", tree.GetSize())
	}

	//Get Root rune
	if tree.GetRoot().GetChar() != '0' {
		t.Errorf("Expected Rune 0 instead had rune %v", tree.GetRoot().GetChar())
	}

}

func TestInsertKeyword(t *testing.T) {
	tree := prefixtree.New("Test")

	tree.InsertKeyword("apple", "becky")
}
