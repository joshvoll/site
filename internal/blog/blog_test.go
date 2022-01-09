package blog

import (
	"testing"
)

func TestLoadBlog(t *testing.T) {
	_, err := LoadPosts("../../blog")
	if err != nil {
		t.Fatalf("erroe loading blog: %v ", err)
	}

}
