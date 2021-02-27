package api

import (
	"bytes"
	"encoding/gob"
	"log"

	"github.com/pkg/errors"
)

func encodeToBytes(p interface{}) []byte {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(p); err != nil {
		log.Fatal(errors.WithStack(err))
		panic(errors.WithStack(err))
	}
	return buf.Bytes()
}