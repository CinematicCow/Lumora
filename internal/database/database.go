package database

import (
	"encoding/gob"
	"io"
	"os"

	"github.com/CinematicCow/Lumora/internal/models"
)

func init() {
	gob.Register(models.Lumora{})
}

func ListLumora() ([]models.Lumora, error) {
	expandedPath := os.ExpandEnv(models.LUMORA_PATH)
	db, err := os.Open(expandedPath)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	var lumora []models.Lumora
	decoder := gob.NewDecoder(db)
	err = decoder.Decode(&lumora)
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
