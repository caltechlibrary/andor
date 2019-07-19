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

// Test data for workflows
var (
	// Three basic workflows for queues draft, review and published
	draftWorkflow = &Workflow{
		Key:    "draft",
		Name:   "Draft",
		Queue:  "draft",
		Create: true,
		Read:   true,
		Update: true,
		Delete: true,
		AssignTo: []string{
			"review",
		},
	}

	reviewWorkflow = &Workflow{
		Key:    "review",
		Name:   "Review",
		Queue:  "review",
		Read:   true,
		Update: true,
		AssignTo: []string{
			"draft",
			"published",
		},
	}

	publishedWorkflow = &Workflow{
		Key:      "published",
		Name:     "Published",
		Queue:    "published",
		Read:     true,
		AssignTo: []string{},
	}

	writer = &User{
		Key:         "writer",
		DisplayName: "Writer One",
		MemberOf: []string{
			"draft",
		},
	}

	reviewer = &User{
		Key:         "reviewer",
		DisplayName: "Reviewer Two",
		MemberOf: []string{
			"review",
			"published",
		},
	}

	draftObject = map[string]interface{}{
		"_Key":   "1",
		"_Queue": "draft",
	}

	reviewObject = map[string]interface{}{
		"_Key":   "2",
		"_Queue": "review",
	}

	publishedObject = map[string]interface{}{
		"_Key":   "3",
		"_Queue": "published",
	}
)

// TestIsAllowed tests if user, workflow, permission, and object
// are accessible
func TestIsAllowed(t *testing.T) {
	cName := path.Join("testout", "is_allowed.ds")
	if _, err := os.Stat(path.Dir(cName)); os.IsNotExist(err) {
		os.MkdirAll(path.Dir(cName), 0777)
	}
	if _, err := os.Stat(cName); os.IsNotExist(err) == false {
		os.RemoveAll(cName)
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

	// Now start testing service IsAllowed()
	s := new(AndOrService)
	s.Repositories = []string{cName}
	s.Workflows = map[string]*Workflow{
		"draft":     draftWorkflow,
		"review":    reviewWorkflow,
		"published": publishedWorkflow,
	}
	s.Queues = makeQueues(s.Workflows)
	s.Users = map[string]*User{
		"writer":   writer,
		"reviewer": reviewer,
	}

	// Writer actions on draft object
	if s.IsAllowed(writer, draftObject, CREATE) != true {
		t.Errorf("expected true for writer and draftObject create, got false")
	}
	if s.IsAllowed(writer, draftObject, READ) != true {
		t.Errorf("expected true for writer and draftObject read, got false")
	}
	if s.IsAllowed(writer, draftObject, UPDATE) != true {
		t.Errorf("expected true for writer and draftObject update, got false")
	}
	if s.IsAllowed(writer, draftObject, DELETE) != true {
		t.Errorf("expected true for writer and draftObject delete, got false")
	}
	// Reviewer actions on draft object
	if s.IsAllowed(reviewer, draftObject, CREATE) != false {
		t.Errorf("expected false for reviewer and draftObject create, got true")
	}
	if s.IsAllowed(reviewer, draftObject, READ) != true {
		t.Errorf("expected true for reviewer and draftObject read, got false")
	}
	if s.IsAllowed(reviewer, draftObject, UPDATE) != true {
		t.Errorf("expected true for reviewer and draftObject update, got false")
	}
	if s.IsAllowed(reviewer, draftObject, DELETE) != false {
		t.Errorf("expected false for reviewer and draftObject delete, got true")
	}
}

// TestCanAssign tests a user, workflow, queue name and object
// is assignable.
func TestCanAssign(t *testing.T) {
	cName := path.Join("testout", "is_allowed.ds")
	if _, err := os.Stat(path.Dir(cName)); os.IsNotExist(err) {
		os.MkdirAll(path.Dir(cName), 0777)
	}
	if _, err := os.Stat(cName); os.IsNotExist(err) == false {
		os.RemoveAll(cName)
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

	// Now start testing service IsAllowed()
	s := new(AndOrService)
	s.Repositories = []string{cName}
	s.Workflows = map[string]*Workflow{
		"draft":     draftWorkflow,
		"review":    reviewWorkflow,
		"published": publishedWorkflow,
	}
	s.Queues = makeQueues(s.Workflows)
	s.Users = map[string]*User{
		"writer":   writer,
		"reviewer": reviewer,
	}

	if s.CanAssign(writer, draftObject, "draft") != false {
		t.Errorf("writer should be restricted from assigning to draft, was allowed")
	}
	if s.CanAssign(writer, draftObject, "review") != true {
		t.Errorf("writer should be able to assign to review, was restricted")
	}
	if s.CanAssign(writer, draftObject, "published") != false {
		t.Errorf("writer should be restricted from published, was allowed")
	}
	if s.CanAssign(reviewer, draftObject, "draft") != true {
		t.Errorf("reviewer should be able to assign to draft, was restricted")
	}
	if s.CanAssign(reviewer, draftObject, "review") != false {
		t.Errorf("reviewer should be restricted from assign to review, was allowed")
	}
	if s.CanAssign(reviewer, draftObject, "published") != true {
		t.Errorf("reviewer should be able to assign to published, was resrticted")
	}
}
