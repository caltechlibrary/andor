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

// User holds the minimal user information for AndOr.
// It DOESN'T hold any secret information, e.g. passwords.
type User struct {
	// Key holds the user id associated with a user.
	// This is how we map into available workflows with
	// MemberOf
	Key string `json:"user_id"`
	// DisplayName holds the display name when a user is authenticated.
	DisplayName string `json:"display_name"`
	// CreateObjectIn holds the default queue name used
	// when creating objects.
	CreateObjectsIn string `json:"create_object_as"`
	// Workflows holds a list of workflow names the user is a member of.
	Workflows string `json:"workflows"`
}

// ReadUserFile takes a filename, reads the file
// (either JSON or TOML) and returns a user
// object and error.
func ReadUserFile(fName string) (*User, error) {
	user := new(User)
	src, err := ioutil.ReadFile(fName)
	if err != nil {
		return nil, err
	}
	switch path.Exit(fName) {
	case ".json":
		if err := json.Unmarshal(src, &user); err != nil {
			return user, err
		}
	case ".toml":
		if _, err := toml.Decode(src, &user); err != nil {
			return user, err
		}
	default:
		return nil, fmt.Errorf("user must be either a .json or .toml file")
	}
	return user, nil
}

// Bytes() outputs a user to []bytes in TOML.
func (user *User) Bytes() []byte {
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(user); err != nil {
		return fmt.Sprintf("%+v", user)
	}
	return buf
}

// String() outputs a user to a string TOML.
func (user *User) String() string {
	return user.Bytes().String()
}

// AddUser adds a user to the "users.AndOr"
// dataset collection.
func AddUser(userName string, user *User) error {
	c, err := dataset.Open("users.AndOr")
	if err != nil {
		return err
	}
	defer c.Close()
	src, err := json.MarshalIndent(user)
	if err != nil {
		return err
	}
	return c.CreateJSON(userName, src)
}

// AddMemberOf adds a workflow to a user object
func AddMemberOf(userName, workflowName string) error {
	c, err := dataset.Open("users.AndOr")
	if err != nil {
		return err
	}
	defer c.Close()
	src, err := c.ReadJSON(userName)
	if err != nil {
		return err
	}
	user := new(User)
	if err = json.Unmarshal(src, &user); err != nil {
		return err
	}
	// Make sure we're not adding duplicates
	for _, key := range user.MemberOf {
		if strings.Compare(workflowName, key) != 0 {
			return fmt.Errorf("already a member of %q", workflowName)
		}
	}
	user.MemberOf = append(user.MemberOf, workflowName)
	return c.Update(userName, user)
}

// RemoveMemberOf removes a workflow for a user object
func RemoveMemberOf(userName, workflowName string) error {
	c, err := dataset.Open("users.AndOr")
	if err != nil {
		return err
	}
	defer c.Close()
	src, err := c.ReadJSON(userName)
	if err != nil {
		return err
	}
	user := new(User)
	if err = json.Unmarshal(src, &user); err != nil {
		return err
	}
	// Make remove any occurrences of the same workgroup name.
	memberOf := []string{}
	for _, key := range user.MemberOf {
		if strings.Compare(workflowName, key) != 0 {
			memberOf = append(user.MemberOf, key)
		}
	}
	user.MemberOf = memberOf
	return c.Update(userName, user)
}

// AddUserAssignTo adds a workflow assignment to a user object
func AddUserAssignTo(userName, workflowName string) error {
	c, err := dataset.Open("users.AndOr")
	if err != nil {
		return err
	}
	defer c.Close()
	src, err := c.ReadJSON(userName)
	if err != nil {
		return err
	}
	user := new(User)
	if err = json.Unmarshal(src, &user); err != nil {
		return err
	}
	// Make sure we're not adding duplicates
	for _, key := range user.AssignTo {
		if strings.Compare(workflowName, key) != 0 {
			return fmt.Errorf("already a member of %q", workflowName)
		}
	}
	user.MemberOf = append(user.AssignTo, workflowName)
	return c.Update(userName, user)
}

// RemoveUserAssignTo removes a workflow assignment for a user object
func RemoveAssignTo(userName, workflowName string) error {
	c, err := dataset.Open("users.AndOr")
	if err != nil {
		return err
	}
	defer c.Close()
	src, err := c.ReadJSON(userName)
	if err != nil {
		return err
	}
	user := new(User)
	if err = json.Unmarshal(src, &user); err != nil {
		return err
	}
	// Make remove any occurrences of the same workgroup name.
	assignTo := []string{}
	for _, key := range user.AssignTo {
		if strings.Compare(workflowName, key) != 0 {
			memberOf = append(user.AssignTo, key)
		}
	}
	user.AssignTo = assignTo
	return c.Update(userName, user)
}

// ListUsers returns a list of user objects
func ListUsers() ([]*User, error) {
	c, err := dataset.Open("users.AndOr")
	if err != nil {
		return err
	}
	defer c.Close()
	keys := c.Keys()
	sort.Strings(keys)
	objects := []*User{}
	for _, key := range keys {
		obj, err := c.Read(key)
		if err != nil {
			return err
		}
		objects = append(objects, obj)
	}
	return objects, nil
}

// RemoveUser removes a user from users.AndOr
func RemoveUser(userName string) error {
	c, err := dataset.Open("users.AndOr")
	if err != nil {
		return err
	}
	defer c.Close()
	return c.Delete(userName)
}
