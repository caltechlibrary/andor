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
	// Repository holds one or more repository names used for web API
	Repositories []string `json:"repositories" toml:"repositories"`

	// Users holds the user map for the service
	Users map[string]*User
	// Workflows holds the workflow map for the service
	Workflows map[string]*Workflow
	// Queues holds teh queue map for the service
	Queues map[string]*Queue
}

// GenerateAndOrTOML generates an example AndOr TOML file
// suitable to edit and then use to run AndOr.
func GenerateAndOrTOML(andorTOML string, collections []string) error {
	s := new(AndOrService)
	s.Port = "8246"
	s.Host = "localhost"
	s.Protocol = "http"
	s.WorkflowsTOML = "workflows.toml"
	s.UsersTOML = "users.toml"
	if len(collections) > 0 {
		s.Repositories = collections
	}
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(s); err != nil {
		return err
	}
	//  Now write out our default config.
	src := []byte(fmt.Sprintf(`#
# Example %q 
#
# Lines starting with "#" are comments.
# This file configuration the AndOr web service.
#
%s
`, andorTOML, buf.String()))
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
	for i, repo := range service.Repositories {
		if len(repo) == 0 {
			suffix := ""
			switch i + 1 {
			case 1:
				suffix = "st"
			case 2:
				suffix = "nd"
			case 3:
				suffix = "nd"
			default:
				suffix = "th"
				return nil, fmt.Errorf("%d%s repository is an empty string", i, suffix)
			}
		}
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
	return service, nil
}

// LoadWorksAndUsers envokes LoadWorkflows() and LoadUsers() for
// a service instance. This is separate from LoadAndOr() because you
// may want to support HUP to reload workflows and users as well as
// support independent validation and testing.
func (s *AndOrService) LoadWorkersAndUsers() error {
	var (
		err error
	)
	s.Workflows, s.Queues, err = LoadWorkflows(s.WorkflowsTOML)
	if err != nil {
		return fmt.Errorf("%q, %s", s.WorkflowsTOML, err)
	}
	s.Users, err = LoadUsers(s.UsersTOML)
	if err != nil {
		return fmt.Errorf("%q, %s", s.UsersTOML, err)
	}

	// Finally check to make sure all the users MemberOf fields
	// are accounted for.
	for _, user := range s.Users {
		for _, workflow := range user.MemberOf {
			if _, ok := s.Workflows[workflow]; ok == false {
				return fmt.Errorf("%s not defined, referenced by %s", workflow, user.Key)
			}
		}
	}
	return nil
}
