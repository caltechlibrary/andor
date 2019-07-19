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
	"path"
	"strings"
	"testing"
)

func TestLoadUsers(t *testing.T) {
	usersTOML := path.Join("testdata", "users.toml")
	usersTOMLSrc := []byte(`
#
# Example Test users file for testing 
# LoadUsers()
#

# User id
["jane.doe@example.edu"]
# Display name
display_name = "Jane Doe"
# By default objects are create in this queue
create_queue = "deposit"
# Jane is a member of the "deposit" workflow/queue
member_of = ["deposit"]
`)
	err := ioutil.WriteFile(usersTOML, usersTOMLSrc, 0666)
	if err != nil {
		t.Errorf("%s, %s", usersTOML, err)
		t.FailNow()
	}
	if _, err := LoadUsers(usersTOML); err != nil {
		t.Errorf("%s", err)
		t.FailNow()
	}
}

func TestUserBytes(t *testing.T) {
	expected := []byte(`user_id = "jane.doe@example.edu"
display_name = "Jane Doe"
member_of = ["deposit"]
`)
	u := new(User)
	u.Key = "jane.doe@example.edu"
	u.DisplayName = "Jane Doe"
	u.MemberOf = []string{"deposit"}
	got := u.Bytes()
	if bytes.Compare(expected, got) != 0 {
		t.Errorf("expected\n%s\n\ngot\n\n%s\n", expected, got)
	}
}

func TestUserString(t *testing.T) {
	expected := `user_id = "jane.doe@example.edu"
display_name = "Jane Doe"
member_of = ["deposit"]
`
	u := new(User)
	u.Key = "jane.doe@example.edu"
	u.DisplayName = "Jane Doe"
	u.MemberOf = []string{"deposit"}
	got := u.String()
	if strings.Compare(expected, got) != 0 {
		t.Errorf("expected\n%s\n\ngot\n\n%s\n", expected, got)
	}
}
