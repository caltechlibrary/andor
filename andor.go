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
	"log"
	"path"
	"strings"

	// 3rd Party
	"github.com/BurntSushi/toml"

	// Caltech Library
	"github.com/caltechlibrary/dataset"
	"github.com/caltechlibrary/wsfn"
)

const (
	Version = `v0.0.0`
)

// AndOrService holds the operating parameters of the AndOr service
type AndOrService struct {
	// CertPEM if running under TLS this is the cert file
	CertPEM string `json:"cert_pem" toml:"cert_pem"`
	// KeyPEM if running under TLS this is the key file
	KeyPEM string `json:"key_pem": toml:"key_pem"`
	// WorkflowsFile
	WorkflowsFile string `json:"workflows_file" toml:"workflows_file"`
	// UsersFile
	UsersFile string `json:"users_file" toml:"users_file"`
	// Scheme is usually either "https" or "http"
	Scheme string `json:"scheme" toml:"scheme"`
	// Port is the port to listen on, usually 8246
	Port string `json:"port" toml:"port"`
	// Post is the hostname/ip address to listen on, e.g. localhost
	Host string `json:"host" toml:"host"`
	// Htdocs is the static page directory if desired (e.g. could
	// host web forms or JavaScript libraries.
	Htdocs string `json:"htdocs" toml:"htdocs"`

	// CORS policy
	CORS *wsfn.CORSPolicy `json:"cors" toml:"cors"`

	// CollectionNames holds one or more dataset
	// collection names used for web API
	CollectionNames []string `json:"collections" toml:"collections"`

	// Users holds the user map for the service
	Users map[string]*User
	// Workflows holds the workflow map for the service
	Workflows map[string]*Workflow
	// Queues holds teh queue map for the service
	Queues map[string]*Queue
	// Collections holds a map of collection name to dataset collection pointer
	Collections map[string]*dataset.Collection

	// Access is based on wsfn.Access model. It supports Basic Auth
	// and could be extended to support Shibboleth, OAuth 2 or
	// JSON web tokens. You use the wsfn tools to maintain this file.
	AccessFile string `json:"access_file,omitempty" toml:"access_file,omitempty"`

	// Access holds an *wsfn.Access object
	Access *wsfn.Access
}

// GenerateAndOr generates an example AndOr TOML or JSON file
// suitable to edit and then use to run AndOr.
func GenerateAndOr(fName string, collections []string) error {
	s := new(AndOrService)
	s.Port = "8246"
	s.Host = "localhost"
	s.Scheme = "http"
	if len(collections) > 0 {
		s.CollectionNames = collections
	}
	buf := new(bytes.Buffer)
	if err := toml.NewEncoder(buf).Encode(s); err != nil {
		return err
	}
	//  Now write out our default config.
	var (
		src []byte
		err error
	)

	src = []byte(fmt.Sprintf(`#
# Example %q 
#
# Lines starting with "#" are comments.
# This file configuration the AndOr web service.
#
%s

#
# You should uncomment these after editing them appropriately
#

# htdocs holds your web document root, e.g. /var/www/htdocs
#htdocs = "htdocs"

# collections holds a list of dataset collections
#collections = ["repository.ds"]

# workflows holds the workflows and queue definitions for this And/Or
#workflows = "workflows.toml"

# users holds user workflow assignments
#users = "users.toml"

# If using basic auth create this file with "webaccess" tool from
# https://github.com/caltechlibrary/wsfn project.
#access = "access.toml"

# If running under TLS you need to include cert_pem and key_pem
# fileds that point at your cert.pem and key.pem certificate files.
#protocol = "https"
#cert_pem = "/etc/ssl/certs/cert.pem"
#key_pem = "/etc/ssl/certs/key.pem"
`, fName, buf.String()))
	switch {
	case strings.HasSuffix(fName, ".json"):
		o := new(AndOrService)
		if _, err = toml.Decode(string(src), &o); err != nil {
			return err
		}
		if src, err = json.MarshalIndent(o, "", "    "); err != nil {
			return err
		}
	}
	return ioutil.WriteFile(fName, src, 0666)
}

// LoadAndOr reads a file, parses it and returns
// an AndOrService object or error. It calls
// LoadWorkflows and LoadUsers internally.
func LoadAndOr(fName string) (*AndOrService, error) {
	service := new(AndOrService)
	src, err := ioutil.ReadFile(fName)
	if err != nil {
		return nil, fmt.Errorf("%q, %s", fName, err)
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

	// Check required values
	if service.WorkflowsFile == "" {
		return nil, fmt.Errorf("%q, missing workflows file", fName)
	}
	if service.UsersFile == "" {
		return nil, fmt.Errorf("%q, missing users file", fName)
	}
	if len(service.CollectionNames) == 0 {
		return nil, fmt.Errorf("%q, missing collection name(s)", fName)
	}
	for i, repo := range service.CollectionNames {
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

	if service.Scheme == "" {
		return nil, fmt.Errorf("Missing scheme (e.g. http, https)")
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
	s.Workflows, s.Queues, err = LoadWorkflows(s.WorkflowsFile)
	if err != nil {
		return fmt.Errorf("%q, %s", s.WorkflowsFile, err)
	}
	s.Users, err = LoadUsers(s.UsersFile)
	if err != nil {
		return fmt.Errorf("%q, %s", s.UsersFile, err)
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

// Start starts the webservice based on the current configuration
// of the service struct.
func (s *AndOrService) Start() error {
	var err error
	// Load an workflows.toml
	s.Workflows, s.Queues, err = LoadWorkflows(s.WorkflowsFile)
	if err != nil {
		log.Printf("Failed to load %q, %s", s.WorkflowsFile, err)
		return err
	}
	// Load an users.toml
	s.Users, err = LoadUsers(s.UsersFile)
	if err != nil {
		log.Printf("Failed to load %q, %s", s.UsersFile, err)
		return err
	}

	// Load an access.toml
	if s.AccessFile != "" {
		s.Access, err = wsfn.LoadAccess(s.AccessFile)
		if err != nil {
			return err
		}
	}

	// Validate AndOrService configuration
	if len(s.Users) == 0 {
		return fmt.Errorf("No users configured, no one can access service")
	}
	if len(s.Workflows) == 0 {
		return fmt.Errorf("No workflows defined, no one can access service")
	}
	if len(s.Queues) == 0 {
		return fmt.Errorf("No object queues defined, nothing to access")
	}
	if len(s.CollectionNames) == 0 {
		return fmt.Errorf("No collections defined, nothing to access")
	}
	// Open any repositories
	log.Printf("Loading %d collection(s), %s", len(s.CollectionNames), strings.Join(s.CollectionNames, ", "))
	s.Collections = map[string]*dataset.Collection{}
	for i, cName := range s.CollectionNames {
		// Add paths for open collections
		log.Printf("Opening (%d) %s", i, cName)
		c, err := dataset.Open(cName)
		if err != nil {
			log.Printf("Can't open (%d) %s, %s", i, cName, err)
			return err
		}
		defer c.Close()
		// Save the map for use by RunService()
		// so we need to trim any .ds suffixed.
		cName = strings.TrimSuffix(cName, ".ds")
		s.Collections[cName] = c
	}

	// Start http(s) service with AndOr end points
	if err = RunService(s); err != nil {
		return err
	}

	// FIXME: SOMEDAY, MAYBE, Watch for signals
	// See https://github.com/golang/go/wiki/SignalHandling
	// See https://golang.org/pkg/os/signal/
	// See https://gist.github.com/reiki4040/be3705f307d3cd136e85
	return nil
}
