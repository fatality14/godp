package godp

import (
	"bytes"
	"encoding/gob"
	"os"
)

func SerializeData[T any](data T, file string) error {
	// Create a new buffer to write the serialized data to
	var b bytes.Buffer

	// Create a new gob encoder and use it to encode the T struct
	enc := gob.NewEncoder(&b)
	if err := enc.Encode(data); err != nil {
		return err
	}

	// The serialized data can now be found in the buffer
	serializedData := b.Bytes()
	err := os.WriteFile(file, serializedData, 0644)
	if err != nil {
		return err
	}

	return nil
}

func DeserializeData[T any](file string) (T, error) {
	// Create a struct to fill from the serialized data
	var res T

	data, err := os.ReadFile(file)
	if err != nil {
		return res, err
	}

	b := bytes.NewBuffer(data)

	// Create a new gob decoder and use it to decode the T struct
	dec := gob.NewDecoder(b)
	if err := dec.Decode(&res); err != nil {
		return res, err
	}

	return res, nil
}
