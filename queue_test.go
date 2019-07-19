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

func TestAddWorkflow(t *testing.T) {
	q := new(Queue)
	if len(q.Workflows) != 0 {
		t.Errorf("Should have an empty q.Workflows -> %+v", q)
	}
	q.AddWorkflow("draft")
	if len(q.Workflows) != 1 {
		t.Errorf("Should have a single workflow -> %+v", q)
	}
	if strings.Compare(q.Workflows[0], "draft") != 0 {
		t.Errorf("Should have draft in q.Workflows[0] -> %+v", q)
	}
}
