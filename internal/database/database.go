package database

import (
	"encoding/gob"
	"io"
	"os"

	"github.com/CinematicCow/Lumora/internal/models"
)

func GetAllFromDB(file *os.File) ([]models.Lumora, error) {
	db := file

	var lumora []models.Lumora
	decoder := gob.NewDecoder(db)
	err := decoder.Decode(&lumora)
	if err == io.EOF {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	return lumora, nil
}

func AddToDB(file *os.File, lumora models.Lumora) error {
	db := file

	encoder := gob.NewEncoder(db)

	if err := encoder.Encode(lumora); err != nil {
		return err
	}

	return nil
}
