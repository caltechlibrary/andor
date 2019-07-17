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
	"io/ioutil"
	"testing"

	// Caltech Library packages
	"github.com/caltechlibrary/dataset"
)

func TestLoadUser(t *testing.T) {
	AndOrUsers = "test_users.AndOr"
	testUsers := "test_users.toml"
	testUsersSrc := []byte(`
#
# Example Test users file for testing 
# LoadUser()
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
	err := ioutil.WriteFile(testUsers, testUsersSrc, 0666)
	if err != nil {
		t.Errorf("%s, %s", testUsers, err)
		t.FailNow()
	}
	// Initial AndOrUsers just in case
	_, err = dataset.InitCollection(AndOrUsers)
	if err != nil {
		t.Errorf("%s, %s", AndOrUsers, err)
		t.FailNow()
	}

	// Reset the test collection for AndOrUsers
	c, err := dataset.Open(AndOrUsers)
	if err != nil {
		t.Errorf("%s, %s", AndOrUsers, err)
		t.FailNow()
	}
	keys := c.Keys()
	for _, key := range keys {
		if err := c.Delete(key); err != nil {
			t.Errorf("%s (delete %s), %s", AndOrUsers, key, err)
			t.FailNow()
		}
	}
	c.Close()
	if err = ioutil.WriteFile(testUsers, testUsersSrc, 0666); err != nil {
		t.Errorf("%s, %s", testUsers, err)
		t.FailNow()
	}
	if err := LoadUser([]string{testUsers}); err != nil {
		t.Errorf("%s (%s), %s", testUsers, AndOrUsers, err)
		t.FailNow()
	}
}

func TestUserToBytes(t *testing.T) {
	t.Errorf("TestUserToBytes() not implemented")
}

func TestUserToString(t *testing.T) {
	t.Errorf("TestUserToString() not implemented")
}

func TestListUsers(t *testing.T) {
	t.Errorf("TestListUsers() not implemented")
}

func TestRemoveUser(t *testing.T) {
	t.Errorf("TestRemoveUser() not implemented")
}
