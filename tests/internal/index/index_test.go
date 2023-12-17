package index__test

import (
	"bytes"
	"github.com/CinematicCow/Lumora/internal/index"
	"log"
	"testing"
)

func TestBTree(t *testing.T) {
	// Create a new BTree
	b, err := index.NewBTree("./testdb")
	if err != nil {
		log.Fatal(err)
	}

	// Add a key to the BTree
	key := []byte("testKey")
	err = index.AddKey(b, key)
	if err != nil {
		log.Fatal(err)
	}

	// Find the key in the BTree
	value, err := index.FindKey(b, key)
	if err != nil {
		log.Fatal(err)
	}

	// Check if the value is correct
	if !bytes.Equal(value, key) {
		t.Errorf("Expected %v, got %v", key, value)
	}

	// Remove the key from the BTree
	err = index.RemoveKey(b, key)
	if err != nil {
		log.Fatal(err)
	}

	// Try to find the key again
	value, err = index.FindKey(b, key)
	if err == nil {
		t.Errorf("Expected an error, got nil")
	}
}
