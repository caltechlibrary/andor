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
func Application(andorTOML, appName string, args []string, in io.Reader, out io.Writer, eOut io.Writer) int {
	if len(args) == 0 {
		fmt.Fprintf(eOut, "Missing a verb like init, gen-user, gen-workflow, start\n")
		return 1
	}
	verb := args[0]
	args = args[1:]
	switch verb {
	case "init":
		if len(args) == 0 {
			fmt.Fprintf(eOut, "Missing collection name(s)")
			return 1
		}
		workflowsTOML := "workflows.toml"
		usersTOML := "users.toml"
		collections := []string{}
		for _, cName := range args {
			_, err := dataset.InitCollection(cName)
			if err != nil {
				fmt.Fprintf(eOut, "%s\n", err)
				return 1
			}
			collections = append(collections, cName)
		}
		// NOTE: We should generate example andor.toml, workflows.toml,
		// and users.toml so it is easy to finish setting AndOr.
		if _, err := os.Stat(andorTOML); os.IsNotExist(err) {
			err = GenerateAndOrTOML(andorTOML, workflowsTOML, usersTOML, collections)
			if err != nil {
				fmt.Printf("generating %q, %s", andorTOML, err)
				os.Exit(1)
			}
		} else {
			s, err := LoadAndOr(andorTOML)
			if err != nil {
				fmt.Printf("WARNING %q, invalid, %s\n", andorTOML, err)
			} else {
				if s.WorkflowsTOML != "" {
					workflowsTOML = s.WorkflowsTOML
				}
				if s.UsersTOML != "" {
					usersTOML = s.UsersTOML
				}
			}
		}
		if _, err := os.Stat(workflowsTOML); os.IsNotExist(err) {
			err = GenerateWorkflowsTOML(workflowsTOML)
			if err != nil {
				fmt.Printf("generating %q, %s", workflowsTOML, err)
				os.Exit(1)
			}
		}
		if _, err := os.Stat("users.toml"); os.IsNotExist(err) {
			err = GenerateUsersTOML(usersTOML)
			if err != nil {
				fmt.Printf("generating %q, %s", usersTOML, err)
				os.Exit(1)
			}
		}
	case "check-config":
		if _, err := LoadAndOr(andorTOML); err != nil {
			fmt.Fprintf(eOut, "Problem with %s, %s", andorTOML, err)
			return 1
		}
		return 0
	case "start":
		if len(args) == 0 {
			fmt.Fprintf(eOut, "Missing TOML config filename")
			return 1
		}
		fmt.Fprintf(eOut, "%q not implemented\n", verb)
		return 1
	default:
		fmt.Fprintf(eOut, "Don't understand \"%s %s\"", verb, strings.Join(args, " "))
		return 1
	}
	return 0
}
