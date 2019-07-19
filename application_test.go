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
	"path"
	"strings"
	"testing"
)

func TestApplication(t *testing.T) {
	input := []byte{}
	output := []byte{}
	errorOut := []byte{}
	Stdin := bytes.NewReader(input)
	Stdout := bytes.NewBuffer(output)
	Stderr := bytes.NewBuffer(errorOut)

	appName := "ApplicationTest"
	andorTOML := path.Join("testdata", "andor.toml")
	args := []string{}

	// Testing init without parameters, then with
	args = append(args, "init")
	if r := Application(appName, andorTOML, args, Stdin, Stdout, Stderr); r != 0 {
		t.Errorf("expected 0, got %d for %s %s", r, appName, strings.Join(args, " "))
	}

	// Now envoke with repository names
	args = []string{}
	args = append(args, "test_repo1.ds", "test_repo2.ds", "test_repo3.ds")
	if r := Application(appName, andorTOML, args, Stdin, Stdout, Stderr); r != 1 {
		t.Errorf("expected 1, got %d for %s %s", r, appName, strings.Join(args, " "))
	}
}
