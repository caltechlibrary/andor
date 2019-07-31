package andor

import (
	"net/http"
	"os"
	"strings"
)

//
// The following is based on the Golang docs file hiding example.
// See https://golang.org/pkg/net/http/#example_FileServer_dotFileHiding
//

func isDotFile(name string) bool {
	parts := strings.Split(name, "/")
	for _, part := range parts {
		if strings.HasPrefix(part, ".") {
			return true
		}
	}
	return false
}

// htdocsFiles is the http.File use in htdocsFilesSystem.
// It is used to wrap the Readdir method of http.File so that we can
// remove files and directories that start with a period from its output.
type htdocsFiles struct {
	http.File
}

// Readdir is a wrapper around the Readdir method of the embedded File
// that filters out all files that start with a period in their name.
func (f htdocsFiles) Readdir(n int) (fis []os.FileInfo, err error) {
	files, err := f.File.Readdir(n)
	for _, file := range files { // Filters out the dot files
		if !strings.HasPrefix(file.Name(), ".") {
			fis = append(fis, file)
		}
	}
	return
}

// htdocsFileSystem is an http.FileSystem that hides
// hidden "dot files" from being served.
type htdocsFileSystem struct {
	http.FileSystem
}

// Open is a wrapper around the Open method of the embedded FileSystem
// that serves a 403 permission error when name has a file or directory
// with whose name starts with a period in its path.
func (fs htdocsFileSystem) Open(name string) (http.File, error) {
	if isDotFile(name) { // If dot file, return 403 response
		return nil, os.ErrPermission
	}

	file, err := fs.FileSystem.Open(name)
	if err != nil {
		return nil, err
	}
	return htdocsFiles{file}, err
}
