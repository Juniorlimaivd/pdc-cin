package global

import (
	"bytes"
	"encoding/gob"
	"log"
)

// Marshaller encodes and decodes objects to and from byte arrays
type Marshaller struct{}

// Marshall ..
func (m *Marshaller) Marshall(data interface{}) []byte {
	buffer := new(bytes.Buffer)

	err := gob.NewEncoder(buffer).Encode(data)

	if err != nil {
		log.Fatal("Encode error:", err)
	}

	return buffer.Bytes()
}

// Unmarshall ...
func (m *Marshaller) Unmarshall(data []byte, result interface{}) error {

	buffer := bytes.NewBuffer(data)
	err := gob.NewDecoder(buffer).Decode(result)
	return err
}
