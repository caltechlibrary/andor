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
		"_Queue": "published",
		"_Key":   "1",
	}
)

func TestObjects(t *testing.T) {
	queue := ObjectQueue(obj)
	key := ObjectKey(obj)
	if strings.Compare(queue, "published") != 0 {
		t.Errorf("expected published, got %q", queue)
	}
	if strings.Compare(key, "1") != 0 {
		t.Errorf("expected 1, got %q", key)
	}
}
