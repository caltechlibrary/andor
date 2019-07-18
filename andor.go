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

	"github.com/BurntSushi/toml"
)

const (
	Version = `v0.0.0`
)

var (
	// WorkflowsTOML holds the default name for workflows TOML file
	WorkflowsTOML = "workflows.toml"
	// UsersTOML holds the default name for users TOML file
	UsersTOML = "users.toml"
	// AndOr holds the default name for the web service configuration TOML file
	AndOrTOML = "andor.toml"
)

// AndOrService holds the operating parameters of the AndOr service
type AndOrService struct {
	// Users holds the user map for the service
	Users map[string]*User
	// Workflows holds the workflow map for the service
	Workflows map[string]*Workflow
	// Queues holds teh queue map for the service
	Queues map[string]*Queue
	// Protocol is usually either https: or http:
	Protocol string
	// Port is the port to listen on, usually 8246
	Port string
	// Post is the hostname/ip address to listen on, e.g. localhost
	Host string
}

// LoadAndOr reads a file, parses it and returns
// an AndOrService object or error. It calls
// LoadWorkflows and LoadUsers internally.
func LoadAndOr(fName string) (*AndOrService, error) {
	service := new(AndOrService)
	src, err := ioutil.ReadFile(fName)
	if err != nil {
		return nil, err
	}
	switch path.Ext(fName) {
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

	service.Workflows, service.Queues, err = LoadWorkflow(WorkflowsTOML)
	if err != nil {
		return nil, err
	}
	service.Users, err = LoadUser(UsersTOML)
	if err != nil {
		return nil, err
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
