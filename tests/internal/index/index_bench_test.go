package index__t_test

import (
	"bytes"
	"fmt"
	"log"
	"testing"

	"github.com/CinematicCow/Lumora/internal/index"
)

func BenchmarkBTree(b *testing.B) {
	// Create a new BTree
	bt, err := index.NewBTree("./testdb")
	if err != nil {
		log.Fatal(err)
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// Add a key to the BTree
		key := []byte("testKey" + fmt.Sprint(i))
		err = index.AddKey(bt, key)
		if err != nil {
			log.Fatal(err)
		}

		// Find the key in the BTree
		value, err := index.FindKey(bt, key)
		if err != nil {
			log.Fatal(err)
		}

		// Check if the value is correct
		if !bytes.Equal(value, key) {
			b.Errorf("Expected %v, got %v", key, value)
		}

		// Remove the key from the BTree
		err = index.RemoveKey(bt, key)
		if err != nil {
			log.Fatal(err)
		}
	}
}
