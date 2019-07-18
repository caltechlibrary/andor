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
	"strings"
)

const (
	_ = iota
	CREATE
	READ
	UPDATE
	DELETE
)

// IsAllowed tests if CRUD operations can be taken on object
// based on user, object  and a permission.
// It returns true if permission is affirmed false otherwise.
func (s *AndOrService) IsAllowed(user *User, object map[string]interface{}, permission int) bool {
	// Get the object's queue
	queue := ObjectQueue(object)
	if q, ok := s.Queues[queue]; ok == true {
		// Get the queue's associated workflow(s)
		for _, workflowName := range q.Workflows {
			// Check if user is in a workflow associated with queue
			if user.IsMemberOf(workflowName) {
				// Get workflow
				if workflow, ok := s.Workflows[workflowName]; ok {
					// Check workflow permission requested
					switch permission {
					case CREATE:
						return workflow.Create
					case READ:
						return workflow.Read
					case UPDATE:
						return workflow.Update
					case DELETE:
						return workflow.Delete
					}
				}

			}
		}
	}
	return false
}

// CanAssign takes a user, object and queue target and checks if the
// asignment is allowed.
func (s *AndOrService) CanAssign(user *User, object map[string]interface{}, targetQueue string) bool {
	// Get the object's queue
	queue := ObjectQueue(object)
	if q, ok := s.Queues[queue]; ok == true {
		// Get the queue's associated workflow(s)
		for _, workflowName := range q.Workflows {
			// Check if user is in a workflow associated with queue
			if user.IsMemberOf(workflowName) {
				// Check what workflow assignments are allowed.
				if workflow, ok := s.Workflows[workflowName]; ok {
					for _, queueName := range workflow.AssignTo {
						if strings.Compare(queueName, targetQueue) == 0 {
							return true
						}
					}
				}
			}
		}
	}
	return false
}
