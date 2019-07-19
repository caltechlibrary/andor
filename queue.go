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
// and sorted id lists and workflows associated with the queue.
type Queue struct {
	// Key holds the id of the queue
	Key string `json:"queue_id"`
	// Workflows thats operating on this queue
	Workflows []string `json:"workflows"`
}

// AddWorkflow associates a workflow with the queue.
func (q *Queue) AddWorkflow(workflow string) {
	for _, key := range q.Workflows {
		if strings.Compare(key, workflow) == 0 {
			return
		}
	}
	q.Workflows = append(q.Workflows, workflow)
}
