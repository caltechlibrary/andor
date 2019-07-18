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
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"

	// Caltech Library Packages

	// Toml package
	"github.com/BurntSushi/toml"
)

// Workflow holds a single workflow description.
// Workflow defines both workflow queue name and
// the permissions about what can be viewed and
// what additional workflows can be assigned to.
type Workflow struct {
	// Key holds the key to be used when referencing the workflow
	// E.g. "editor", "curator", "public"
	Key string `json:"workflow_id" toml:"workflow_id"`
	// Name, the display name, e.g. "Editor", "Curator", "Public View"
	Name string `json:"workflow_name" toml:"workflow_name"`
	// Queues hold name of the queue this workflow can operating on.
	Queue string `json:"queue" toml:"queue"`
	// Create permissions in .Queue
	Create bool `json:"create" toml:"create"`
	// Read permissions in .Queue
	Read bool `json:"read" toml:"read"`
	// Update permissions in .Queue
	Update bool `json:"update" toml:"update"`
	// Delete permissions in .Queue
	Delete bool `json:"delete" toml:"delete"`
	// AssignTo defines a list of queues that this workflow
	// can send objects to.
	AssignTo []string `json:"assign_to" toml:"assign_to"`
}

// LoadWorkflow reads a file (either JSON or TOML) at
// start up of AndOr web service and sets up workflows and
// queues. It returns a map[string]*Workflow,
// a map[string]*Queue and an error
func LoadWorkflow(fName string) (map[string]*Workflow, map[string]*Queue, error) {
	workflows := map[string]*Workflow{}
	queues := map[string]*Queue{}

	// Parse our workflows
	src, err := ioutil.ReadFile(fName)
	if err != nil {
		return nil, nil, err
	}
	switch path.Ext(fName) {
	case ".json":
		if err := json.Unmarshal(src, &workflows); err != nil {
			return nil, nil, err
		}
	case ".toml":
		if _, err := toml.Decode(string(src), &workflows); err != nil {
			return nil, nil, err
		}
	default:
		return nil, nil, fmt.Errorf("workflow must be either a .json or .toml file")
	}
	// Create Queues from workflows.Queue and workflows.AssignTo
	for _, workflow := range workflows {
		// For each queue mentioned in workflow, check if it
		// exists and update it with the workflow information.
		queueList := append([]string{workflow.Name}, workflow.AssignTo...)
		for _, queue := range queueList {
			q, ok := queues[queue]
			if ok == false {
				q := new(Queue)
				q.Name = workflow.Queue
			}
			q.AddWorkflow(workflow.Name)
			queues[queue] = q
		}
	}
	return workflows, queues, nil
}

// Bytes() outputs a workflow to []bytes in TOML.
func (workflow *Workflow) Bytes() []byte {
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(workflow); err != nil {
		src, _ := json.MarshalIndent(workflow, "", "    ")
		return src
	}
	return buf.Bytes()
}

// String() outputs a workflow to a string TOML.
func (workflow *Workflow) String() string {
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(workflow); err != nil {
		src, _ := json.MarshalIndent(workflow, "", "    ")
		return string(src)
	}
	return buf.String()
}
