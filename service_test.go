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
	"fmt"
	"log"
	"os"
	"path"
	"testing"

	// Caltech Library Packages
	"github.com/caltechlibrary/dataset"
)

// TestRunService creates an *AndOrService and runs both
// it and an http client to test functionality.
func TestRunService(t *testing.T) {
	testFolder := path.Join("testout", "runservice")

	if _, err := os.Stat(testFolder); os.IsNotExist(err) {
		if err := os.MkdirAll(testFolder, 0777); err != nil {
			t.Errorf("Could not recreate %s, %s", testFolder, err)
			t.FailNow()
		}
	} else {
		log.Printf("Recreating %s", testFolder)
		if err := os.RemoveAll(testFolder); err != nil {
			t.Errorf("Could not remove %s, %s", testFolder, err)
			t.FailNow()
		}
		if err := os.MkdirAll(testFolder, 0777); err != nil {
			t.Errorf("Could not recreate %s, %s", testFolder, err)
			t.FailNow()
		}
	}

	andorFile := path.Join(testFolder, "andor.toml")
	cName := path.Join(testFolder, "collection.ds")

	service := new(AndOrService)
	service.CollectionNames = []string{cName}
	service.Htdocs = path.Join(testFolder, "htdocs")
	service.RolesFile = path.Join(testFolder, "roles.toml")
	service.UsersFile = path.Join(testFolder, "users.toml")
	service.AccessFile = path.Join(testFolder, "access.toml")

	// Dump our configured service
	if err := service.DumpAndOr(andorFile); err != nil {
		t.Errorf("Cound not create %s, %s", cName, err)
		t.FailNow()
	}

	// Create our test htdocs folder
	if err := os.MkdirAll(service.Htdocs, 0777); err != nil {
		t.Errorf("Cound not create %s, %s", service.Htdocs, err)
		t.FailNow()
	}

	// Create an empty collection then populate it
	c, err := dataset.InitCollection(cName)
	if err != nil {
		t.Errorf("Cound not create %s, %s", cName, err)
		t.FailNow()
	}
	defer c.Close()

	// Populate our collection with test objects in various
	// states
	objects := []map[string]interface{}{
		map[string]interface{}{
			"title":       "The Short Reign of Pippin IV",
			"description": "A political satire based in 1950s Paris",
			"creators": []map[string]interface{}{
				map[string]interface{}{
					"display_name": "John Steinbeck",
					"family":       "Steinbeck",
					"given":        "John",
				},
			},
			"pubDate": "1957-04-15",
			"_State":  "published",
		},
		map[string]interface{}{
			"title":       "A woman on Avalon",
			"description": "A fictional espionage story set in early 21st century California",
			"creators": []map[string]interface{}{
				map[string]interface{}{
					"display_name": "R. S. Doiel",
					"family":       "Doiel",
					"given":        "R. S.",
				},
			},
			"pubDate": "circa 2010",
			"_State":  "draft",
		},
		map[string]interface{}{
			"title":       "The Adventures of Chet Blacke -- Plastic Man",
			"description": "Comic story about expectations and adventures",
			"creators": []map[string]interface{}{
				map[string]interface{}{
					"display_name": "Richard Weekley",
					"family":       "Weekley",
					"given":        "Richard",
				},
			},
			"pubDate": "1975",
			"_State":  "review",
		},
		map[string]interface{}{
			"title":       "Bandido!",
			"description": "A fictionalized account of Tiburcio Vasquez",
			"creators": []map[string]interface{}{
				map[string]interface{}{
					"display_name": "Luis Valdez",
					"family":       "Valdez",
					"given":        "Luis",
				},
			},
			"pubDate": "1994-06",
			"_State":  "published",
		},
	}
	for i, object := range objects {
		key := fmt.Sprintf("%d", i+100)
		if err := c.Create(key, object); err != nil {
			t.Errorf("could not create object %q, in %q, %s", key, cName, err)
			t.FailNow()
		}
	}

	// Create our test roles.toml
	roles := map[string]*Role{
		"curator": &Role{
			Key:      "curator",
			Name:     "Curator",
			States:   []string{"review", "embargoed", "published"},
			Create:   false,
			Read:     true,
			Update:   true,
			Delete:   true,
			AssignTo: []string{"*"},
		},
		"public": &Role{
			Key:      "public",
			Name:     "Public",
			States:   []string{"published"},
			Create:   false,
			Read:     true,
			Update:   false,
			Delete:   false,
			AssignTo: []string{},
		},
		"depositor": &Role{
			Key:      "depositor",
			Name:     "Depositor",
			States:   []string{"review"},
			Create:   true,
			Read:     false,
			Update:   false,
			Delete:   false,
			AssignTo: []string{},
		},
		"reviewer": &Role{
			Key:      "reviewer",
			Name:     "Reviewer",
			States:   []string{"review"},
			Create:   true,
			Read:     false,
			Update:   false,
			Delete:   false,
			AssignTo: []string{"published", "embargoed"},
		},
	}
	service.Roles = roles
	if err = service.DumpRoles(service.RolesFile); err != nil {
		t.Errorf("could not create %q, %s", path.Base(service.RolesFile), err)
		t.FailNow()
	}

	// Create our test users.toml
	users := map[string]*User{
		"innez": &User{
			Key:         "innez",
			DisplayName: "Innez",
			Roles:       []string{"depositor", "curator"},
		},
		"jane": &User{
			Key:         "jane",
			DisplayName: "Jane",
			Roles:       []string{"depositor", "curator"},
		},
		"millie": &User{
			Key:         "millie",
			DisplayName: "Millie",
			Roles:       []string{"depositor", "reviewer", "public"},
		},
		"bea": &User{
			Key:         "bea",
			DisplayName: "Bea",
			Roles:       []string{"depositor", "public"},
		},
		"anonynous": &User{
			Key:         "anonymous",
			DisplayName: "Anonymous",
			Roles:       []string{"public"},
		},
	}
	service.Users = users
	if err = service.DumpUsers(service.UsersFile); err != nil {
		t.Errorf("could not create %q, %s", path.Base(service.UsersFile), err)
		t.FailNow()
	}

	// Create our test access.toml
	t.Errorf("FIXME: create test access.toml")

	// start service
	t.Errorf("FIXME: need to start service in a Go routine")

	// start client and run tests
	t.Errorf("FIXME: need to start client and run tests")

	// cleanup stop service and clean
	t.Errorf("FIXME: need to stop service and cleanup after tests")
}
