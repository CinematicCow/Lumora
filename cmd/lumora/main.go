package main

import (
	"fmt"
	"log"
	"os"

	"github.com/CinematicCow/Lumora/internal/database"
	"github.com/CinematicCow/Lumora/internal/models"
)

func main() {
	fmt.Println("Hello, Lumora!")

	expandedPath := os.ExpandEnv(models.LUMORA_PATH)

	db, err := os.Open(expandedPath)
	if err != nil {
		log.Panicf("something went wrong at main in opening file: %v", err)
	}

	defer db.Close()

	dbForWrite, err := os.OpenFile(expandedPath, os.O_APPEND|os.O_WRONLY, 0600)

	if err != nil {
		log.Panicf("something went wrong at dbForWrite: %v", err)
	}

	defer dbForWrite.Close()

	d := &models.Data{
		Key:   []byte("cube"),
		Value: []byte("solved"),
	}

	if err := database.WriteToDB(dbForWrite, d); err != nil {
		log.Panicf("something went wrong at add: %v", err)
	}

	data, err := database.ReadFromDB(db)

	if err != nil {
		log.Panicf("something went wrong at main readDB: %v", err)
	}

	fmt.Print(data)

}
