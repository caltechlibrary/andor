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

// Queue describes a queue's state, the object ideas and
// and sorted id lists and roles associated with the queue.
type Queue struct {
	// Key holds the id of the queue
	Key string `json:"queue_id"`
	// Roles thats operating on this queue
	Roles []string `json:"roles"`
}

// AddRole associates a role with the queue.
func (q *Queue) AddRole(role string) {
	for _, key := range q.Roles {
		if strings.Compare(key, role) == 0 {
			return
		}
	}
	q.Roles = append(q.Roles, role)
}
