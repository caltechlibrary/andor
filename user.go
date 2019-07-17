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
	"strings"

	// Caltech Library packages
	"github.com/caltechlibrary/dataset"

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
	// CreateQueue holds the default queue name used
	// when creating objects.
	CreateQueue string `json:"create_queue" toml:"create_queue"`
	// MemberOf holds a list of workflow names the user is a member of.
	MemberOf []string `json:"member_of" toml:"member_of"`
}

// LoadUser takes one or more filenames, reads the file
// (either JSON or TOML) and loads the users into AndOr
// NOTE: it add/updates records and does NOT remove records.
func LoadUser(fNames []string) error {
	c, err := dataset.Open(AndOrUsers)
	if err != nil {
		return err
	}
	defer c.Close()

	for _, fName := range fNames {
		users := map[string]*User{}
		src, err := ioutil.ReadFile(fName)
		if err != nil {
			return err
		}
		fmt.Printf("DEBUG src -> %s\n", src)
		switch path.Ext(fName) {
		case ".json":
			if err := json.Unmarshal(src, &users); err != nil {
				return err
			}
		case ".toml":
			if _, err := toml.Decode(string(src), &users); err != nil {
				return err
			}
		default:
			return fmt.Errorf("user must be either a .json or .toml file")
		}
		fmt.Printf("DEBUG users => %+v\n", users)
		for key, user := range users {
			user.Key = key
			fmt.Printf("DEBUG user -> %s\n", user.Bytes())
			src, err := json.MarshalIndent(user, "", "    ")
			if err != nil {
				return err
			}
			if c.HasKey(key) == true {
				if err := c.UpdateJSON(key, src); err != nil {
					return err
				}
			} else {
				if err := c.CreateJSON(key, src); err != nil {
					return err
				}
			}
		}
	}
	return nil
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

// AddUser adds a user to the "users.AndOr"
// dataset collection.
func AddUser(userName string, user *User) error {
	c, err := dataset.Open(AndOrUsers)
	if err != nil {
		return err
	}
	defer c.Close()
	src, err := json.MarshalIndent(user, "", "    ")
	if err != nil {
		return err
	}
	return c.CreateJSON(userName, src)
}

// AddMemberOf adds a workflow to a user object
func AddMemberOf(userName, workflowName string) error {
	c, err := dataset.Open(AndOrUsers)
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
	src, err = json.MarshalIndent(user, "", "    ")
	if err != nil {
		return err
	}
	return c.UpdateJSON(userName, src)
}

// ListUsers returns a list of user objects
func ListUsers() ([]*User, error) {
	c, err := dataset.Open(AndOrUsers)
	if err != nil {
		return nil, err
	}
	defer c.Close()
	keys := c.Keys()
	sort.Strings(keys)
	objects := []*User{}
	for _, key := range keys {
		src, err := c.ReadJSON(key)
		if err != nil {
			return nil, err
		}
		obj := new(User)
		err = json.Unmarshal(src, &obj)
		if err != nil {
			return nil, err
		}
		objects = append(objects, obj)
	}
	return objects, nil
}

// RemoveUser removes user(s) from "users.AndOr"
func RemoveUser(userNames []string) error {
	c, err := dataset.Open(AndOrUsers)
	if err != nil {
		return err
	}
	defer c.Close()
	for _, userName := range userNames {
		if userName == "*" {
			keys := c.Keys()
			return RemoveUser(keys)
		}
		if err := c.Delete(userName); err != nil {
			return err
		}
	}
	return nil
}
