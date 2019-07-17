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
	deposit = Workflow{
		Key:               "deposit",
		Name:              "Deposit",
		ObjectPermissions: []string{"create", "update"},
		AssignTo:          []string{"review"},
		Queues:            []string{"deposit"},
	}

	review = Workflow{
		Key:               "review",
		Name:              "Review",
		ObjectPermissions: []string{"read", "update"},
		AssignTo:          []string{"deposit", "published"},
		Queues:            []string{"deposit", "review", "published"},
	}
	published = Workflow{
		Key:               "published",
		Name:              "Published",
		ObjectPermissions: []string{"read"},
		AssignTo:          []string{},
		Queues:            []string{"published"},
	}

	user1 = User{
		Key:         "UserOne",
		DisplayName: "User One",
		CreateQueue: "deposit",
		MemberOf:    []string{"deposit", "review", "published"},
	}

	user2 = User{
		Key:         "UserTwo",
		DisplayName: "User Tow",
		CreateQueue: "",
		MemberOf:    []string{"published"},
	}
)

// TestUserInWorkflow tests user in workflow
func TestUserInWorkflow(t *testing.T) {
	if expected, got := true, UserInWorkflow(user1, deposit); expected != got {
		t.Errorf("expected %t, got %t", expected, got)

	}
	if expected, got := true, UserInWorkflow(user1, review); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
	if expected, got := true, UserInWorkflow(user1, published); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}

	if expected, got := false, UserInWorkflow(user2, deposit); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
	if expected, got := false, UserInWorkflow(user2, review); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
	if expected, got := true, UserInWorkflow(user2, published); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
}

// TestObjectInWorkflow test object and workflow
func TestObjectInWorkflow(t *testing.T) {
	t.Errorf("TestObjectInWorkflow() not implemented.")
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
