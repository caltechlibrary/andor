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

// requestKeys is the API version of `dataset keys COLLECTION_NAME`
// We only support GET on keys.
func requestKeys(cName string, c *dataset.Collection, w http.ResponseWriter, r *http.Request) {
	//FIXME: Need to apply user/workflow/queue rules.
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

// requestObject is the API version of
//     `dataset read -c -p COLLECTION_NAME KEY`
func requestObject(cName string, c *dataset.Collection, w http.ResponseWriter, r *http.Request) {
	//FIXME: Need to apply user/workflow/queue rules.
	key := strings.TrimPrefix(r.URL.Path, "/"+cName+"/objects/")
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

func addAccessRoute(a *wsfn.Access, p string) {
	if a != nil {
		if a.Routes == nil {
			a.Routes = []string{}
		}
		a.Routes = append(a.Routes, p)
	}
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

	for cName, c := range s.Collections {
		//NOTE: We create a function handler based on on the
		// current collection being processed.
		log.Printf("Adding %q collection handlers", cName)
		p := "/" + path.Base(cName) + "/objects/"
		addAccessRoute(access, p)
		mux.HandleFunc(p, func(w http.ResponseWriter, r *http.Request) {
			// Do we have an object request or keys request?
			if strings.Compare(r.URL.Path, p) == 0 {
				requestKeys(cName, c, w, r)
				return
			}
			if strings.HasPrefix(r.URL.Path, p) {
				requestObject(cName, c, w, r)
				return
			}
			// Unsupported request ...
			http.NotFound(w, r)
		})
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
