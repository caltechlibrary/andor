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
	"log"
	"net/http"
	"net/url"
	"path"
	"sort"
	"strings"

	// Caltech Library Packages
	"github.com/caltechlibrary/dataset"
	"github.com/caltechlibrary/wsfn"
)

var webService *wsfn.WebService

func writeResponse(w http.ResponseWriter, r *http.Request, src []byte) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(src); err != nil {
		log.Printf("Response write error, %s %s", r.URL.Path, err)
		return
	}
	log.Printf("FIXME: Log successful requests here ... %s", r.URL.Path)
}

func (s *AndOrService) requestAccessInfo(w http.ResponseWriter, r *http.Request) {
	//FIXME: This should really be JSON Web Token based ...
	username, _, ok := r.BasicAuth()
	if ok == false || username == "" {
		username = "anonymous"
	}

	log.Printf("DEBUG username %q", username)
	// Are we logged in?
	if u, ok := s.Users[username]; ok == true {
		roleMap := make(map[string]*Role)
		// Is user member of role?
		for _, key := range u.MemberOf {
			if role, ok := s.Roles[key]; ok == true {
				roleMap[key] = role
			}
		}
		src, err := json.MarshalIndent(map[string]interface{}{
			"user":  u,
			"roles": roleMap,
		}, "", "    ")
		if err != nil {
			log.Printf("Failed to marshal %q, %s", username, err)
			http.Error(w, "Internal Server error", http.StatusInternalServerError)
		}
		// return payload appropriately
		writeResponse(w, r, src)
		return
	}
	// Otherwise return 404, Not Found
	http.NotFound(w, r)
}

// requestKeys is the API version of `dataset keys COLLECTION_NAME`
// We only support GET on keys.
func requestKeys(cName string, c *dataset.Collection, w http.ResponseWriter, r *http.Request) {
	//FIXME: Need to apply users/roles/states rules.
	keys := c.Keys()
	sort.Strings(keys)
	src, err := json.MarshalIndent(keys, "", "    ")
	if err != nil {
		log.Printf("Internal Server error, %s %s", cName, err)
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
		return
	}
	writeResponse(w, r, src)
}

// requestCreate is the API version of
//	`dataset create COLLECTION_NAME OBJECT_ID OBJECT_JSON`
func requestCreate(cName string, c *dataset.Collection, w http.ResponseWriter, r *http.Request) {
	//FIXME: Need to apply users/roles/states rules.
	//FIXME: Need to make sure this part of the service is behind
	// muxtex.
	log.Printf("requestCreate(%q, ...) not implemented", cName)
	http.Error(w, "Internal Server error", http.StatusInternalServerError)
}

// requestRead is the API version of
//     `dataset read -c -p COLLECTION_NAME KEY`
func requestRead(cName string, c *dataset.Collection, w http.ResponseWriter, r *http.Request) {
	//FIXME: Need to apply user/role/queue rules.
	key := strings.TrimPrefix(r.URL.Path, "/"+cName+"/read/")
	if c.HasKey(key) == false {
		log.Printf("%s, %s, unknown key", cName, r.URL.Path)
		http.NotFound(w, r)
		return
	}
	src, err := c.ReadJSON(key)
	if err != nil {
		log.Printf("Error reading key %q from %q, %s", key, cName, err)
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
		return
	}
	writeResponse(w, r, src)
}

// requestUpdate is the API version of
//	`dataset update COLLECTION_NAME OBJECT_ID OBJECT_JSON`
func requestUpdate(cName string, c *dataset.Collection, w http.ResponseWriter, r *http.Request) {
	//FIXME: Need to apply users/roles/states rules.
	//FIXME: Need to make sure this part of the service is behind
	// muxtex.
	log.Printf("requestUpdate(%q, ...) not implemented", cName)
	http.Error(w, "Internal Server error", http.StatusInternalServerError)
}

// requestDelete is the API version of
//	`dataset Delete COLLECTION_NAME OBJECT_ID`
// except is doesn't actually delete the object. It changes the
// object's `._State` value.
func requestDelete(cName string, c *dataset.Collection, w http.ResponseWriter, r *http.Request) {
	//FIXME: Need to apply users/roles/states rules.
	//FIXME: Need to make sure this part of the service is behind
	// muxtex.
	log.Printf("requestDelete(%q, ...) not implemented", cName)
	http.Error(w, "Internal Server error", http.StatusInternalServerError)
}

// requestAttach is the API version of
//	`dataset attach COLLECTION_NAME OBJECT_ID FILENAMES`
func requestAttach(cName string, c *dataset.Collection, w http.ResponseWriter, r *http.Request) {
	//FIXME: Need to apply users/roles/states rules.
	//FIXME: Need to make sure this part of the service is behind
	// muxtex.
	log.Printf("requestAttach(%q, ...) not implemented", cName)
	http.Error(w, "Internal Server error", http.StatusInternalServerError)
}

// requestAttachments is the API version of
//	`dataset attachments COLLECTION_NAME OBJECT_ID`
func requestAttachments(cName string, c *dataset.Collection, w http.ResponseWriter, r *http.Request) {
	//FIXME: Need to apply users/roles/states rules.
	//FIXME: Need to make sure this part of the service is behind
	// muxtex.
	log.Printf("requestAttachments(%q, ...) not implemented", cName)
	http.Error(w, "Internal Server error", http.StatusInternalServerError)
}

// requestDetach is the API version of
//	`dataset detach COLLECTION_NAME OBJECT_ID FILENAME`
func requestDetach(cName string, c *dataset.Collection, w http.ResponseWriter, r *http.Request) {
	//FIXME: Need to apply users/roles/states rules.
	log.Printf("requestDetach(%q, ...) not implemented", cName)
	http.Error(w, "Internal Server error", http.StatusInternalServerError)
}

// requestPrune is the API version of
//	`dataset prune COLLECTION_NAME OBJECT_ID`
func requestPrune(cName string, c *dataset.Collection, w http.ResponseWriter, r *http.Request) {
	//FIXME: Need to apply users/roles/states rules.
	//FIXME: Need to make sure this part of the service is behind
	// muxtex.
	log.Printf("requestPrune(%q, ...) not implemented", cName)
	http.Error(w, "Internal Server error", http.StatusInternalServerError)
}

func addAccessRoute(a *wsfn.Access, p string) {
	log.Printf("DEBUG a -> %+v", a)
	if a != nil {
		if a.Routes == nil {
			a.Routes = []string{}
		}
		a.Routes = append(a.Routes, p)
	}
}

// assignHandlers generates the /keys, /create, /read, /delete
// end points for accessing a collection in And/Or.
func (s *AndOrService) assignHandlers(mux *http.ServeMux, c *dataset.Collection) {
	cName := c.Name
	access := s.Access
	//NOTE: We create a function handler based on on the
	// current collection being processed.
	log.Printf("Adding collection %q", cName)
	base := "/" + path.Base(cName)
	log.Printf("Adding access route %q", base)
	if s.IsAccessRestricted() {
		addAccessRoute(access, base)
	}
	// End points based on dataset
	p := base + "/keys"
	log.Printf("Adding handler %s", p)
	mux.HandleFunc(p, func(w http.ResponseWriter, r *http.Request) {
		requestKeys(cName, c, w, r)
	})
	// dataset object handling
	p = base + "/create"
	mux.HandleFunc(p, func(w http.ResponseWriter, r *http.Request) {
		requestCreate(cName, c, w, r)
	})
	p = base + "/read"
	mux.HandleFunc(p, func(w http.ResponseWriter, r *http.Request) {
		requestRead(cName, c, w, r)
	})
	p = base + "/update"
	mux.HandleFunc(p, func(w http.ResponseWriter, r *http.Request) {
		requestUpdate(cName, c, w, r)
	})
	p = base + "/delete"
	mux.HandleFunc(p, func(w http.ResponseWriter, r *http.Request) {
		requestDelete(cName, c, w, r)
	})
	// dataset attachment handling
	p = base + "/attach"
	mux.HandleFunc(p, func(w http.ResponseWriter, r *http.Request) {
		requestAttach(cName, c, w, r)
	})
	p = base + "/attachments"
	mux.HandleFunc(p, func(w http.ResponseWriter, r *http.Request) {
		requestAttachments(cName, c, w, r)
	})
	p = base + "/detach"
	mux.HandleFunc(p, func(w http.ResponseWriter, r *http.Request) {
		requestDetach(cName, c, w, r)
	})
	p = base + "/prune"
	mux.HandleFunc(p, func(w http.ResponseWriter, r *http.Request) {
		requestPrune(cName, c, w, r)
	})

	// Additional And/Or specific end points
	p = "/" + path.Base(cName) + "/access/"
	log.Printf("Adding handler %s", p)
	mux.HandleFunc(p, s.requestAccessInfo)
}

// RunService runs the http/https web service of AndOr.
func RunService(s *AndOrService) error {
	var (
		access *wsfn.Access
		cors   *wsfn.CORSPolicy
	)
	// Setup our web service from our *AndOrService
	u := new(url.URL)
	u.Scheme = s.Scheme
	u.Host = s.Host + ":" + s.Port
	if s.Access != nil {
		access = s.Access
	}
	if s.CORS != nil {
		cors = s.CORS
	}
	mux := http.NewServeMux()

	log.Printf("Have %d collection(s)", len(s.Collections))
	// NOTE: For each collection we assign our set of end points.
	for _, c := range s.Collections {
		s.assignHandlers(mux, c)
	}
	if s.Htdocs != "" {
		fs, err := wsfn.MakeSafeFileSystem(s.Htdocs)
		if err != nil {
			return err
		}
		mux.Handle("/", http.FileServer(fs))
	}
	hostname := fmt.Sprintf("%s:%s", u.Hostname(), u.Port())
	log.Printf("Starting service %s", hostname)
	switch s.Scheme {
	case "http":
		return http.ListenAndServe(hostname, wsfn.RequestLogger(cors.Handler(access.Handler(mux))))
	case "https":
		return http.ListenAndServeTLS(hostname, s.CertPEM, s.KeyPEM, wsfn.RequestLogger(cors.Handler(access.Handler(mux))))
	default:
		return fmt.Errorf("%q url scheme not supported", s.Scheme)
	}
}
