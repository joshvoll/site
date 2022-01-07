package front

import (
	"bytes"

	"gopkg.in/yaml.v2"
)

// Delimeter
var delim = []byte("---")

// Unmarshal is going to resolve everything
func Unmarshal(b []byte, v interface{}) (content []byte, err error) {
	if !bytes.HasPrefix(b, delim) {
		return b, nil
	}
	parts := bytes.SplitN(b, delim, 3)
	content = parts[2]
	err = yaml.Unmarshal(parts[1], v)
	return
}
