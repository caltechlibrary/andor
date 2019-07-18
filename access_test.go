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
	// Three basic workflows for queues draft, review and published
	depositQueue = &Workflow{
		Key:    "writer",
		Name:   "Writer",
		Queue:  "draft",
		Create: true,
		Read:   true,
		Update: true,
		Delete: true,
		AssignTo: []string{
			"review",
		},
	}

	reviewQueue = &Workflow{
		Key:    "reviewer",
		Name:   "Reviewer",
		Queue:  "review",
		Read:   true,
		Update: true,
		AssignTo: []string{
			"draft",
			"published",
		},
	}

	publishedQueue = &Workflow{
		Key:      "published",
		Name:     "Published",
		Queue:    "published",
		Read:     true,
		AssignTo: []string{},
	}

	writer = &User{
		Key:         "writer1",
		DisplayName: "Writer One",
		MemberOf: []string{
			"writer",
		},
	}

	reviewer = &User{
		Key:         "reviewer2",
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

	reviewedObject = map[string]interface{}{
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
	t.Errorf("TestIsAllowed() not implemented")
}

// TestCanAssign tests a user, workflow, queue name and object
// is assignable.
func TestCanAssign(t *testing.T) {
	t.Errorf("TestCanAssign() not implemented.")
}
