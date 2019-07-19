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
	"path"
	"testing"
)

// TestLoadAndOr() tests if we can create a service object from
// loading test copies of testdata/workflows.toml, testdata/users.toml
/// and testdata/andor.toml.
func TestLoadAndOr(t *testing.T) {
	andorTOML := path.Join("testdata", "andor.toml")
	if _, err := LoadAndOr(andorTOML); err != nil {
		t.Errorf("LoadAndOr(%q) %s", andorTOML, err)
		t.FailNow()
	}
}
