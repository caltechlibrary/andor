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
	"fmt"
	"log"
	"net/http"
	"net/url"
	"path"
	"sort"
	"strings"

	// Caltech Library Packages
	"github.com/caltechlibrary/dataset"
)

// requestKeys is the API version of `dataset keys COLLECTION_NAME`
func requestKeys(cName string, c *dataset.Collection, p string, w http.ResponseWriter, r *http.Request) {
	keys := c.Keys()
	sort.Strings(keys)
	log.Printf("requestKeys(%q, obj, %q, obj, obj) not implemented", cName, p)
}

// requestObject is the API version of
//     `dataset read -c -p COLLECTION_NAME KEY`
func requestObject(cName string, c *dataset.Collection, p string, w http.ResponseWriter, r *http.Request) {
	log.Printf("requestObject(%q, obj, %q, obj, obj) not implemented", cName, p)
}

// RunService runs the http/https web service of AndOr.
func RunService(s *AndOrService) error {
	u := new(url.URL)
	u.Scheme = s.Scheme
	u.Host = s.Host + ":" + s.Port

	log.Printf("Have %d collection(s)", len(s.Collections))
	for cName, c := range s.Collections {
		log.Printf("Adding %q collection handlers", cName)
		//NOTE: We create a function handler based on on the
		// current collection being processed.
		p := "/" + path.Base(cName) + "/objects/"
		http.HandleFunc(p, func(w http.ResponseWriter, r *http.Request) {
			// Do we have an object request or keys request?
			if strings.Compare(r.URL.Path, p) == 0 {
				requestKeys(cName, c, p, w, r)
				return
			}
			if strings.HasPrefix(r.URL.Path, p) {
				requestObject(cName, c, r.URL.Path, w, r)
				return
			}
			// Unsupported request ...
			//FIXME: need to log misses
			http.NotFound(w, r)
		})
	}
	if len(s.Htdocs) > 0 {
		fs := htdocsFileSystem{http.Dir(s.Htdocs)}
		http.Handle("/", http.FileServer(fs))
	}
	hostname := fmt.Sprintf("%s:%s", u.Hostname(), u.Port())
	log.Printf("Starting service %s", hostname)
	switch s.Scheme {
	case "http":
		return http.ListenAndServe(hostname, nil)
	case "https":
		return http.ListenAndServeTLS(hostname, s.CertPEM, s.KeyPEM, nil)
	default:
		return fmt.Errorf("%q url scheme not supported", s.Scheme)
	}
}
