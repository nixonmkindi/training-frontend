package helpers

import (
	"bytes"
	"encoding/json"
	"training-frontend/package/log"
)

// converts maps into struct
func Decode(in, out interface{}) {
	buf := new(bytes.Buffer)
	err := json.NewEncoder(buf).Encode(in)
	if err != nil {
		log.Errorf("error creating new encoder: %v", err)
	}
	err = json.NewDecoder(buf).Decode(out)
	if err != nil {
		log.Errorf("error creating new decoder: %v", err)
	}
}
