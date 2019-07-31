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
	roles, queues, err := LoadRoles(roleFile)
	if err != nil {
		t.Errorf("Failed to load %q, %s", roleFile, err)
		t.FailNow()
	}
	if roles == nil || len(roles) != 3 {
		t.Errorf("expected 3 roles, got %d", len(roles))
		t.FailNow()
	}
	if queues == nil || len(queues) != 3 {
		t.Errorf("expected 3 queues, got %d", len(queues))
		t.FailNow()
	}
	// Check for roles, then check queues
	//queues := []string{"draft", "review", "published"}
	for _, roleName := range []string{"writer", "editor", "public"} {
		if role, ok := roles[roleName]; ok == false || role == nil {
			t.Errorf("expected %q, not found in %q", roleName, roleFile)
		} else {
			if role.Queues == nil || len(role.Queues) == 0 {
				t.Errorf("expected at least one state in queue, %s %s", roleName, roleFile)
				t.FailNow()
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
	for _, k := range []string{"writer", "reviewer", "public"} {
		v, ok := w[k]
		if ok {
			src = append(src, []byte(fmt.Sprintf("[%s]\n", k))...)
			src = append(src, v.Bytes()...)
			src = append(src, []byte("\n")...)
		} else {
			t.Errorf("Missing %s in test data, test is in error for %s", k, roleFile)
			t.FailNow()
		}
	}
	if len(src) == 0 {
		t.Errorf("expected a []byte with data for file %s", roleFile)
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
	for _, k := range []string{"writer", "reviewer", "public"} {
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
		t.Errorf("expected sources to match, got\n%s\n------\n%s", src, roleSrc)
	}
}
