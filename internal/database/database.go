package database

import (
	"io"
	"os"

	"github.com/CinematicCow/Lumora/internal/models"
	"github.com/CinematicCow/Lumora/internal/serde"
)

// Retrieves all data from the given database file.
//
// It takes a pointer to an os.File as its parameter.
// It returns a slice of values and an error.
func GetAllFromDB(db *os.File) ([]models.DecodedData, error) {

	var result []models.DecodedData

	buffer := make([]byte, 1024)

	for {
		n, err := db.Read(buffer)

		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		data := buffer[:n]

		key, value, err := serde.Deserialize(data)
		if err != nil {
			return nil, err
		}

		result = append(result, models.DecodedData{
			Key:   key,
			Value: value,
		})
	}

	return result, nil
}

// Adds a key-value pair to the database.
//
// It takes a pointer to a file, key, and value as parameters.
// It returns an error if there was a problem serializing the data or writing to the database.
func AddToDB(db *os.File, key, value *[]byte) error {
	data, err := serde.Serialize(key, value)

	if err != nil {
		return err
	}

	if _, err := db.Write(data); err != nil {
		return err
	}

	return nil
}
