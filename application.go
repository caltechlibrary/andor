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
func Application(appName, andorTOML string, args []string, in io.Reader, out io.Writer, eOut io.Writer) int {
	var (
		collections   []string
		workflowsTOML string
		usersTOML     string
		verb          string
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
		// NOTE: We should generate example andor.toml, workflows.toml,
		// and users.toml so it is easy to finish setting AndOr.
		if _, err := os.Stat(andorTOML); os.IsNotExist(err) {
			err = GenerateAndOrTOML(andorTOML, collections)
			if err != nil {
				fmt.Fprintf(eOut, "generating %q, %s\n", andorTOML, err)
				os.Exit(1)
			}
			workflowsTOML = "workflows.toml"
			usersTOML = "users.toml"
		} else {
			fmt.Fprintf(eOut, "Using existing %q\n", andorTOML)
			if s, err := LoadAndOr(andorTOML); err != nil {
				fmt.Fprintf(eOut, "Found invalid %q, %s\n", andorTOML, err)
				os.Exit(0)
			} else {
				workflowsTOML = s.WorkflowsTOML
				usersTOML = s.UsersTOML
			}
		}
		if _, err := os.Stat(workflowsTOML); os.IsNotExist(err) {
			err = GenerateWorkflowsTOML(workflowsTOML)
			if err != nil {
				fmt.Fprintf(eOut, "generating %q, %s\n", workflowsTOML, err)
				os.Exit(1)
			}
		} else {
			fmt.Fprintf(eOut, "Using existing %q\n", workflowsTOML)
		}
		if _, err := os.Stat(usersTOML); os.IsNotExist(err) {
			err = GenerateUsersTOML(usersTOML)
			if err != nil {
				fmt.Fprintf(eOut, "generating %q, %s\n", usersTOML, err)
				os.Exit(1)
			}
		} else {
			fmt.Fprintf(eOut, "Using existing %q\n", usersTOML)
		}
		// Now read back our toml files (existing or the ones we created)
		s, err := LoadAndOr(andorTOML)
		if err != nil {
			fmt.Fprintf(eOut, "WARNING (1) %q, invalid, %s\n", andorTOML, err)
			return 1
		}
		if _, _, err := LoadWorkflows(s.WorkflowsTOML); err != nil {
			fmt.Fprintf(eOut, "WARNING (2) %q, invalid, %s\n", s.WorkflowsTOML, err)
			return 1
		}
		if _, err := LoadUsers(s.UsersTOML); err != nil {
			fmt.Fprintf(eOut, "WARNING (3) %q, invalid, %s\n", s.UsersTOML, err)
			return 1
		}
		fmt.Fprintln(out, "OK")
		return 0
	case "check":
		s, err := LoadAndOr(andorTOML)
		if err != nil {
			fmt.Fprintf(eOut, "Problem with %s\n", err)
			return 1
		}
		if _, _, err := LoadWorkflows(s.WorkflowsTOML); err != nil {
			fmt.Fprintf(eOut, "Problem with %s\n", err)
			return 1
		}
		if _, err := LoadUsers(s.UsersTOML); err != nil {
			fmt.Fprintf(eOut, "Problem with %s\n", err)
			return 1
		}
		fmt.Fprintln(out, "OK")
		return 0
	case "start":
		if len(args) == 0 {
			if _, err := os.Stat(andorTOML); os.IsNotExist(err) {
				fmt.Fprintf(eOut, "Missing %q\n", andorTOML)
				return 1
			}
		} else {
			andorTOML = args[1]
		}
		fmt.Fprintf(eOut, "%q not implemented\n", verb)
		return 1
	default:
		fmt.Fprintf(eOut, "Don't understand \"%s %s\"", verb, strings.Join(args, " "))
		return 1
	}
	fmt.Fprintln(out, "OK")
	return 0
}
