package index

import (
	"bytes"

	"github.com/google/btree"
)

// BTree is a struct that wraps a BTree from the google/btree package.
// It is used to store keys in a B+ tree.
type BTree struct {
	tree *btree.BTree
}

// KeyItem is a struct that wraps a []byte key.
// It is used to store keys in the B+ tree.
type KeyItem struct {
	Key []byte
}

// NewBTree creates a new BTree with a BTree of degree 32.
func NewBTree() (*BTree, error) {
	t := btree.New(32)
	return &BTree{tree: t}, nil
}

// Less is a method that implements the Item interface from the google/btree package.
// It compares two KeyItem instances based on their keys.
func (k KeyItem) Less(than btree.Item) bool {
	return bytes.Compare(k.Key, than.(KeyItem).Key) < 0
}

// Insert adds a key to the B+ tree.
func (b *BTree) Insert(key []byte) {
	b.tree.ReplaceOrInsert(KeyItem{Key: key})
}

// Has checks if a key exists in the B+ tree.
// returns true if key exists.
func (b *BTree) Has(key []byte) bool {
	return b.tree.Has(KeyItem{Key: key})
}

// Remove deletes a key from the B+ tree.
func (b *BTree) Remove(key []byte) {
	b.tree.Delete(KeyItem{Key: key})
}
