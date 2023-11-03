package database__test

import (
	"os"
	"testing"
)

func createDBFile(t *testing.T) *os.File {
	f, err := os.CreateTemp("", "testdb")
	if err != nil {
		t.Fatal("Error while creating test db: ", err)
	}
	return f
}

func readDBFile(t *testing.T) *os.File {

	tf := createDBFile(t)
	defer tf.Close()

	d, err := os.Open(tf.Name())
	if err != nil {
		t.Fatal("Error while opening test db: ", err)
	}
	return d
}

func TestReadFromDB(t *testing.T) {

	// read from empty db

}

func TestWriteToDB(t *testing.T) {

}
