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
	"io/ioutil"
	"os"
	"path"
	"strings"

	// Toml package
	//"github.com/BurntSushi/toml"
	// forked version to CaltechLibrary so we have json dropin interfaces.
	"github.com/caltechlibrary/toml"
)

// User holds the minimal user information for AndOr.
// It DOESN'T hold any secret information, e.g. passwords.
type User struct {
	// Key holds the user id associated with a user
	Key string `json:"user_id" toml:"user_id"`
	// DisplayName holds the display name when a user is authenticated.
	DisplayName string `json:"display_name" toml:"display_name"`
	// Role holds a list of role names the user is a member of.
	Roles []string `json:"roles" toml:"roles"`
}

// HasRole takes a role name and returns true if it
// is in the list, false otherwise.
func (u *User) HasRole(roleName string) bool {
	for _, name := range u.Roles {
		if strings.Compare(name, roleName) == 0 {
			return true
		}
	}
	return false
}

// GenerateUsers generates an example users.toml file.
// suitable to edit when setting up AndOr.
func GenerateUsers(fName string) error {
	userID := os.Getenv("USER")
	if userID == "" {
		userID = "admin"
	}
	src := []byte(fmt.Sprintf(`#
#
# Example %q. Lines starting with "#" are comments.
# This shows example users.
#
[%q]
display_name = %q
member_of = [ "admin" ]
`, fName, userID, userID))
	return ioutil.WriteFile(fName, src, 0666)
}

// LoadUsers takes a file name, reads the file
// (either JSON or TOML) and returns a map[string]*User
// and an error
func LoadUsers(fName string) (map[string]*User, error) {
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
	for key, user := range users {
		if user.Key == "" {
			user.Key = key
		}
	}
	return users, nil
}

// Bytes() outputs a user to []bytes in file..
func (u *User) Bytes() []byte {
	if src, err := toml.MarshalIndent(u, "", "    "); err != nil {
		return []byte("")
	} else {
		return src
	}
}

// String() outputs a user to a string file..
func (user *User) String() string {
	if src, err := toml.MarshalIndent(user, "", "    "); err != nil {
		return ""
	} else {
		return string(src)
	}
}
