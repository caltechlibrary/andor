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
	"path"
	"strings"

	"github.com/BurntSushi/toml"
)

const (
	Version = `v0.0.0`
)

// AndOrService holds the operating parameters of the AndOr service
type AndOrService struct {
	// WorkflowsTOML
	WorkflowsTOML string `json:"workflows_toml" toml:"workflows_toml"`
	// UsersTOML
	UsersTOML string `json:"users_toml" toml:"users_toml"`
	// Protocol is usually either https: or http:
	Protocol string `json:"protocol" toml:"protocol"`
	// Port is the port to listen on, usually 8246
	Port string `json:"port" toml:"port"`
	// Post is the hostname/ip address to listen on, e.g. localhost
	Host string `json:"host" toml:"host"`

	// Users holds the user map for the service
	Users map[string]*User
	// Workflows holds the workflow map for the service
	Workflows map[string]*Workflow
	// Queues holds teh queue map for the service
	Queues map[string]*Queue
}

// GenerateAndOrTOML generates an example AndOr TOML file
// suitable to edit and then use to run AndOr.
func GenerateAndOrTOML(andorTOML, workflowsTOML, usersTOML string, collections []string) error {
	src := []byte(fmt.Sprintf(`#
# Example %q. Lines starting with "#" are comments.
# This file configuration the AndOr web service.
#
port = "8246"
host = "localhost"
protocol = "http"
workflows_toml = %q
users_toml = %q
repository = [ "%s" ]

`, andorTOML, workflowsTOML, usersTOML,
		strings.Join(collections, "\", \"")))
	return ioutil.WriteFile(andorTOML, src, 0666)
}

// LoadAndOr reads a file, parses it and returns
// an AndOrService object or error. It calls
// LoadWorkflows and LoadUsers internally.
func LoadAndOr(andorTOML string) (*AndOrService, error) {
	service := new(AndOrService)
	src, err := ioutil.ReadFile(andorTOML)
	if err != nil {
		return nil, fmt.Errorf("%q, %s", andorTOML, err)
	}
	switch path.Ext(andorTOML) {
	case ".json":
		if err := json.Unmarshal(src, &service); err != nil {
			return nil, err
		}
	case ".toml":
		if _, err := toml.Decode(string(src), &service); err != nil {
			return nil, err
		}
	default:
		return nil, fmt.Errorf("service must be either a .json or .toml file")
	}

	// Set optional values if unset.
	if service.WorkflowsTOML == "" {
		service.WorkflowsTOML = "workflows.toml"
	}
	if service.UsersTOML == "" {
		service.UsersTOML = "users.toml"
	}

	service.Workflows, service.Queues, err = LoadWorkflows(service.WorkflowsTOML)
	if err != nil {
		return nil, fmt.Errorf("%q, %s", service.WorkflowsTOML, err)
	}
	service.Users, err = LoadUsers(service.UsersTOML)
	if err != nil {
		return nil, fmt.Errorf("%q, %s", service.UsersTOML, err)
	}
	if service.Protocol == "" {
		return nil, fmt.Errorf("Missing protocol (e.g. http, https)")
	}
	if service.Port == "" {
		return nil, fmt.Errorf("Port not set")
	}
	if service.Host == "" {
		return nil, fmt.Errorf("Host not set")
	}
	// Finally check to make sure all the users MemberOf fields
	// are accounted for.
	for _, user := range service.Users {
		for _, workflow := range user.MemberOf {
			if _, ok := service.Workflows[workflow]; ok == false {
				return nil, fmt.Errorf("%s not defined, referenced by %s", workflow, user.Key)
			}
		}
	}
	return service, nil
}
