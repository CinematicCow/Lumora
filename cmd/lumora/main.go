package main

import (
	"fmt"

	"github.com/CinematicCow/Lumora/internal/database"
)

func main() {
	fmt.Println("Hello, Lumora!")

	database.ListLumora()
}
