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
	"bytes"
	"io/ioutil"
	"os"
	"path"
	"strings"
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

// TestGenerateAndOrTOML() generates testout/andor.toml
// and then makes sure it can read it back.
func TestGenerateAndOrTOML(t *testing.T) {
	andorTOML := path.Join("testout", "andor.toml")
	collection := path.Join("testout", "repository.ds")
	if _, err := os.Stat("testout"); os.IsNotExist(err) {
		os.MkdirAll("testout", 0777)
	}
	if err := GenerateAndOrTOML(andorTOML, []string{collection}); err != nil {
		t.Errorf("Expected success, got %s", err)
		t.FailNow()
	}
	src, err := ioutil.ReadFile(andorTOML)
	if err != nil {
		t.Errorf("Can't read back %q, %s", andorTOML, err)
		t.FailNow()
	}
	if bytes.Contains(src, []byte(collection)) == false {
		t.Errorf("%q is missing %q", andorTOML, collection)
	}
	s, err := LoadAndOr(andorTOML)
	if err != nil {
		t.Errorf("problem loading %q, %s", andorTOML, err)
	}
	if len(s.Repositories) != 1 {
		t.Errorf("expected one repository got %d", len(s.Repositories))
	}
	if len(s.Repositories) == 1 {
		if strings.Compare(s.Repositories[0], collection) != 0 {
			t.Errorf("Repositories value not correct, %+v", s.Repositories)
		}
	}
}
