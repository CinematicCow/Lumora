package database__test

import (
	"os"
	"testing"

	"github.com/CinematicCow/Lumora/internal/database"
	"github.com/CinematicCow/Lumora/internal/models"
)

func BenchmarkWrite(b *testing.B) {
	// setup
	db, err := os.CreateTemp("../../../tmp/", "testdb")
	if err != nil {
		b.Fatal("Error while creating test db: ", err)
	}

	defer db.Close()

	d := &models.Data{
		Key:   []byte("test-key"),
		Value: []byte("test-value"),
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		err := database.WriteToDB(db, d)
		if err != nil {
			b.Fatal("Error while writing to db at bench: ", err)
		}
	}

	b.ReportAllocs()

}

func BenchmarkRead(b *testing.B) {
	// Setup
	db, err := os.CreateTemp("../../../tmp/", "testdb")
	if err != nil {
		b.Fatal("Error while creating test db: ", err)
	}
	defer db.Close()

	data := &models.Data{
		Key:   []byte("test-key"),
		Value: []byte("test-value"),
	}

	// Write some data to the db
	for i := 0; i < 1000; i++ {
		err := database.WriteToDB(db, data)
		if err != nil {
			b.Fatal("Error while writing to db: ", err)
		}
	}

	// Reset timer to exclude setup time
	b.ResetTimer()

	// Benchmark
	for i := 0; i < b.N; i++ {
		_, err := database.ReadFromDB(db)
		if err != nil {
			b.Fatal("Error while reading from db: ", err)
		}
	}

	// Report allocations
	b.ReportAllocs()
}
