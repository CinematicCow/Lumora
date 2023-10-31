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
		log.Panicf("something went wrong: %v", err)
	}

	if err := database.AddToDB(db, models.Lumora{
		Key:   "test",
		Value: "test",
	}); err != nil {
		log.Panicf("something went wrong at add: %v", err)
	}

	data, err := database.GetAllFromDB(db)

	if err != nil {
		log.Panicf("something went wrong: %v", err)
	}

	fmt.Print(data)
	defer db.Close()

}
