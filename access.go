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
	"os"
	"strings"
)

var debugOn = false

func debug_msg(fString string, args ...interface{}) {
	if debugOn {
		fmt.Fprintf(os.Stderr, "DEBUG "+fString+"\n", args...)
	}
}

// UserInWorkflow takes user and workflow and sees if
// the user is indeed in the workflow or not.
func UserInWorkflow(user *User, workflow *Workflow) bool {
	debug_msg("user (%T) -> %+v\n", user, workflow)
	debug_msg("Workflow (%T) -> %+v\n", workflow, workflow.String())
	for _, queue := range user.MemberOf {
		debug_msg("compare(%q, %q)\n\n", queue, workflow.Key)
		if strings.Compare(queue, workflow.Key) == 0 {
			debug_msg("matched\n\n")
			return true
		}
	}
	return false
}

// ObjectInWorkflow takes an object and workflow and sees if
// the object is in the workflow's queue(s)
func ObjectInWorkflow(object map[string]interface{}, workflow *Workflow) bool {
	if s, ok := object["_Queue"]; ok == true {
		debug_msg("s (%T) -> %+v\n", s, s)
		debug_msg("Workflow (%T) -> %+v\n", workflow, workflow.String())
		switch s.(type) {
		case string:
			queueName := s.(string)
			for _, queue := range workflow.Queues {
				debug_msg("compare(%q, %q)\n\n", queueName, queue)
				if strings.Compare(queueName, queue) == 0 {
					debug_msg("matched\n\n")
					return true
				}
			}
		}
	}
	debug_msg("NOT matched\n\n")
	return false
}

// CanAccess takes a user, workflow, permission, and object
// it returns true if permission is affirmed false otherwise.
func CanAccess(user *User, workflow *Workflow, permission string, object map[string]interface{}) bool {
	// Check if user is in workflow
	// Check if object is in workflow's queues
	if UserInWorkflow(user, workflow) && ObjectInWorkflow(object, workflow) {
		// Check if work flow has rights on object
		for _, objectPermission := range workflow.ObjectPermissions {
			if strings.Compare(objectPermission, "*") == 0 ||
				strings.Compare(permission, objectPermission) == 0 {
				return true
			}
		}
		return false
	}

	return false
}

// CanAssign takes a user, workflow, queue name and object
// it returns true if assignment is allowed, false otherwise
func CanAssign(user *User, workflow *Workflow, queueName string, object map[string]interface{}) bool {
	// Check if user is in workflow
	// Check if object is in workflow's queues
	if UserInWorkflow(user, workflow) && ObjectInWorkflow(object, workflow) {
		// Check if work flow has rights to assign to new queue
		for _, assignTo := range workflow.AssignTo {
			if strings.Compare(assignTo, "*") == 0 ||
				strings.Compare(queueName, assignTo) == 0 {
				return true
			}
		}
		return false
	}
	return false
}
