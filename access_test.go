// Package andor provides support for building simple digital
// object repositories in Go where objects are stored in a
// dataset collection and the UI of the repository is static
// HTML 5 documents using JavaScript to access a web API.
//
// @Author R. S. Doiel, <rsdoiel@library.caltech.edu>
//
package andor

import (
	"os"
	"path"
	"testing"

	// Caltech Library packages
	"github.com/caltechlibrary/dataset"
)

// Test data for roles
var (
	// Three basic roles for writer, reviewer covering
	// object states of draft, review, published, embargoed
	writerRole = &Role{
		Key:    "writer",
		Name:   "Writer",
		States: []string{"draft"},
		Create: true,
		Read:   true,
		Update: true,
		Delete: true,
		AssignTo: []string{
			"review",
		},
	}

	reviewerRole = &Role{
		Key:    "reviewer",
		Name:   "Reviewer",
		States: []string{"review", "published", "embargoed"},
		Read:   true,
		Update: true,
		Delete: true,
		AssignTo: []string{
			"draft",
			"review",
			"published",
			"embargoed",
		},
	}

	publicRole = &Role{
		Key:      "public",
		Name:     "public",
		States:   []string{"published"},
		Read:     true,
		AssignTo: []string{},
	}

	// Define a couple of users
	writer = &User{
		Key:         "writer",
		DisplayName: "A Writer",
		Roles: []string{
			"writer",
			"public",
		},
	}

	reviewer = &User{
		Key:         "reviewer",
		DisplayName: "A Reviewer",
		Roles: []string{
			"reviewer",
			"public",
		},
	}

	draftObject = map[string]interface{}{
		"_Key":   "1",
		"_State": "draft",
	}

	reviewObject = map[string]interface{}{
		"_Key":   "2",
		"_State": "review",
	}

	publishedObject = map[string]interface{}{
		"_Key":   "3",
		"_State": "published",
	}
)

// TestIsAllowed tests if user, role, permission, and object
// are accessible
func TestIsAllowed(t *testing.T) {
	folder := path.Join("testout", "isallowed")
	cName := path.Join(folder, "is_allowed.ds")
	if _, err := os.Stat(path.Dir(folder)); os.IsNotExist(err) {
		os.MkdirAll(path.Dir(folder), 0777)
	} else {
		if err := os.RemoveAll(folder); err != nil {
			t.Errorf("Can't cleanup statle %q, %s", folder, err)
			t.FailNow()
		}
		os.MkdirAll(path.Dir(folder), 0777)
	}
	c, err := dataset.InitCollection(cName)
	if err != nil {
		t.Errorf("setup failed, %s", err)
		t.FailNow()
	}
	if err = c.Create("1", draftObject); err != nil {
		t.Errorf("Can't create draftObject in %s, %s", cName, err)
		t.FailNow()
	}
	if err = c.Create("2", reviewObject); err != nil {
		t.Errorf("Can't create reviewObject in %s, %s", cName, err)
		t.FailNow()
	}
	if err = c.Create("3", publishedObject); err != nil {
		t.Errorf("Can't create publishedObject in %s, %s", cName, err)
		t.FailNow()
	}

	// Now start testing service isAllowed()
	s := new(AndOrService)
	s.CollectionNames = []string{cName}
	s.Roles = map[string]*Role{
		"writer":   writerRole,
		"reviewer": reviewerRole,
		"public":   publicRole,
	}
	s.States = makeStates(s.Roles)
	s.Users = map[string]*User{
		"writer":   writer,
		"reviewer": reviewer,
	}

	// Writer actions on draft object
	roles, ok := s.getUserRoles("writer")
	if ok == false {
		t.Errorf("Should be able to get the writer's roles")
		t.FailNow()
	}
	if len(roles) != 2 {
		t.Errorf("s.Users -> %+v", s.Users)
		t.Errorf("Expected two roles for A Writer, got %d", len(roles))
	}
	if s.isAllowed(roles, "draft", CREATE) == false {
		t.Errorf("expected true for writer and draftObject create, got false")
	}
	if s.isAllowed(roles, "draft", READ) == false {
		t.Errorf("expected true for writer and draftObject read, got false")
	}
	if s.isAllowed(roles, "draft", UPDATE) == false {
		t.Errorf("expected true for writer and draftObject update, got false")
	}
	if s.isAllowed(roles, "draft", DELETE) == false {
		t.Errorf("expected true for writer and draftObject delete, got false")
	}
	if s.isAllowed(roles, "draft", ASSIGN) == false {
		t.Errorf("expected true for writer and draftObject, AssignTo is NOT empty")
	}
	if s.isAllowed(roles, "review", READ) == true {
		t.Errorf("writer should be restricted from review state")
	}
	if s.isAllowed(roles, "embargoed", READ) == true {
		t.Errorf("writer should be restricted from review state")
	}
	if s.isAllowed(roles, "published", READ) == false {
		t.Errorf("writer should NOT be restricted from public state")
	}

	// Reviewer actions on draft object
	roles, ok = s.getUserRoles("reviewer")
	if ok == false {
		t.Errorf("Should be able to get the reviewer's roles")
		t.FailNow()
	}
	if s.isAllowed(roles, "draft", CREATE) == true {
		t.Errorf("expected false for reviewer and draftObject create")
	}
	if s.isAllowed(roles, "draft", READ) == true {
		t.Errorf("expected false for reviewer and draftObject read")
	}
	if s.isAllowed(roles, "draft", UPDATE) == true {
		t.Errorf("expected false for reviewer and draftObject update")
	}
	if s.isAllowed(roles, "draft", DELETE) == true {
		t.Errorf("expected false for reviewer and draftObject delete")
	}
	if s.isAllowed(roles, "draft", ASSIGN) == true {
		t.Errorf("expected false for reviewer, the draft state isn't listed in states so can't assign from draft, but can assign to draft")
	}
	if s.isAllowed(roles, "review", READ) == false {
		t.Errorf("expected true for reviewer and reviewObject read")
	}
	if s.isAllowed(roles, "review", ASSIGN) == false {
		t.Errorf("expected true for reviewer and reviewObject, can assign")
	}
	if s.isAllowed(roles, "published", READ) == false {
		t.Errorf("expected true for reviewer and publishedObject read")
	}
}

// TestCanAssign tests a user, role, state name and object
// is assignable.
func TestCanAssign(t *testing.T) {
	folder := path.Join("testout", "can_assign")
	cName := path.Join(folder, "is_allowed.ds")
	if _, err := os.Stat(path.Dir(folder)); os.IsNotExist(err) {
		os.MkdirAll(path.Dir(folder), 0777)
	} else {
		if err := os.RemoveAll(folder); err != nil {
			t.Errorf("Can't cleanup stale %q, %s", folder, err)
			t.FailNow()
		}
		os.MkdirAll(path.Dir(folder), 0777)
	}
	c, err := dataset.InitCollection(cName)
	if err != nil {
		t.Errorf("setup failed, %s", err)
		t.FailNow()
	}
	if err = c.Create("1", draftObject); err != nil {
		t.Errorf("Can't create draftObject in %s, %s", cName, err)
		t.FailNow()
	}
	if err = c.Create("2", reviewObject); err != nil {
		t.Errorf("Can't create reviewObject in %s, %s", cName, err)
		t.FailNow()
	}
	if err = c.Create("3", publishedObject); err != nil {
		t.Errorf("Can't create publishedObject in %s, %s", cName, err)
		t.FailNow()
	}

	// Now start testing service canAssign()
	s := new(AndOrService)
	s.CollectionNames = []string{cName}
	s.Roles = map[string]*Role{
		"writer":   writerRole,
		"reviewer": reviewerRole,
		"public":   publicRole,
	}
	s.States = makeStates(s.Roles)
	s.Users = map[string]*User{
		"writer":   writer,
		"reviewer": reviewer,
	}

	roles, ok := s.getUserRoles("writer")
	if ok == false {
		t.Errorf("Should be able to get writer's roles")
		t.FailNow()
	}
	if s.canAssign(roles, "draft", "draft") != false {
		t.Errorf("writer should be restricted from assigning objects, AssignTo is empty")
	}
	if s.canAssign(roles, "draft", "review") != true {
		t.Errorf("writer should be able to assign to review, was restricted")
	}
	if s.canAssign(roles, "draft", "published") != false {
		t.Errorf("writer should be restricted from published, was allowed")
	}
	roles, ok = s.getUserRoles("reviewer")
	if ok == false {
		t.Errorf("Should be able to get reviewer's roles")
		t.FailNow()
	}
	if s.canAssign(roles, "review", "draft") != true {
		t.Errorf("reviewer should be able to assign to draft, was restricted")
	}
	if s.canAssign(roles, "draft", "review") != false {
		t.Errorf("reviewer should be restricted from assign to review, was allowed")
	}
	if s.canAssign(roles, "review", "published") != true {
		t.Errorf("reviewer should be able to assign to published, was resrticted")
	}
}
