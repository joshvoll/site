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

var mark = []byte(`---
title: The Beautiful in the Ugly
date: 2018-04-23
authors: 
   - Silver
---

# The Beautiful in the Ugly

Functional programming is nice and all, but sometimes you just need to have
things get done regardless of the consequences. Sometimes a dirty little hack
will suffice in place of a branching construct. This is a story of one of these
times.`)

type article struct {
	Title   string
	Authors []string
}

func TestMardownParse(t *testing.T) {
	var a article
	content, err := Unmarshal(mark, &a)
	if err != nil {
		t.Fatalf("cannot unmarshal de conentn: %v ", err)
	}
	fmt.Printf("%#v\n", a)
	fmt.Printf("%s\n", content)

}
