package database

import (
	"os"
	"sync"
	"testing"

	"github.com/CinematicCow/Lumora/internal/database"
	"github.com/CinematicCow/Lumora/internal/models"
)

var tempFilePool = sync.Pool{
	New: func() interface{} {
		tempFile, _ := os.CreateTemp("../../../tmp/", "test-lumora.gob")
		return tempFile
	},
}

var dataPool = sync.Pool{
	New: func() interface{} {
		return new(models.Data)
	},
}

func setupTestFile(t *testing.T) *os.File {
	return tempFilePool.Get().(*os.File)
}

func openTestFile(t *testing.T) *os.File {
	tempFile := setupTestFile(t)
	db := new(os.File)
	*db = *tempFile
	return db
}

type TestCase struct {
	name string
	test func(*testing.T)
}

var AddToDBCases = []TestCase{
	{
		name: "AddToDB",
		test: func(t *testing.T) {
			db := openTestFile(t)
			defer db.Close()

			data := dataPool.Get().(*models.Data)
			data.Key = []byte("key")
			data.Value = []byte("value")
			err := database.AddToDB(db, &data.Key, &data.Value)
			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}
			dataPool.Put(data)
		},
	},
}

var ListLumoraCases = []TestCase{
	{
		name: "GetAllFromDB",
		test: func(t *testing.T) {
			db := openTestFile(t)
			_, err := database.GetAllFromDB(db)
			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}
		},
	},
	{
		name: "empty database",
		test: func(t *testing.T) {
			db := openTestFile(t)
			lumora, err := database.GetAllFromDB(db)
			if err != nil {
				t.Fatalf("Expected no error, got %v", err)
			}
			if lumora != nil {
				t.Fatalf("Expected nil, got %v", lumora)
			}
		},
	},
}

func TestAddToDBCases(t *testing.T) {
	for _, test := range AddToDBCases {
		t.Run(test.name, test.test)
	}
}

func TestGetAllFromDBCases(t *testing.T) {
	for _, test := range ListLumoraCases {
		t.Run(test.name, test.test)
	}
}
