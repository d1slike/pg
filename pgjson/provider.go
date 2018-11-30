package pgjson

import (
	"encoding/json"
	"io"
)

type jsonProvider interface {
	Marshal(v interface{}) ([]byte, error)
	NewDecoder(r io.Reader) interface {
		Decode(v interface{}) error
	}
}

type defualtJsonProvider struct {
}

func (p *defualtJsonProvider) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (p *defualtJsonProvider) NewDecoder(r io.Reader) interface {
	Decode(v interface{}) error
} {
	return json.NewDecoder(r)
}

var (
	JsonProvider jsonProvider = &defualtJsonProvider{}
)
