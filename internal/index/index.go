package index

import (
	"bytes"
	"encoding/binary"
	"fmt"

	"github.com/timtadh/fs2/bptree"
	"github.com/timtadh/fs2/fmap"
)

// BTree is a struct that wraps a BTree from the google/btree package.
// It is used to store keys in a B+ tree.
type BTree struct {
	tree *bptree.BpTree
}

// NewBTree creates a new BTree with a BTree of degree 32.
func NewBTree(filepath string) (*bptree.BpTree, error) {
	bf, err := fmap.CreateBlockFile(filepath)
	if err != nil {
		return nil, err
	}
	bpt, err := bptree.New(bf, 8, -1)
	if err != nil {
		bf.Close()
		return nil, err
	}
	return bpt, nil
}

func OpenBTree(filepath string) (*bptree.BpTree, error) {
	bf, err := fmap.OpenBlockFile(filepath)
	if err != nil {
		return nil, err
	}
	bpt, err := bptree.Open(bf)
	if err != nil {
		bf.Close()
		return nil, err
	}
	return bpt, nil

}

// AddKey adds a key to the B+ tree.
func AddKey(b *bptree.BpTree, key []byte) error {
	kBytes := make([]byte, 8)
	binary.PutUvarint(kBytes, uint64(len(key)))
	err := b.Add(kBytes, key)
	if err != nil {
		return err
	}
	return nil
}

// FindKey checks if a key exists in the B+ tree.
// returns true if key exists.
func FindKey(b *bptree.BpTree, key []byte) ([]byte, error) {
	kBytes := make([]byte, 8)
	binary.PutUvarint(kBytes, uint64(len(key)))
	kvi, err := b.Find(kBytes)
	if err != nil {
		return nil, err
	}
	var value []byte
	for key, value, err, kvi = kvi(); kvi != nil; key, value, err, kvi = kvi() {
		if bytes.Equal(key, kBytes) {
			return value, nil
		}
	}

	return nil, fmt.Errorf("key not found")
}

// RemoveKey deletes a key from the B+ tree.
func RemoveKey(b *bptree.BpTree, key []byte) error {
	kBytes := make([]byte, 8)
	binary.PutUvarint(kBytes, uint64(len(key)))
	err := b.Remove(kBytes, func(value []byte) bool {
		return true
	})
	if err != nil {
		return err
	}
	return nil
}
