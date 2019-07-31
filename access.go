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
	// Get the object's state
	state := ObjectState(object)
	if q, ok := s.States[state]; ok == true {
		// Get the state's associated role(s)
		for _, roleName := range q.Roles {
			// Check if user is in a role associated with state
			if user.IsMemberOf(roleName) {
				// Get role
				if role, ok := s.Roles[roleName]; ok {
					// Check role permission requested
					switch permission {
					case CREATE:
						return role.Create
					case READ:
						return role.Read
					case UPDATE:
						return role.Update
					case DELETE:
						return role.Delete
					}
				}

			}
		}
	}
	return false
}

// CanAssign takes a user, object and state target and checks if the
// asignment is allowed.
func (s *AndOrService) CanAssign(user *User, object map[string]interface{}, targetState string) bool {
	// Get the object's state
	state := ObjectState(object)
	if q, ok := s.States[state]; ok == true {
		// Get the state's associated role(s)
		for _, roleName := range q.Roles {
			// Check if user is in a role associated with state
			if user.IsMemberOf(roleName) {
				// Check what role assignments are allowed.
				if role, ok := s.Roles[roleName]; ok {
					for _, stateName := range role.AssignTo {
						if strings.Compare(stateName, targetState) == 0 {
							return true
						}
					}
				}
			}
		}
	}
	return false
}
