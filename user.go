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
	"strings"

	// Toml package
	"github.com/BurntSushi/toml"
)

// User holds the minimal user information for AndOr.
// It DOESN'T hold any secret information, e.g. passwords.
type User struct {
	// Key holds the user id associated with a user.
	// This is how we map into available workflows with
	// MemberOf
	Key string `json:"user_id" toml:"user_id"`
	// DisplayName holds the display name when a user is authenticated.
	DisplayName string `json:"display_name" toml:"display_name"`
	// MemberOf holds a list of workflow names the user is a member of.
	MemberOf []string `json:"member_of" toml:"member_of"`
}

// IsMemberOf takes a workflow name and returns true if it
// is in the list, false otherwise.
func (u *User) IsMemberOf(workflowName string) bool {
	for _, name := range u.MemberOf {
		if strings.Compare(name, workflowName) == 0 {
			return true
		}
	}
	return false
}

// LoadUser takes a file name, reads the file
// (either JSON or TOML) and returns a map[string]*User
// and an error
func LoadUser(fName string) (map[string]*User, error) {
	users := map[string]*User{}
	src, err := ioutil.ReadFile(fName)
	if err != nil {
		return nil, err
	}
	switch path.Ext(fName) {
	case ".json":
		if err := json.Unmarshal(src, &users); err != nil {
			return nil, err
		}
	case ".toml":
		if _, err := toml.Decode(string(src), &users); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("user must be either a .json or .toml file")
	}
	return users, nil
}

// Bytes() outputs a user to []bytes in TOML.
func (user *User) Bytes() []byte {
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(user); err != nil {
		src, _ := json.Marshal(user)
		return src
	}
	return buf.Bytes()
}

// String() outputs a user to a string TOML.
func (user *User) String() string {
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(user); err != nil {
		src, _ := json.Marshal(user)
		return string(src)
	}
	return buf.String()
}
