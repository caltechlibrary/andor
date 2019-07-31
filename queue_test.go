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

func TestAddRole(t *testing.T) {
	q := new(Queue)
	if len(q.Roles) != 0 {
		t.Errorf("Should have an empty q.Roles -> %+v", q)
	}
	q.AddRole("draft")
	if len(q.Roles) != 1 {
		t.Errorf("Should have a single role -> %+v", q)
	}
	if strings.Compare(q.Roles[0], "draft") != 0 {
		t.Errorf("Should have draft in q.Roles[0] -> %+v", q)
	}
}
