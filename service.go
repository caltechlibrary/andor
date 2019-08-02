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
	"sync"

	// Caltech Library Packages
	"github.com/caltechlibrary/dataset"
	"github.com/caltechlibrary/wsfn"
)

var (
	webService *wsfn.WebService
	mutex      = new(sync.Mutex)
)

// safeDatasetOp wraps Create, Update, Delete in a mutex
// to prevent corruption of items on disc like collection.json
func safeDatasetOp(c *dataset.Collection, key string, object map[string]interface{}, op int) error {
	mutex.Lock()
	defer mutex.Unlock()
	switch op {
	case CREATE:
		return c.Create(key, object)
	case UPDATE:
		return c.Update(key, object)
	case DELETE:
		return c.Delete(key)
	default:
		return fmt.Errorf("Unsupported operation %d", op)
	}
}

// writeError
func writeError(w http.ResponseWriter, statusCode int) {
	http.Error(w, http.StatusText(statusCode), statusCode)
}

// writeJSON
func writeJSON(w http.ResponseWriter, r *http.Request, src []byte) {
	w.Header().Add("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	if _, err := w.Write(src); err != nil {
		log.Printf("Response write error, %s %s", r.URL.Path, err)
		return
	}
	log.Printf("FIXME: Log successful requests here ... %s", r.URL.Path)
}

func (s *AndOrService) requestAccessInfo(w http.ResponseWriter, r *http.Request) {
	// Who am I?
	username := s.getUsername(r)
	// What roles do I have?
	if roles, ok := s.getUserRoles(username); ok == true {
		src, err := json.MarshalIndent(roles, "", "    ")
		if err != nil {
			log.Printf("Failed to marshal %q, %s", username, err)
			writeError(w, http.StatusInternalServerError)
			return
		}
		// return payload appropriately
		writeJSON(w, r, src)
		return
	}
	// Otherwise return 404, Not Found
	http.NotFound(w, r)
}

// requestKeys is the API version of `dataset keys COLLECTION_NAME`
// We only support GET on keys.
func (s *AndOrService) requestKeys(cName string, c *dataset.Collection, w http.ResponseWriter, r *http.Request) {
	//FIXME: Need to apply users/roles/states rules.
	keys := c.Keys()
	sort.Strings(keys)
	src, err := json.MarshalIndent(keys, "", "    ")
	if err != nil {
		log.Printf("Internal Server error, %s %s", cName, err)
		writeError(w, http.StatusInternalServerError)
		return
	}
	writeJSON(w, r, src)
}

// requestCreate is the API version of
//	`dataset create COLLECTION_NAME OBJECT_ID OBJECT_JSON`
func (s *AndOrService) requestCreate(cName string, c *dataset.Collection, w http.ResponseWriter, r *http.Request) {
	var (
		src []byte
		err error
	)
	// Make sure we have the right http Method
	if r.Method != "POST" {
		writerError(w, http.StatusMethodNotAllowed)
		return
	}

	// Make sure we can determine permissions before reading
	// post data.
	username := s.getUsername(r)
	if username == "" {
		writeError(w, http.StatusUnauthorized)
		return
	}
	roles, ok := s.getUserRoles(username)
	if ok == false {
		writeError(w, http.StatusUnauthorized)
		return
	}
	// We need to get the submitted object before checking
	// isAllowed.
	src, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("%s %s", r.URL.Path, err)
		writeError(w, http.StatusNotAcceptable)
		return
	}
	// We only accept content in JSON form with /create.
	object := make(map[string]interface{})
	decoder := json.NewDecoder(r.Body)
	decoder.UseNumber()
	err := decoder.Unmarshal(src, &object)
	if err != nil {
		log.Printf("%s %s", r.URL.Path, err)
		writeError(w, http.StatusUnsupportedMediaType)
		return
	}
	// Need to apply users/roles/states rules.
	state := getState(object)
	if s.isAllowed(roles, state, CREATE) == false {
		writeError(w, http.StatusUnauthorized)
		return
	}

	// Now get the proposed key.
	key := getKey(r.URL.Path, "/"+cName+"/create/")

	// Need to make sure this part of the service is behind
	// the mutex.
	if err := safeDatasetOp(c, key, object, CREATE); err != nil {
		log.Printf("%s %s", r.URL.Path, err)
		writeError(w, http.StatusNotAcceptable)
		return
	}
	log.Printf("%s created %s in %s", username, key, c.Name)
}

// requestRead is the API version of
//     `dataset read -c -p COLLECTION_NAME KEY`
func (s *AndOrService) requestRead(cName string, c *dataset.Collection, w http.ResponseWriter, r *http.Request) {
	var (
		src []byte
		err error
	)
	username := s.getUsername(r)
	if username == "" {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	roles, ok := s.getUserRoles(username)
	if ok == false {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	//FIXME: need to apply state filtering to keys requested
	keys := getKeys(r.URL.Path, "/"+cName+"/read/")
	if len(keys) == 0 {
		writeError(w, http.StatusBadRequest)
		return
	}
	unauthorized := false
	objects := []map[string]interface{}{}
	for _, key = range keys {
		key = strings.TrimSpace(key)
		if key != "" {
			object := make(map[string]interface{})
			if err = c.Read(strings.TrimSpace(key), object, false); err != nil {
				//FIXME: what do we do if one of a list of keys not found?
				log.Printf("Error reading key %q from %q, %s", key, c.Name, err)
			} else {
				state := getState(object)
				if s.isAllowed(roles, state, READ) {
					objects = append(objects, object)
				} else {
					unauthorized = true
					log.Printf("%q not allowed to read %q from %q", username, key, c.Name)
				}
			}
		}
	}
	switch len(objects) {
	case 0:
		if unauthorized {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		http.Error(w, "Not found", http.StatusNotFound)
		return
	case 1:
		src, err := json.Marshal(objects[0])
	default:
		src, err = json.Marshal(objects)
		if err != nil {
			log.Printf("Error reading key %q from %q, %s", key, cName, err)
			http.Error(w, "Internal Server error", http.StatusInternalServerError)
			return
		}
	}
	writeJSON(w, r, src)
}

// requestUpdate is the API version of
//	`dataset update COLLECTION_NAME OBJECT_ID OBJECT_JSON`
func (s *AndOrService) requestUpdate(cName string, c *dataset.Collection, w http.ResponseWriter, r *http.Request) {
	var (
		src []byte
		err error
	)
	// Make sure we have the right http Method
	if r.Method != "POST" {
		writerError(w, http.StatusMethodNotAllowed)
		return
	}

	// Make sure we can determine permissions before reading
	// post data.
	username := s.getUsername(r)
	if username == "" {
		writeError(w, http.StatusUnauthorized)
		return
	}
	roles, ok := s.getUserRoles(username)
	if ok == false {
		writeError(w, http.StatusUnauthorized)
		return
	}
	// We need to get the submitted object before checking
	// isAllowed.
	src, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("%s %s", r.URL.Path, err)
		writeError(w, http.StatusNotAcceptable)
		return
	}
	// We only accept content in JSON form with /create.
	object := make(map[string]interface{})
	decoder := json.NewDecoder(r.Body)
	decoder.UseNumber()
	err := decoder.Unmarshal(src, &object)
	if err != nil {
		log.Printf("%s %s", r.URL.Path, err)
		writeError(w, http.StatusUnsupportedMediaType)
		return
	}
	// Need to apply users/roles/states rules.
	state := getState(object)
	if s.isAllowed(roles, state, UPDATE) == false {
		writeError(w, http.StatusUnauthorized)
		return
	}

	// Now get the proposed key.
	key := getKey(r.URL.Path, "/"+cName+"/update/")

	// Need to make sure this part of the service is behind
	// the mutex.
	if err := safeDatasetOp(c, key, object, UPDATE); err != nil {
		log.Printf("%s %s", r.URL.Path, err)
		writeError(w, http.StatusNotAcceptable)
		return
	}
	log.Printf("%s created %s in %s", username, key, c.Name)
}

// requestDelete is the API version of
//	`dataset Delete COLLECTION_NAME OBJECT_ID`
// except is doesn't actually delete the object. It changes the
// object's `._State` value.
func (s *AndOrService) requestDelete(cName string, c *dataset.Collection, w http.ResponseWriter, r *http.Request) {
	//FIXME: Need to apply users/roles/states rules.
	//FIXME: Need to make sure this part of the service is behind
	// muxtex.
	log.Printf("requestDelete(%q, ...) not implemented", cName)
	http.Error(w, "Internal Server error", http.StatusInternalServerError)
}

// requestAssignment retrieves an object, updates ._State and
// writes it back out.
func (s *AndOrService) requestAssignment(cName string, c *dataset.Collection, w http.ResponseWriter, r *http.Request) {
}

// FIXME: requestAttach is the API version of
//	`dataset attach COLLECTION_NAME OBJECT_ID FILENAMES`
func (s *AndOrService) requestAttach(cName string, c *dataset.Collection, w http.ResponseWriter, r *http.Request) {
	//FIXME: Need to apply users/roles/states rules.
	//FIXME: Need to make sure this part of the service is behind
	// muxtex.
	log.Printf("requestAttach(%q, ...) not implemented", cName)
	http.Error(w, "Internal Server error", http.StatusInternalServerError)
}

// FIXME: requestAttachments is the API version of
//	`dataset attachments COLLECTION_NAME OBJECT_ID`
func (s *AndOrService) requestAttachments(cName string, c *dataset.Collection, w http.ResponseWriter, r *http.Request) {
	//FIXME: Need to apply users/roles/states rules.
	//FIXME: Need to make sure this part of the service is behind
	// muxtex.
	log.Printf("requestAttachments(%q, ...) not implemented", cName)
	http.Error(w, "Internal Server error", http.StatusInternalServerError)
}

// FIXME: requestDetach is the API version of
//	`dataset detach COLLECTION_NAME OBJECT_ID FILENAME`
func (s *AndOrService) requestDetach(cName string, c *dataset.Collection, w http.ResponseWriter, r *http.Request) {
	//FIXME: Need to apply users/roles/states rules.
	log.Printf("requestDetach(%q, ...) not implemented", cName)
	http.Error(w, "Internal Server error", http.StatusInternalServerError)
}

// FIXME: requestPrune is the API version of
//	`dataset prune COLLECTION_NAME OBJECT_ID`
func (s *AndOrService) requestPrune(cName string, c *dataset.Collection, w http.ResponseWriter, r *http.Request) {
	//FIXME: Need to apply users/roles/states rules.
	//FIXME: Need to make sure this part of the service is behind
	// muxtex.
	log.Printf("requestPrune(%q, ...) not implemented", cName)
	http.Error(w, "Internal Server error", http.StatusInternalServerError)
}

// addAccessRoute makes a route require an authentication mechanism,
// currently this is BasicAUTH but will likely become JWT.
func addAccessRoute(a *wsfn.Access, p string) {
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
	cName := strings.TrimSuffix(c.Name, ".ds")
	access := s.Access
	//NOTE: We create a function handler based on on the
	// current collection being processed.
	log.Printf("Adding collection %q", c.Name)
	base := "/" + path.Base(cName)
	log.Printf("Adding access route %q", base)
	if s.IsAccessRestricted() {
		addAccessRoute(access, base)
	}
	// End points based on dataset
	p := base + "/keys"
	log.Printf("Adding handler %s", p)
	mux.HandleFunc(p, func(w http.ResponseWriter, r *http.Request) {
		s.requestKeys(cName, c, w, r)
	})
	// dataset object handling
	p = base + "/create/"
	log.Printf("Adding handler %s", p)
	mux.HandleFunc(p, func(w http.ResponseWriter, r *http.Request) {
		s.requestCreate(cName, c, w, r)
	})
	p = base + "/read/"
	log.Printf("Adding handler %s", p)
	mux.HandleFunc(p, func(w http.ResponseWriter, r *http.Request) {
		s.requestRead(cName, c, w, r)
	})
	p = base + "/update/"
	mux.HandleFunc(p, func(w http.ResponseWriter, r *http.Request) {
		s.requestUpdate(cName, c, w, r)
	})
	p = base + "/delete/"
	mux.HandleFunc(p, func(w http.ResponseWriter, r *http.Request) {
		s.requestDelete(cName, c, w, r)
	})

	/*FIXME: add these after basic object ops are working.
	// dataset attachment handling
	p = base + "/attach/"
	mux.HandleFunc(p, func(w http.ResponseWriter, r *http.Request) {
		s.requestAttach(cName, c, w, r)
	})
	p = base + "/attachments/"
	mux.HandleFunc(p, func(w http.ResponseWriter, r *http.Request) {
		s.requestAttachments(cName, c, w, r)
	})
	p = base + "/detach/"
	mux.HandleFunc(p, func(w http.ResponseWriter, r *http.Request) {
		s.requestDetach(cName, c, w, r)
	})
	p = base + "/prune/"
	mux.HandleFunc(p, func(w http.ResponseWriter, r *http.Request) {
		s.requestPrune(cName, c, w, r)
	})
	*/

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
