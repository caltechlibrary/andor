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
	"testing"
)

// Test data for workflows
var (
	// Three basic workflows
	deposit = &Workflow{
		Key:               "deposit",
		Name:              "Deposit",
		ObjectPermissions: []string{"create", "read", "update"},
		AssignTo:          []string{"review"},
		Queues:            []string{"deposit"},
	}

	review = &Workflow{
		Key:               "review",
		Name:              "Review",
		ObjectPermissions: []string{"read", "update"},
		AssignTo:          []string{"deposit", "published"},
		Queues:            []string{"review", "published"},
	}

	published = &Workflow{
		Key:               "published",
		Name:              "Published",
		ObjectPermissions: []string{"read"},
		AssignTo:          []string{},
		Queues:            []string{"published"},
	}

	depositor = &User{
		Key:         "UserOne",
		DisplayName: "User One",
		CreateQueue: "deposit",
		MemberOf:    []string{"deposit"},
	}

	reviewer = &User{
		Key:         "UserTwo",
		DisplayName: "User Tow",
		CreateQueue: "",
		MemberOf:    []string{"review", "published"},
	}

	depositedObject = map[string]interface{}{
		"_Key":   "1",
		"_Queue": "deposit",
	}
	reviewedObject = map[string]interface{}{
		"_Key":   "2",
		"_Queue": "review",
	}
	publishedObject = map[string]interface{}{
		"_Key":   "3",
		"_Queue": "published",
	}
)

// TestUserInWorkflow tests user in workflow
func TestUserInWorkflow(t *testing.T) {
	if expected, got := true, UserInWorkflow(depositor, deposit); expected != got {
		t.Errorf("expected %t, got %t", expected, got)

	}
	if expected, got := false, UserInWorkflow(depositor, review); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
	if expected, got := false, UserInWorkflow(depositor, published); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}

	if expected, got := false, UserInWorkflow(reviewer, deposit); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
	if expected, got := true, UserInWorkflow(reviewer, review); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
	if expected, got := true, UserInWorkflow(reviewer, published); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
}

// TestObjectInWorkflow test object and workflow
func TestObjectInWorkflow(t *testing.T) {
	if expected, got := true, ObjectInWorkflow(depositedObject, deposit); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
	if expected, got := false, ObjectInWorkflow(depositedObject, review); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
	if expected, got := false, ObjectInWorkflow(depositedObject, published); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
	if expected, got := false, ObjectInWorkflow(reviewedObject, deposit); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
	if expected, got := true, ObjectInWorkflow(reviewedObject, review); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
	if expected, got := false, ObjectInWorkflow(reviewedObject, published); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
	if expected, got := false, ObjectInWorkflow(publishedObject, deposit); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
	if expected, got := true, ObjectInWorkflow(publishedObject, review); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
	if expected, got := true, ObjectInWorkflow(publishedObject, published); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
}

// TestCanAccess tests if user, workflow, permission, and object
// are accessible
func TestCanAccess(t *testing.T) {
	t.Errorf("TestCanAccess() not implemented.")
}

// TestCanAssign tests a user, workflow, queue name and object
// is assignable.
func TestCanAssign(t *testing.T) {
	t.Errorf("TestCanAssign() not implemented.")
}
