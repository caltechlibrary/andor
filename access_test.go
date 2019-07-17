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
	depositQueue = &Workflow{
		Key:               "deposit",
		Name:              "Deposit",
		ObjectPermissions: []string{"create", "read", "update"},
		AssignTo:          []string{"review"},
		Queues:            []string{"deposit"},
	}

	reviewQueue = &Workflow{
		Key:               "review",
		Name:              "Review",
		ObjectPermissions: []string{"read", "update"},
		AssignTo:          []string{"deposit", "published"},
		Queues:            []string{"review", "published"},
	}

	publishedQueue = &Workflow{
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
	if expected, got := true, UserInWorkflow(depositor, depositQueue); expected != got {
		t.Errorf("expected %t, got %t", expected, got)

	}
	if expected, got := false, UserInWorkflow(depositor, reviewQueue); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
	if expected, got := false, UserInWorkflow(depositor, publishedQueue); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}

	if expected, got := false, UserInWorkflow(reviewer, depositQueue); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
	if expected, got := true, UserInWorkflow(reviewer, reviewQueue); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
	if expected, got := true, UserInWorkflow(reviewer, publishedQueue); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
}

// TestObjectInWorkflow test object and workflow
func TestObjectInWorkflow(t *testing.T) {
	if expected, got := true, ObjectInWorkflow(depositedObject, depositQueue); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
	if expected, got := false, ObjectInWorkflow(depositedObject, reviewQueue); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
	if expected, got := false, ObjectInWorkflow(depositedObject, publishedQueue); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
	if expected, got := false, ObjectInWorkflow(reviewedObject, depositQueue); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
	if expected, got := true, ObjectInWorkflow(reviewedObject, reviewQueue); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
	if expected, got := false, ObjectInWorkflow(reviewedObject, publishedQueue); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
	if expected, got := false, ObjectInWorkflow(publishedObject, depositQueue); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
	if expected, got := true, ObjectInWorkflow(publishedObject, reviewQueue); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
	if expected, got := true, ObjectInWorkflow(publishedObject, publishedQueue); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
}

// TestIsAllowed tests if user, workflow, permission, and object
// are accessible
func TestIsAllowed(t *testing.T) {
	if expected, got := true, IsAllowed(depositor, depositQueue, depositedObject, "create"); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
	if expected, got := true, IsAllowed(depositor, depositQueue, depositedObject, "read"); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
	if expected, got := true, IsAllowed(depositor, depositQueue, depositedObject, "update"); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
	if expected, got := false, IsAllowed(depositor, reviewQueue, depositedObject, "create"); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
	if expected, got := false, IsAllowed(depositor, reviewQueue, depositedObject, "read"); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
	if expected, got := false, IsAllowed(depositor, reviewQueue, depositedObject, "update"); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
	if expected, got := false, IsAllowed(depositor, publishedQueue, depositedObject, "create"); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
	if expected, got := false, IsAllowed(depositor, publishedQueue, depositedObject, "read"); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
	if expected, got := false, IsAllowed(depositor, publishedQueue, depositedObject, "update"); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
}

// TestCanAssign tests a user, workflow, queue name and object
// is assignable.
func TestCanAssign(t *testing.T) {
	if expected, got := true, CanAssign(depositor, depositQueue, depositedObject, "review"); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
	if expected, got := false, CanAssign(depositor, depositQueue, depositedObject, "deposit"); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
	if expected, got := false, CanAssign(depositor, depositQueue, depositedObject, "published"); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
	if expected, got := false, CanAssign(depositor, reviewQueue, depositedObject, "deposit"); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
	if expected, got := false, CanAssign(depositor, reviewQueue, depositedObject, "review"); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
	if expected, got := false, CanAssign(depositor, reviewQueue, depositedObject, "published"); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
	if expected, got := false, CanAssign(depositor, publishedQueue, depositedObject, "deposit"); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
	if expected, got := false, CanAssign(depositor, publishedQueue, depositedObject, "review"); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
	if expected, got := false, CanAssign(depositor, publishedQueue, depositedObject, "published"); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}

	if expected, got := false, CanAssign(reviewer, depositQueue, reviewedObject, "review"); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
	if expected, got := false, CanAssign(reviewer, depositQueue, reviewedObject, "deposit"); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
	if expected, got := false, CanAssign(reviewer, depositQueue, reviewedObject, "published"); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
	if expected, got := true, CanAssign(reviewer, reviewQueue, reviewedObject, "deposit"); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
	if expected, got := false, CanAssign(reviewer, reviewQueue, reviewedObject, "review"); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
	if expected, got := true, CanAssign(reviewer, reviewQueue, reviewedObject, "published"); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
	if expected, got := false, CanAssign(reviewer, publishedQueue, reviewedObject, "deposit"); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
	if expected, got := false, CanAssign(reviewer, publishedQueue, reviewedObject, "review"); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
	if expected, got := false, CanAssign(reviewer, publishedQueue, reviewedObject, "published"); expected != got {
		t.Errorf("expected %t, got %t", expected, got)
	}
}
