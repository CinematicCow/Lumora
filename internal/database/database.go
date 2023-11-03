package database

import (
	"bufio"
	"io"
	"log"
	"os"

	"github.com/CinematicCow/Lumora/internal/models"
	"github.com/CinematicCow/Lumora/internal/serde"
)

func WriteToDB(db *os.File, data *models.Data) error {

	sd, err := serde.Serialize(data)

	if err != nil {
		log.Fatal("Error while serializing", err)
		return err
	}

	n, err := db.Write(append(sd, '\n'))

	if err != nil {
		log.Fatal("Error while writing to db", err)
		return err
	}

	log.Default().Printf("data wrote: %d\ndata length: %d", n, len(sd))

	if n-1 != len(sd) {
		log.Fatal("Mismatch in number of bytes written", err)
		return err
	}

	return nil

}

func ReadFromDB(db *os.File) ([]models.DecodedData, error) {

	scanner := bufio.NewScanner(db)

	var result []models.DecodedData

	for scanner.Scan() {
		line := scanner.Bytes()

		d, err := serde.Deserialize(line)

		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}

		Key := string(d.Key)
		Value := string(d.Value)
		result = append(result, models.DecodedData{Key, Value})
	}

	if err := scanner.Err(); err != nil {
		log.Fatal("Error while reading from db", err)
		return nil, err
	}

	return result, nil
}
