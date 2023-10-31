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
	defer db.Close()

	data, err := database.ListLumora(db)

	if err != nil {
		log.Panicf("something went wrong: %v", err)
	}

	fmt.Print(data)

}
