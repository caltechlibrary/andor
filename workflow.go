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
	"sort"

	// Caltech Library Packages
	"github.com/caltechlibrary/dataset"

	// Toml package
	"github.com/BurntSushi/toml"
)

// Workflow holds a single workflow description.
// Workflow defines both workflow queue name and
// the permissions about what can be viewed and
// what additional workflows can be assigned to.
type Workflow struct {
	// Key holds the key to be used when saving the workflow
	// to workflows.AndOr. e.g. "editor", "curator", "public"
	Key string `json:"workflow_id"`
	// Name, the display name, e.g. "Editor", "Curator", "Public View"
	Name string `json:"workflow_name"`
	// Object level permissions, i.e. "create", "read", "update"
	ObjectPermissions []string `json:"object_permissions"`
	// AssignTo defines a list of workflows that this workflow
	// can send objects to.
	AssignTo []string `json:"assign_to"`
	// Queues holds an list of queue names of this workflow
	// can view objects in. The queue name is the same as a
	// defined workflow name. E.g. objects is the review queue
	// would be in the review workflow state with those rights.
	Queues []string `json:"queues"`
}

// LoadWorkflow takes a filename, reads the file
// (either JSON or TOML) and updates the workflow.AndOr
// collection.
func LoadWorkflow(fNames []string) error {
	c, err := dataset.Open(andOrWorkflows)
	if err != nil {
		return err
	}
	defer c.Close()

	for _, fName := range fNames {
		workflows := map[string]*Workflow{}
		src, err := ioutil.ReadFile(fName)
		if err != nil {
			return err
		}
		switch path.Ext(fName) {
		case ".json":
			if err := json.Unmarshal(src, &workflows); err != nil {
				return err
			}
		case ".toml":
			if _, err := toml.Decode(string(src), &workflows); err != nil {
				return err
			}
		default:
			return fmt.Errorf("workflow must be either a .json or .toml file")
		}

		for key, workflow := range workflows {
			workflow.Key = key
			src, err := json.MarshalIndent(workflow, "", "    ")
			if err != nil {
				return err
			}
			if c.HasKey(key) {
				err = c.UpdateJSON(key, src)
			} else {
				err = c.CreateJSON(key, src)
			}
			if err != nil {
				return nil
			}
		}
	}
	return nil
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

// ListWorkflow returns a list of workflow objects
func ListWorkflow(keys []string) ([]*Workflow, error) {
	c, err := dataset.Open(andOrWorkflows)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	if len(keys) == 0 {
		keys = c.Keys()
	}
	sort.Strings(keys)
	objects := []*Workflow{}
	for _, key := range keys {
		src, err := c.ReadJSON(key)
		if err != nil {
			return nil, err
		}
		obj := new(Workflow)
		err = json.Unmarshal(src, &obj)
		if err != nil {
			return nil, err
		}
		objects = append(objects, obj)
	}
	return objects, nil
}

// RemoveWorkflow removes one or more workflows from workflows.AndOr
func RemoveWorkflow(workflowNames []string) error {
	c, err := dataset.Open(andOrWorkflows)
	if err != nil {
		return err
	}
	defer c.Close()
	for _, key := range workflowNames {
		if err := c.Delete(key); err != nil {
			return err
		}
	}
	return nil
}
