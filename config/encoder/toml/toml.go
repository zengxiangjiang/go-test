package toml

import (
	"bytes"
)
import "github.com/BurntSushi/toml"

type tomlEncoder struct{}

func (t tomlEncoder) Encode(i interface{}) ([]byte, error) {
	buffer := bytes.NewBuffer(nil)
	defer buffer.Reset()
	err := toml.NewEncoder(buffer).Encode(i)
	if err != nil {
		return nil, err
	}
	return buffer.Bytes(), nil
}

func (t tomlEncoder) Decode(bytes []byte, i interface{}) error {

	return toml.Unmarshal(bytes, i)
}

func (t tomlEncoder) String() string {
	return "toml"
}
