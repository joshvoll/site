package front

import (
	"fmt"
	"testing"
)

var markdown = []byte(`---
title: Ferrets
authors:
  - Tobi
  - Loki
  - Jane
---
Some content here, so
interesting, you just
want to keep reading.`)

type article struct {
	Title   string
	Authors []string
}

// TestFrontParse is going to parse the
func TestFrontParse(t *testing.T) {
	var a article
	content, err := Unmarshal(markdown, &a)
	if err != nil {
		t.Errorf("generating error: %v ", err)
	}
	fmt.Printf("%#v\n", a)
	fmt.Printf("%s\n", string(content))
}
