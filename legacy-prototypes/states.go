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

// State describes a state's state, the object ideas and
// and sorted id lists and roles associated with the state.
type State struct {
	// Key holds the id of the state
	Key string `json:"state_id"`
	// Roles thats operating on this state
	Roles []string `json:"roles"`
}

// AddRole associates a role with the state.
func (s *State) AddRole(role string) {
	for _, key := range s.Roles {
		if strings.Compare(key, role) == 0 {
			return
		}
	}
	s.Roles = append(s.Roles, role)
}
