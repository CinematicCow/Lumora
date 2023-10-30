package database

import (
	"os"
	"testing"

	"github.com/CinematicCow/Lumora/internal/database"
)

func setupTestFile(t *testing.T) *os.File {
	// create a temp file to mock
	tempFile, err := os.CreateTemp("", "test-lumora.gob")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	return tempFile
}

func openTestFile(t *testing.T) *os.File {
	tempFile := setupTestFile(t)
	defer tempFile.Close()

	// open the mock file
	db, err := os.Open(tempFile.Name())
	if err != nil {
		t.Fatalf("Failed to open temporary file: %v", err)
	}
	return db
}

type ListLumoraCase struct {
	name string
	test func(t *testing.T)
}

var ListLumoraCases = []ListLumoraCase{
	{
		name: "ListLumora",
		test: func(t *testing.T) {
			db := openTestFile(t)
			_, err := database.ListLumora(db)
			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}
		},
	},
	{
		name: "empty database",
		test: func(t *testing.T) {

			db := openTestFile(t)
			// call the ListLumora function
			lumora, err := database.ListLumora(db)

			// assert no error
			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}

			// assert lumora returned EOF
			if lumora != nil {
				t.Fatalf("Expected nil, got %v", lumora)
			}
		},
	},
}

func TestListLumoraCases(t *testing.T) {
	for _, test := range ListLumoraCases {
		t.Run(test.name, test.test)
	}
}
