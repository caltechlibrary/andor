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
	"encoding/json"
	"fmt"
)

// Queue describes a queue's state, the object ideas and
// and sorted id lists and workflows associated with the queue.
type Queue struct {
	// Name holds the name of the queue
	Name string `json:"name"`
	// Workflows thats operating on this queue
	Workflows []string `json:"workflows"`
}

// AddWorkflow associates a workflow with the queue.
func (q *Queue) AddWorkflow(workflow string) {
	hasWorkflow := false
	for _, key := range q.Workflows {
		if key == workflow {
			hasWorkflow = true
			break
		}
	}
	if hasWorkflow == false {
		q.Workflows = append(q.Workflows, workflow)
	}
}

// objKey inspects a map[string]interface{} (an object)
// and returns the `._Key` value if it is set, otherwise an
// empty string.
func objKey(obj map[string]interface{}) string {
	if key, ok := obj["_Key"]; ok == true {
		switch key.(type) {
		case json.Number:
			return key.(json.Number).String()
		case float64:
			return fmt.Sprintf("%f", key.(float64))
		case int:
			return fmt.Sprintf("%d", key.(int))
		case string:
			return key.(string)
		}
	}
	return ""
}
