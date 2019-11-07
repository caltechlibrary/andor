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
	"flag"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"testing"

	// Caltech Library Packages
	"github.com/caltechlibrary/dataset"
	"github.com/caltechlibrary/wsfn"
)

var (
	runWebService bool
	service       *AndOrService
)

// TestRunService creates an *AndOrService and runs both
// it and an http client to test functionality.
func TestRunService(t *testing.T) {
	username := "bea"
	roles, ok := service.getUserRoles(username)
	if ok == false {
		t.Errorf("Expected user bea to have roles")
		t.FailNow()
	}
	for rName, role := range roles {
		switch rName {
		case "depositor":
			if (len(role.States) == 1 &&
				role.States[0] == "review" &&
				role.Create == true &&
				role.Read == false &&
				role.Update == false &&
				role.Delete == false &&
				len(role.AssignTo) == 0) == false {
				t.Errorf("%s's role %q has unexpected permissions, %s", username, rName, role)
				t.FailNow()

			}
		case "public":
			if (len(role.States) == 1 &&
				role.States[0] == "published" &&
				role.Create == false &&
				role.Read == true &&
				role.Update == false &&
				role.Delete == false &&
				len(role.AssignTo) == 0) == false {
				t.Errorf("%s's role %q has unexpected permissions, %s", username, rName, role)
				t.FailNow()

			}
		default:
			t.Errorf("%s's role %q is not depositor or public", username, rName)
		}
	}

	cName := service.CollectionNames[0]
	c, err := dataset.Open(cName)
	if err != nil {
		t.Errorf("Could note open %q, %s", cName, err)
		t.FailNow()
	}
	defer c.Close()
	key := "101"
	object := make(map[string]interface{})

	if err = c.Read(key, object, false); err != nil {
		t.Errorf("Expected to read object %q from %q, %s", key, c.Name, err)
		t.FailNow()
	}
	state := getState(object)
	if strings.Compare(state, "draft") != 0 {
		t.Errorf("Expected object (%q) in draft, got %q", key, state)
		t.FailNow()
	}
	if service.isAllowed(roles, "draft", READ) == true {
		t.Errorf("bea should not be able to read drafts")
		t.FailNow()
	}

	// start service
	if runWebService {
		log.Println("Starting service for testing")
		if err := service.Start(); err != nil {
			log.Fatalf("service.Start() failed, %s", err)
		}

		// create client
		t.Errorf("FIXME: need to create client and run tests")

		// now run some tests.
		t.Errorf("FIXME: TestRun(t *testing.T) not implemented.")
	}
}

// TestMain creates an *AndOrService and populates with a test
// dataset for testing.
func TestMain(m *testing.M) {
	testFolder := path.Join("testout", "runservice")

	if _, err := os.Stat(testFolder); os.IsNotExist(err) {
		if err := os.MkdirAll(testFolder, 0777); err != nil {
			log.Fatalf("Could not recreate %s, %s", testFolder, err)
		}
	} else {
		log.Printf("Recreating %s", testFolder)
		if err := os.RemoveAll(testFolder); err != nil {
			log.Fatalf("Could not remove %s, %s", testFolder, err)
		}
		if err := os.MkdirAll(testFolder, 0777); err != nil {
			log.Fatalf("Could not recreate %s, %s", testFolder, err)
		}
	}

	andorFile := path.Join(testFolder, "andor.toml")
	cName := path.Join(testFolder, "collection.ds")

	service = new(AndOrService)
	service.Scheme = "http"
	service.Host = "localhost"
	service.Port = "8246"
	service.Htdocs = path.Join(testFolder, "htdocs")
	service.CollectionNames = []string{cName}
	service.RolesFile = path.Join(testFolder, "roles.toml")
	service.UsersFile = path.Join(testFolder, "users.toml")
	service.AccessFile = path.Join(testFolder, "access.toml")

	// Dump our configured service
	if err := service.DumpAndOr(andorFile); err != nil {
		log.Fatalf("Cound not create %s, %s", cName, err)
	}

	// Create our test htdocs folder
	if err := os.MkdirAll(service.Htdocs, 0777); err != nil {
		log.Fatalf("Cound not create %s, %s", service.Htdocs, err)
	}

	// Create an empty collection then populate it
	c, err := dataset.InitCollection(cName)
	if err != nil {
		log.Fatalf("Cound not create %s, %s", cName, err)
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
			"title":       "The Adventures of Chet Blake -- Plastic Man",
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
			log.Fatalf("could not create object %q, in %q, %s", key, cName, err)
		}
	}

	// Create our test roles.toml
	roles := map[string]*Role{
		"publisher": &Role{
			Key:      "publisher",
			Name:     "Publisher",
			States:   []string{"*"},
			Create:   false,
			Read:     true,
			Update:   true,
			Delete:   true,
			AssignTo: []string{"*"},
		},
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
		log.Fatalf("could not create %q, %s", path.Base(service.RolesFile), err)
	}

	// Create our test users.toml
	users := map[string]*User{
		"ester": &User{
			Key:         "ester",
			DisplayName: "Ester",
			Roles:       []string{"publisher"},
		},
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
		log.Fatalf("could not create %q, %s", path.Base(service.UsersFile), err)
	}

	// Create our test access.toml
	service.Access = new(wsfn.Access)
	service.Access.AuthType = "basic"
	service.Access.AuthName = "Testing RunService()"
	if service.IsAccessRestricted() == false {
		log.Fatalf("service should be access restricted!")
	}
	for _, user := range service.Users {
		// Create our access accounts with common password for testing.
		service.Access.UpdateAccess(user.Key, "hello")
	}
	if err = service.DumpAccess(service.AccessFile); err != nil {
		log.Fatalf("could not create %q, %s", path.Base(service.AccessFile), err)
	}

	// call flag.Parse() here if TestMain uses flags.
	flag.BoolVar(&runWebService, "web", false, "run web service")
	flag.Parse()
	os.Exit(m.Run())
}
