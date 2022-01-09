package front

import (
	"bytes"

	"gopkg.in/yaml.v2"
)

// Delimeter ..
var delimeter = []byte("---")

// Unmarshal parses YAML frontmatter and returns the content. When no
// frontmatter delimiters are present the original content is returned.
func Unmarshal(b []byte, v interface{}) (content []byte, err error) {
	if !bytes.HasPrefix(b, delimeter) {
		return b, nil
	}
	parts := bytes.SplitN(b, delimeter, 3)
	content = parts[2]
	err = yaml.Unmarshal(parts[1], v)
	return
}
