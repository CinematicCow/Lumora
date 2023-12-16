package index__test

import (
	"testing"

	"github.com/CinematicCow/Lumora/internal/index"
)

func TestBTree(t *testing.T) {

	// create a new btree
	tree, err := index.NewBTree()
	if err != nil {
		t.Error(err)
	}

	// insert some keys
	keys := [][]byte{
		[]byte("key1"),
		[]byte("key2"),
		[]byte("key3"),
		[]byte("key4"),
		[]byte("key5"),
	}

	for _, key := range keys {
		tree.Insert(key)
	}

	// check if the keys exist in the tree

	for _, key := range keys {
		if !tree.Has(key) {
			t.Errorf("Key %s does not exist in the tree", key)
		}
	}

	// remove a key
	tree.Remove(keys[2])

	// check if the key has been removed
	if tree.Has(keys[2]) {
		t.Errorf("Key %s has not been removed from the tree", keys[2])
	}
}
