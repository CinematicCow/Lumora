package database

import (
	"encoding/gob"
	"io"
	"os"

	"github.com/CinematicCow/Lumora/internal/models"
)

func ListLumora(file *os.File) ([]models.Lumora, error) {
	db := file
	defer db.Close()

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

func AddLumora(lumora models.Lumora) {
	db, err := os.Open(models.LUMORA_PATH)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	encoder := gob.NewEncoder(db)

	err = encoder.Encode(lumora)
	if err != nil {
		panic(err)
	}

}
