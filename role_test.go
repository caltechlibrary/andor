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
	"fmt"
	"io/ioutil"
	"path"
	"strings"
	"testing"
)

// TestLoadRoles ...
func TestLoadRoles(t *testing.T) {
	roleFile := path.Join("testdata", "roles.toml")
	w, q, err := LoadRoles(roleFile)
	if err != nil {
		t.Errorf("Failed to load %q, %s", roleFile, err)
	}
	if len(w) != 3 {
		t.Errorf("expected 3 roles, got %d", len(w))
	}
	if len(q) != 3 {
		t.Errorf("expected 3 queues, got %d", len(q))
	}
	for _, wName := range []string{"draft", "review", "published"} {
		if role, ok := w[wName]; ok == false {
			t.Errorf("expected %q, not found -> %+v", wName, w)
			if strings.Compare(role.Queue, wName) != 0 {
				t.Errorf("expected %q, got %q for Queue", wName, role.Queue)
			}
		}
	}
}

// TestBytes for role structs
func TestBytes(t *testing.T) {
	roleFile := path.Join("testdata", "roles2.toml")
	roleSrc, err := ioutil.ReadFile(roleFile)
	if err != nil {
		t.Errorf("expected to read %s, %s", roleFile, err)
		t.FailNow()
	}
	w, _, err := LoadRoles(roleFile)
	if err != nil {
		t.Errorf("expected to load %q, got %s", roleFile, err)
	}
	src := []byte{}
	for _, k := range []string{"draft", "review", "published"} {
		v, _ := w[k]
		src = append(src, []byte(fmt.Sprintf("[%s]\n", k))...)
		src = append(src, v.Bytes()...)
		src = append(src, []byte("\n")...)
	}
	if len(src) == 0 {
		t.Errorf("expected a []byte with data for %+v", w)
	}
	if bytes.Compare(roleSrc, src) != 0 {
		t.Errorf("expected sources to match, got\n%s\n", src)
	}
}

// TestString for role structs
func TestString(t *testing.T) {
	roleFile := path.Join("testdata", "roles2.toml")
	roleSrc, err := ioutil.ReadFile(roleFile)
	if err != nil {
		t.Errorf("expected to read %s, %s", roleFile, err)
		t.FailNow()
	}
	w, _, err := LoadRoles(roleFile)
	if err != nil {
		t.Errorf("expected to load %q, got %s", roleFile, err)
	}
	s := []string{}
	for _, k := range []string{"draft", "review", "published"} {
		v, _ := w[k]
		s = append(s, fmt.Sprintf("[%s]\n", k))
		s = append(s, v.String())
		s = append(s, "\n")
	}
	src := strings.Join(s, "")
	if len(src) == 0 {
		t.Errorf("expected a string with data for %+v", w)
	}
	if strings.Compare(string(roleSrc), src) != 0 {
		t.Errorf("expected sources to match, got\n%s\n", src)
	}
}
