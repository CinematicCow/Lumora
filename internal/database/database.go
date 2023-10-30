package database

import (
	"encoding/gob"
	"fmt"
	"os"

	"github.com/CinematicCow/Lumora/internal/models"
)

func init() {
	gob.Register(models.Lumora{})
}

func ListLumora() []models.Lumora {
	expandedPath := os.ExpandEnv(models.LUMORA_PATH)
	fmt.Println(expandedPath)
	db, err := os.Open(expandedPath)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	decoder := gob.NewDecoder(db)

	var lumora []models.Lumora

	err = decoder.Decode(&lumora)
	if err != nil {
		panic(err)
	}

	return lumora
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
