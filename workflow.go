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
	AssignTo string `json:"assign_to"`
	// Queues holds an list of queue names of this workflow
	// can view objects in. The queue name is the same as a
	// defined workflow name. E.g. objects is the review queue
	// would be in the review workflow state with those rights.
	Queues string `json:"queues"`
}

// ReadWorkflowFile takes a filename, reads the file
// (either JSON or TOML) and returns a workflow
// object and error.
func ReadWorkflowFile(fName string) (*Workflow, error) {
	workflow := new(Workflow)
	src, err := ioutil.ReadFile(fName)
	if err != nil {
		return nil, err
	}
	switch path.Exit(fName) {
	case ".json":
		if err := json.Unmarshal(src, &workflow); err != nil {
			return workflow, err
		}
	case ".toml":
		if _, err := toml.Decode(src, &workflow); err != nil {
			return workflow, err
		}
	default:
		return nil, fmt.Errorf("workflow must be either a .json or .toml file")
	}
	return workflow, nil
}

// Bytes() outputs a workflow to []bytes in TOML.
func (workflow *Workflow) Bytes() []byte {
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(workflow); err != nil {
		return fmt.Sprintf("%+v", workflow)
	}
	return buf
}

// String() outputs a workflow to a string TOML.
func (workflow *Workflow) String() string {
	return workflow.Bytes().String()
}

// AddWorkflow adds a workflow to the "workflows.AndOr"
// dataset collection.
func AddWorkflow(worflowName string, workflow *Workflow) error {
	c, err := dataset.Open("workflows.AndOr")
	if err != nil {
		return err
	}
	defer c.Close()
	src, err := json.MarshalIndent(workflow)
	if err != nil {
		return err
	}
	return c.CreateJSON(workflowName, src)
}

// AddQueue adds a workflow name to Queues attribute.
func AddQueue(workflowName, queueName string) error {
	c, err := dataset.Open("workflows.AndOr")
	if err != nil {
		return err
	}
	defer c.Close()
	src, err := c.ReadJSON(workflowName)
	if err != nil {
		return err
	}
	workflow := new(Workflow)
	if err = json.Unmarshal(src, &workflow); err != nil {
		return err
	}
	// Make sure we're not adding duplicates
	for _, key := range workflow.Queues {
		if strings.Compare(queueName, key) != 0 {
			return fmt.Errorf("already has queue of %q", queueName)
		}
	}
	workflow.Queues = append(workflow.Queues, queueName)
	return c.Update(workflowName, workflow)
}

// RemoveQueue adds a workflow name to Queues attribute.
func RemoveQueue(workflowName, queueName string) error {
	c, err := dataset.Open("workflows.AndOr")
	if err != nil {
		return err
	}
	defer c.Close()
	src, err := c.ReadJSON(workflowName)
	if err != nil {
		return err
	}
	workflow := new(Workflow)
	if err = json.Unmarshal(src, &workflow); err != nil {
		return err
	}
	queues := []string{}
	for _, key := range workflow.Queues {
		if strings.Compare(queueName, key) != 0 {
			queues = append(queues, key)
		}
	}
	workflow.Queues = queues
	return c.Update(workflowName, workflow)
}

// AddAssignTo adds a workflow to AssignTo attribute.
func AddAssignTo(workflowName, queueName string) error {
	c, err := dataset.Open("workflows.AndOr")
	if err != nil {
		return err
	}
	defer c.Close()
	src, err := c.ReadJSON(workflowName)
	if err != nil {
		return err
	}
	workflow := new(Workflow)
	if err = json.Unmarshal(src, &workflow); err != nil {
		return err
	}
	// Make sure we're not adding duplicates
	for _, key := range workflow.AssignTo {
		if strings.Compare(queueName, key) != 0 {
			return fmt.Errorf("already has assign to of %q", queueName)
		}
	}
	workflow.Queues = append(workflow.AssignTo, queueName)
	return c.Update(workflowName, workflow)
}

// RemoveAssignTo adds a workflow to AssignTo attribute.
func RemoveAssignTo(workflowName, queueName string) error {
	c, err := dataset.Open("workflows.AndOr")
	if err != nil {
		return err
	}
	defer c.Close()
	src, err := c.ReadJSON(workflowName)
	if err != nil {
		return err
	}
	workflow := new(Workflow)
	if err = json.Unmarshal(src, &workflow); err != nil {
		return err
	}
	queues := []string{}
	for _, key := range workflow.AssignTo {
		if strings.Compare(queueName, key) != 0 {
			queues = append(queues, key)
		}
	}
	workflow.AssignTo = queues
	return c.Update(workflowName, workflow)
}

// ListWorkflows returns a list of workflow objects
func ListWorkflows() ([]*Workflow, error) {
	c, err := dataset.Open("workflows.AndOr")
	if err != nil {
		return err
	}
	defer c.Close()
	keys := c.Keys()
	sort.Strings(keys)
	objects := []*Workflow{}
	for _, key := range keys {
		obj, err := c.Read(key)
		if err != nil {
			return err
		}
		objects = append(objects, obj)
	}
	return objects, nil
}

// RemoveWorkflow removes a workflow from workflows.AndOr
func RemoveWorkflow(workflowName string) error {
	c, err := dataset.Open("workflows.AndOr")
	if err != nil {
		return err
	}
	defer c.Close()
	return c.Delete(workflowName)
}

// UserInWorkflow takes user and workflow and sees if
// the user is indeed in the workflow or not.
func UserInWorkflow(user *User, workflow *Workflow) bool {
	for _, queue := range user.Workflows {
		if strings.Compare(queue, workflow.WorkflowID) == 0 {
			return true
		}
	}
	return false
}

// ObjectInWorkflow takes an object and workflow and sees if
// the object is in the workflow's queue(s)
func ObjectInWorkflow(object map[string]interface{}, workflow *Workflow) bool {
	if s, ok := object["_Queue"]; ok == true {
		switch s.(type) {
		case string:
			queueName := s.(string)
			for _, queue := range workflow.Queues {
				if strings.Compare(queueName, queue) == 0 {
					return true
				}
			}
		}
	}
	return false
}

// HasAccess takes a user, workflow, permission, and object
// it returns true if permission is affirmed false otherwise.
func HasAccess(user *User, workflow *Workflow, permission string, object map[string]*interface{}) bool {
	// Check if user is in workflow
	// Check if object is in workflow's queues
	if UserInWorkflow(user, workflow) && ObjectInWorkflow(object, workflow) {
		// Check if work flow has rights on object
		for _, objectPermission := range workflow.ObjectPermissions {
			if strings.Compare(permission, objectPermission) == 0 {
				return true
			}
		}
		return false
	}

	return false
}

// CanAssign takes a user, workflow, queue name and object
// it returns true if assignment is allowed, false otherwise
func CanAssign(user *User, workflow *Workflow, queueName string, object map[string]*interface{}) bool {
	// Check if user is in workflow
	// Check if object is in workflow's queues
	if UserInWorkflow(user, workflow) && ObjectInWorkflow(object, workflow) {
		// Check if work flow has rights to assign to new queue
		for _, assignTo := range workflow.AssignTo {
			if strings.Compare(queueName, assignTo) == 0 {
				return true
			}
		}
		return false
	}
	return false
}
