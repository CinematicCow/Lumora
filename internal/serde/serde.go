package serde

import (
	"github.com/CinematicCow/Lumora/internal/models"
	"github.com/shamaton/msgpack/v2"
)

func Serialize(data *models.Data) ([]byte, error) {
	d, err := msgpack.Marshal(data)

	if err != nil {
		return nil, err
	}

	return d, nil
}

func Deserialize(data []byte) (models.Data, error) {

	var result models.Data
	err := msgpack.Unmarshal(data, &result)

	if err != nil {
		return result, err
	}

	return result, nil

}
