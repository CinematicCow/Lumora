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

	dbForWrite, err := os.Create(expandedPath)

	if err != nil {
		log.Panicf("something went wrong at dbForWrite: %v", err)
	}

	defer dbForWrite.Close()

	Key := []byte("hello1")
	Value := []byte("world2")

	if err := database.AddToDB(dbForWrite, &Key, &Value); err != nil {
		log.Panicf("something went wrong at add: %v", err)
	}

	data, err := database.GetAllFromDB(db)

	if err != nil {
		log.Panicf("something went wrong at main getAll: %v", err)
	}

	fmt.Print(data)

}
