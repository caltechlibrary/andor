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
	"io"
	"os"
	"strings"

	// Caltech Library packages
	"github.com/caltechlibrary/dataset"
)

// Application runs the command line interaction
// for AndOr. It returns an exit status (e.g. 0
// if everything is OK, non-zero for error).
func Application(appName, andorFile string, args []string, in io.Reader, out io.Writer, eOut io.Writer) int {
	var (
		collections []string
		rolesFile   string
		usersFile   string
		verb        string
	)
	if len(args) == 0 {
		fmt.Fprintf(eOut, "Expecting either 'init' or 'start' action\n")
		return 1
	}
	verb = args[0]
	args = args[1:]
	switch verb {
	case "init":
		if len(args) == 0 {
			// If see if repository.ds exists, if not create it.
			if _, err := os.Stat("repository.ds"); os.IsNotExist(err) {
				args = append(args, "repository.ds")
			} else if err != nil {
				fmt.Fprintf(eOut, "No repository name(s) provided, %s\n", err)
				return 1
			} else {
				fmt.Fprintf(eOut, "Using existing %q\n", "reposotory.ds")
			}
		}
		for _, cName := range args {
			if cName != "" {
				_, err := dataset.InitCollection(cName)
				if err != nil {
					fmt.Fprintf(eOut, "%s\n", err)
					return 1
				}
				collections = append(collections, cName)
			}
		}
		// NOTE: We should generate example andor.toml, roles.toml,
		// and users.toml so it is easy to finish setting AndOr.
		if _, err := os.Stat(andorFile); os.IsNotExist(err) {
			err = GenerateAndOr(andorFile, collections)
			if err != nil {
				fmt.Fprintf(eOut, "generating %q, %s\n", andorFile, err)
				os.Exit(1)
			}
			rolesFile = "roles.toml"
			usersFile = "users.toml"
		} else {
			fmt.Fprintf(eOut, "Using existing %q\n", andorFile)
			if s, err := LoadAndOr(andorFile); err != nil {
				fmt.Fprintf(eOut, "Found invalid %q, %s\n", andorFile, err)
				os.Exit(0)
			} else {
				rolesFile = s.RolesFile
				usersFile = s.UsersFile
			}
		}
		if _, err := os.Stat(rolesFile); os.IsNotExist(err) {
			err = GenerateRoles(rolesFile)
			if err != nil {
				fmt.Fprintf(eOut, "generating %q, %s\n", rolesFile, err)
				os.Exit(1)
			}
		} else {
			fmt.Fprintf(eOut, "Using existing %q\n", rolesFile)
		}
		if _, err := os.Stat(usersFile); os.IsNotExist(err) {
			err = GenerateUsers(usersFile)
			if err != nil {
				fmt.Fprintf(eOut, "generating %q, %s\n", usersFile, err)
				os.Exit(1)
			}
		} else {
			fmt.Fprintf(eOut, "Using existing %q\n", usersFile)
		}
		// Now read back our toml files (existing or the ones we created)
		s, err := LoadAndOr(andorFile)
		if err != nil {
			fmt.Fprintf(eOut, "WARNING (1) %q, invalid, %s\n", andorFile, err)
			return 1
		}
		if _, _, err := LoadRoles(s.RolesFile); err != nil {
			fmt.Fprintf(eOut, "WARNING (2) %q, invalid, %s\n", s.RolesFile, err)
			return 1
		}
		if _, err := LoadUsers(s.UsersFile); err != nil {
			fmt.Fprintf(eOut, "WARNING (3) %q, invalid, %s\n", s.UsersFile, err)
			return 1
		}
		fmt.Fprintln(out, "OK")
		return 0
	case "check":
		s, err := LoadAndOr(andorFile)
		if err != nil {
			fmt.Fprintf(eOut, "Problem with %s\n", err)
			return 1
		}
		fmt.Printf("DEBUG s -> %+v\n", s)
		if _, _, err := LoadRoles(s.RolesFile); err != nil {
			fmt.Fprintf(eOut, "Problem with %s\n", err)
			return 1
		}
		if _, err := LoadUsers(s.UsersFile); err != nil {
			fmt.Fprintf(eOut, "Problem with %s\n", err)
			return 1
		}
		fmt.Fprintln(out, "OK")
		return 0
	case "start":
		if len(args) > 0 {
			andorFile = args[0]
		}
		service, err := LoadAndOr(andorFile)
		if err != nil {
			fmt.Fprintf(eOut, "Error starting service, %s\n", err)
			return 1
		}
		if err := service.Start(); err != nil {
			fmt.Fprintf(eOut, "%s\n", err)
			return 1
		}
		return 0
	default:
		fmt.Fprintf(eOut, "Don't understand \"%s %s\"", verb, strings.Join(args, " "))
		return 1
	}
	fmt.Fprintln(out, "OK")
	return 0
}
