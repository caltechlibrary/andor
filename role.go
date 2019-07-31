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

// Role holds a single role description.
// Role defines both role queue name and
// the permissions about what can be viewed and
// what additional roles can be assigned to.
type Role struct {
	// Key holds the key to be used when referencing the role
	// E.g. "editor", "curator", "public"
	Key string `json:"role_id" toml:"role_id"`
	// Name, the display name, e.g. "Editor", "Curator", "Public View"
	Name string `json:"role_name" toml:"role_name"`
	// Queues hold name of the queue this role can operating on.
	Queue string `json:"queue" toml:"queue"`
	// Create permissions in .Queue
	Create bool `json:"create" toml:"create"`
	// Read permissions in .Queue
	Read bool `json:"read" toml:"read"`
	// Update permissions in .Queue
	Update bool `json:"update" toml:"update"`
	// Delete permissions in .Queue
	Delete bool `json:"delete" toml:"delete"`
	// AssignTo defines a list of queues that this role
	// can send objects to.
	AssignTo []string `json:"assign_to" toml:"assign_to"`
}

// GenerateRoles generates the role.toml file
// example suitable for editing when setting up AndOr.
func GenerateRoles(fName string) error {
	src := []byte(fmt.Sprintf(`#
# Example %q. Lines starting with "#" are comments.
# This file setups up the roles used by AndOr.
#
[admin]
role_name = "Administrator"
queue = "*"
create = true
read = true
update = true
delete = true
assign_to = [ "*" ]

[depositor]
role_name = "Depositor"
queue = "review"
create = true
read = false
update = false
delete = false
assign_to = [ ]

[reviewer]
role_name = "Reviewer"
queue = "review"
create = false
read = true
update = true
delete = true
assign_to = [ "published" ]

[published]
role_name = "Published"
queue = "published"
create = false
read = true
update = false
delete = false
assign_to = [ ]

`, fName))
	return ioutil.WriteFile(fName, src, 0666)
}

// makeQueues() takes a set of roles and returns a list of queues.
func makeQueues(roles map[string]*Role) map[string]*Queue {
	queues := map[string]*Queue{}
	// Create Queues from roles.Queue and roles.AssignTo
	for key, role := range roles {
		role.Key = key
		// For each queue mentioned in role, check if it
		// exists and update it with the role information.
		queueList := append([]string{role.Queue}, role.AssignTo...)
		for _, queue := range queueList {
			q, ok := queues[queue]
			if ok == false {
				q = new(Queue)
				q.Key = role.Queue
			}
			q.AddRole(role.Key)
			queues[queue] = q
		}
	}
	return queues
}

// LoadRoles reads a file (either JSON or TOML) at
// start up of AndOr web service and sets up roles and
// queues. It returns a map[string]*Role,
// a map[string]*Queue and an error
func LoadRoles(fName string) (map[string]*Role, map[string]*Queue, error) {
	roles := map[string]*Role{}

	// Parse our roles
	src, err := ioutil.ReadFile(fName)
	if err != nil {
		return nil, nil, err
	}
	switch path.Ext(fName) {
	case ".json":
		if err := json.Unmarshal(src, &roles); err != nil {
			return nil, nil, err
		}
	case ".toml":
		if _, err := toml.Decode(string(src), &roles); err != nil {
			return nil, nil, err
		}
	default:
		return nil, nil, fmt.Errorf("role must be either a .json or .toml file")
	}
	queues := makeQueues(roles)
	return roles, queues, nil
}

// Bytes() outputs a role to []bytes in TOML.
func (role *Role) Bytes() []byte {
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(role); err != nil {
		src, _ := json.MarshalIndent(role, "", "    ")
		return src
	}
	return buf.Bytes()
}

// String() outputs a role to a string TOML.
func (role *Role) String() string {
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(role); err != nil {
		src, _ := json.MarshalIndent(role, "", "    ")
		return string(src)
	}
	return buf.String()
}
