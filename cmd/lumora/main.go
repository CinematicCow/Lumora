package main

import (
	"fmt"
	"log"

	"github.com/CinematicCow/Lumora/internal/database"
)

func main() {
	fmt.Println("Hello, Lumora!")

	data, err := database.ListLumora()

	if err != nil {
		log.Panicf("something went wrong: %v", err)
	}

	fmt.Print(data)

}
