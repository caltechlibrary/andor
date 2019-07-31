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
// Role defines both role state name and
// the permissions about what can be viewed and
// what additional roles can be assigned to.
type Role struct {
	// Key holds the key to be used when referencing the role
	// E.g. "editor", "curator", "public"
	Key string `json:"role_id" toml:"role_id"`
	// Name, the display name, e.g. "Editor", "Curator", "Public View"
	Name string `json:"role_name" toml:"role_name"`
	// States holds the state names this role can operating on.
	States []string `json:"states,omitempty" toml:"states,omitempty"`
	// Create permissions for states
	Create bool `json:"create" toml:"create"`
	// Read permissions for states
	Read bool `json:"read" toml:"read"`
	// Update permissions for states
	Update bool `json:"update" toml:"update"`
	// Delete permissions for states
	Delete bool `json:"delete" toml:"delete"`
	// AssignTo defines a list of states that this role
	// can send objects to.
	AssignTo []string `json:"assign_to,omitempty" toml:"assign_to,omitempty"`
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
states = [ "*" ]
create = true
read = true
update = true
delete = true
assign_to = [ "*" ]

[depositor]
role_name = "Depositor"
states = [ "review" ]
create = true
read = false
update = false
delete = false
assign_to = [ ]

[reviewer]
role_name = "Reviewer"
states = [ "review", "embargoed" ]
create = false
read = true
update = true
delete = true
assign_to = [ "review", "embargoed", "published" ]

[curator]
role_name = "curator"
states = [ "review", "embargoed", "published" ]
create = true
read = true
update = true
delete = true
assign_to = [ "*" ]

`, fName))
	return ioutil.WriteFile(fName, src, 0666)
}

// makeStates() takes a set of roles and returns a list of states.
func makeStates(roles map[string]*Role) map[string]*State {
	states := map[string]*State{}
	// Create States from roles.State and roles.AssignTo
	for key, role := range roles {
		role.Key = key
		// For each state mentioned in role, check if it
		// exists and update it with the role information.
		stateList := append(role.States, role.AssignTo...)
		for _, state := range stateList {
			q, ok := states[state]
			if ok == false {
				q = new(State)
				q.Key = state
			}
			q.AddRole(role.Key)
			states[state] = q
		}
	}
	return states
}

// LoadRoles reads a file (either JSON or TOML) at
// start up of AndOr web service and sets up roles and
// states. It returns a map[string]*Role,
// a map[string]*State and an error
func LoadRoles(fName string) (map[string]*Role, map[string]*State, error) {
	roles := make(map[string]*Role)

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
	states := makeStates(roles)
	return roles, states, nil
}

// Bytes() outputs a role to []bytes in TOML.
func (role *Role) Bytes() []byte {
	if role == nil {
		return []byte{}
	}
	buf := new(bytes.Buffer)
	encoder := toml.NewEncoder(buf)
	if err := encoder.Encode(role); err != nil {
		fmt.Printf("DEBUG encoder.Encode(role) error in Bytes() -> %s\n", err)
		return []byte{}
	}
	return buf.Bytes()
}

// String() outputs a role to a string TOML.
func (role *Role) String() string {
	if role == nil {
		return ""
	}
	src := role.Bytes()
	return string(src)
}
