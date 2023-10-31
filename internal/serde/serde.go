package serde

import (
	"bytes"
	"encoding/gob"

	"github.com/CinematicCow/Lumora/internal/models"
)

// Serialize the given key and value into a byte slice.
//
// It takes in two byte pointers, key and value, and returns a byte slice and an
// error. The byte slice represents the serialized data, and the error indicates
// any issues encountered during serialization.
func Serialize(key, value *[]byte) ([]byte, error) {
	var buf bytes.Buffer
	encoder := gob.NewEncoder(&buf)

	data := models.Data{
		Key:   *key,
		Value: *value,
	}

	if err := encoder.Encode(data); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// Deserialize the given byte slice into a key and value string.
//
// It takes a byte slice as a parameter and returns a key string, a value string, and an error.
func Deserialize(data []byte) (key string, value string, err error) {
	decoder := gob.NewDecoder(bytes.NewReader(data))

	var model models.Data

	if err = decoder.Decode(&model); err != nil {
		return
	}

	key = string(model.Key)
	value = string(model.Value)

	return key, value, nil
}
