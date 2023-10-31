package yaml

import (
	"github.com/ghodss/yaml"
)

type yamlEnCoder struct{}

func (encoder yamlEnCoder) Encode(i interface{}) ([]byte, error) {
	return yaml.Marshal(i)
}

func (encoder yamlEnCoder) Decode(bytes []byte, i interface{}) error {
	return yaml.Unmarshal(bytes, i)
}

func (encoder yamlEnCoder) String() string {

	return "yaml"
}
