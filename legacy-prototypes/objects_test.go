//
// Package andor provides support for building simple digital
// object repositories in Go where objects are stored in a
// dataset collection and the UI of the repository is static
// HTML 5 documents using JavaScript to access a web API.
//
// @Author R. S. Doiel, <rsdoiel@library.caltech.edu>
//
package andor

import (
	"strings"
	"testing"
)

var (
	obj = map[string]interface{}{
		"_State": "published",
		"_Key":   "1",
	}
)

func TestObjects(t *testing.T) {
	state := ObjectState(obj)
	key := ObjectKey(obj)
	if strings.Compare(state, "published") != 0 {
		t.Errorf("expected published, got %q", state)
	}
	if strings.Compare(key, "1") != 0 {
		t.Errorf("expected 1, got %q", key)
	}
}
