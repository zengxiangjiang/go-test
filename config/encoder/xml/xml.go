package xml

import (
	"encoding/xml"
)

type xmlEncoder struct{}

func (encoder xmlEncoder) Encode(i interface{}) ([]byte, error) {
	return xml.Marshal(i)
}

func (encoder xmlEncoder) Decode(bytes []byte, i interface{}) error {
	return xml.Unmarshal(bytes, i)
}

func (encoder xmlEncoder) String() string {
	return "xml"
}
