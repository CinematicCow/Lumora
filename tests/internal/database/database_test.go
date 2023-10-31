package database

import (
	"os"
	"testing"

	"github.com/CinematicCow/Lumora/internal/database"
	"github.com/CinematicCow/Lumora/internal/models"
)

func setupTestFile(t *testing.T) *os.File {
	// create a temp file to mock
	tempFile, err := os.CreateTemp("../../../temp/", "test-lumora.gob")
	if err != nil {
		t.Fatalf("Failed to create temporary file: %v", err)
	}
	return tempFile
}

func openTestFile(t *testing.T) *os.File {
	tempFile := setupTestFile(t)

	// open the mock file
	db, err := os.Open(tempFile.Name())
	if err != nil {
		t.Fatalf("Failed to open temporary file: %v", err)
	}
	return db
}

type TestCase struct {
	name string
	test func(t *testing.T)
}

var AddLumoraCases = []TestCase{
	{
		name: "AddLumora",
		test: func(t *testing.T) {

			db := setupTestFile(t)
			defer db.Close()

			data := models.Lumora{
				Key:   "test",
				Value: "test",
			}
			err := database.AddLumora(db, data)
			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}
		},
	},
}

var ListLumoraCases = []TestCase{
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

func TestAddLumoraCases(t *testing.T) {
	for _, test := range AddLumoraCases {
		t.Run(test.name, test.test)
	}
}

func TestListLumoraCases(t *testing.T) {
	for _, test := range ListLumoraCases {
		t.Run(test.name, test.test)
	}
}
