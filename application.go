package andor

import (
	"fmt"
	"io"

	// Caltech Library packages
	"github.com/caltechlibrary/dataset"
)

// Application runs the command line interaction
// for AndOr. It returns an exit status (e.g. 0
// if everything is OK, non-zero for error).
func Application(appName string, args []string, in io.Reader, out io.Writer, eOut io.Writer) int {
	if len(args) == 0 {
		fmt.Fprintf(eOut, "Missing a verb like init, gen-user, gen-workflow, start\n")
		return 1
	}
	verb := args[0]
	switch verb {
	case "init":
		if len(args) < 2 {
			fmt.Fprintf(eOut, "Missing collection name(s)")
			return 1
		}
		for _, cName := range args[1:] {
			_, err := dataset.InitCollection(cName)
			if err != nil {
				fmt.Fprintf(eOut, "%s\n", err)
				return 1
			}
		}
	case "add-workflow":
		//FIXME: We want to also provide an interactive version.
		if len(args) < 3 {
			fmt.Fprintf(eOut, "Missing workflow name and JSON/TOML document")
			return 1
		}
		//FIXME: Need to handle case where we have JSON or TOML document
		workflowName := args[1]
		workflow, err := ReadWorkflowFile(args[2])
		if err != nil {
			fmt.Fprintf(eOut, "Can't read %s, %s", args[2], err)
			return 1
		}
		if err = AddWorkflow(workflowName, workflow); err != nil {
			fmt.Fprintf(eOut, "%s", err)
			return 1
		}
		return 0
	case "list-workflows":
		objects, err := ListWorkflows()
		if err != nil {
			fmt.Fprintf(eOut, "Can't read %s, %s", args[2], err)
			return 1
		}
		for _, obj := range objects {
			fmt.Fprintln(out, "#")
			fmt.Fprintf(out, "# Workflow: %s\n", obj.Name)
			fmt.Fprintln(out, "#")
			fmt.Fprintf(out, "%s\n\n", obj.String())
		}
		return 0
	case "remove-workflows":
		if len(args) < 2 {
			fmt.Fprintf(eOut, "Missing workflow name to remove")
			return 1
		}
		err := RemoveWorkflow(args[1])
		if err != nil {
			fmt.Fprintf(eOut, "%s\n", err)
			return 1
		}
		return 0
	case "add-user":
		fmt.Fprintf(eOut, "%q not implemented\n", verb)
		return 1
	case "add-user-workflow":
		fmt.Fprintf(eOut, "%q not implemented\n", verb)
		return 1
	case "remove-user-workflow":
		fmt.Fprintf(eOut, "%q not implemented\n", verb)
		return 1
	case "list-users":
		fmt.Fprintf(eOut, "%q not implemented\n", verb)
		return 1
	case "remove-user":
		fmt.Fprintf(eOut, "%q not implemented\n", verb)
		return 1
	case "config":
		fmt.Fprintf(eOut, "%q not implemented\n", verb)
		return 1
	case "start":
		fmt.Fprintf(eOut, "%q not implemented\n", verb)
		return 1
	case "restart":
		fmt.Fprintf(eOut, "%q not implemented\n", verb)
		return 1
	case "stop":
		fmt.Fprintf(eOut, "%q not implemented\n", verb)
		return 1
	default:
		fmt.Fprintf(eOut, "%q, unknown verb\n", verb)
		return 1
	}
	return 0
}
