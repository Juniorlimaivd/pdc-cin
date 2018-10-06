package middleware

import (
	"bytes"
	"encoding/gob"
	"log"
)

// Marshaller encodes and decodes objects to and from byte arrays
type Marshaller struct{}

func (m *Marshaller) marshall(data interface{}) []byte {
	buffer := new(bytes.Buffer)

	err := gob.NewEncoder(buffer).Encode(data)

	if err != nil {
		log.Fatal("Encode error:", err)
	}

	return buffer.Bytes()
}

func (m *Marshaller) unmarshall(data []byte, result interface{}) {

	buffer := bytes.NewBuffer(data)
	gob.NewDecoder(buffer).Decode(result)
}
