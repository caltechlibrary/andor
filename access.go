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
	"net/http"
	"strings"
)

const (
	//
	// These define the operations that can be performed on an Object
	//
	UNDEFINED = iota
	// CREATE an object with the requested state
	CREATE
	// READ an object with the requested state
	READ
	// UPDATE an object with the requested state
	UPDATE
	// DELETE an object with the requested state
	DELETE
	// ASSIGN an object to the requested next state
	ASSIGN
)

// getState takes a map[string]interface{} and returns the
// ._State value as a string or an empty string if not found.
func getState(object map[string]interface{}) string {
	if state, ok := object["_State"]; ok == true {
		switch state.(type) {
		case string:
			return state.(string)
		}
	}
	return ""
}

// getOp looks at q request and returns the operation (CRUD)
// requested.
func getOp(r *http.Request) int {
	p := strings.Split(r.URL.Path, "/")
	// NOTE: we should have more than three path elements.
	// empty, collection name, operation, object ids
	if len(p) < 3 {
		return UNDEFINED
	}
	switch p[2] {
	case "create":
		return CREATE
	case "read":
		return READ
	case "update":
		return UPDATE
	case "delete":
		return DELETE
	case "assign":
		return ASSIGN
	default:
		return UNDEFINED
	}
}

// getUsername looks at a request and determines who the
// user is (e.g. via BasicAUTH or JWT). Returns a string
// or error.
func (s *AndOrService) getUsername(r *http.Request) string {
	//FIXME: This should should also handle JSON Web Token based
	// authentication/assertions ...
	username, _, ok := r.BasicAuth()
	if ok == false {
		username = ""
	}
	return strings.TrimSpace(username)
}

// getUserRoles looks at a request and determines who the
// user is and retrieves their roles. It returns a list of
// roles that can be evaluated by hasPermission() or error
func (s *AndOrService) getUserRoles(username string) (map[string]*Role, bool) {
	u, ok := s.Users[username]
	if ok == true {
		roleMap := make(map[string]*Role)
		// Is user member of role?
		for _, key := range u.Roles {
			if role, ok := s.Roles[key]; ok == true {
				roleMap[key] = role
			}
		}
		if len(roleMap) > 0 {
			return roleMap, true
		}
	}
	return nil, false
}

// isAllowed applies the policies defined in user's roles,
// object state and operation requested.  NOTE: if the operation
// is ASSIGN then an additional check is needed with
// CanAssign which takes roles, current object state, next object
// state.
//
// if roles, ok := s.getUserRoles("jane"); ok == true {
//     if s.IsAllowed(roles, "review", CREATE) {
//         ... user allowed to create object in review. ...
//     }
// }
func (s *AndOrService) isAllowed(roles map[string]*Role, state string, op int) bool {
	for _, role := range roles {
		for _, rState := range role.States {
			if strings.Compare(state, rState) == 0 || strings.Compare(rState, "*") == 0 {
				switch op {
				case CREATE:
					if role.Create {
						return true
					}
				case READ:
					if role.Read {
						return true
					}
				case UPDATE:
					if role.Update {
						return true
					}
				case DELETE:
					if role.Delete {
						return true
					}
				case ASSIGN:
					if len(role.AssignTo) > 0 {
						return true
					}
				}
			}
		}
	}
	return false
}

// canAssign takes a map of roles, a current state, and a next state
// and returns true is assignment is permitted, false otherwise.
func (s *AndOrService) canAssign(roles map[string]*Role, current string, next string) bool {
	for _, role := range roles {
		// Can role make assignments?
		if len(role.AssignTo) > 0 {
			// Is the object in a state covered by role?
			for _, state := range role.States {
				if strings.Compare(state, current) == 0 {
					// Can that state assign to next?
					for _, assignment := range role.AssignTo {
						if strings.Compare(assignment, next) == 0 {
							return true
						}
					}
				}
			}
		}
	}
	return false
}
