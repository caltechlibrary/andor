package andor

import (
	"fmt"
	"io"
	"strings"

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
	args = args[1:]
	switch verb {
	case "init":
		if len(args) == 0 {
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
	case "load-workflow":
		if len(args) == 0 {
			fmt.Fprintf(eOut, "Missing workflow TOML file name")
			return 1
		}
		//NOTE: We can actually read JSON or TOML files ...
		err := LoadWorkflow(args)
		if err != nil {
			fmt.Fprintf(eOut, "Can't read %s, %s", args[1], err)
			return 1
		}
		return 0
	case "list-workflow":
		objects, err := ListWorkflow(args)
		if err != nil {
			fmt.Fprintf(eOut, "Can't read %s, %s", args[2], err)
			return 1
		}
		for _, obj := range objects {
			fmt.Fprintln(out, "#")
			fmt.Fprintf(out, "# Workflow: %s\n", obj.Key)
			fmt.Fprintf(out, "# Display Name: %s\n", obj.Name)
			fmt.Fprintln(out, "#")
			fmt.Fprintf(out, "%s\n\n", obj.String())
		}
		return 0
	case "remove-workflow":
		if len(args) == 0 {
			fmt.Fprintf(eOut, "Missing workflow name to remove")
			return 1
		}
		err := RemoveWorkflow(args)
		if err != nil {
			fmt.Fprintf(eOut, "%s\n", err)
			return 1
		}
		return 0
	case "load-user":
		if len(args) == 0 {
			fmt.Fprintf(eOut, "Missing user TOML file to load")
			return 1
		}
		fmt.Fprintf(eOut, "%q not implemented\n", verb)
		return 1
	case "list-user":
		fmt.Fprintf(eOut, "%q not implemented\n", verb)
		return 1
	case "remove-user":
		if len(args) == 0 {
			fmt.Fprintf(eOut, "Missing user id to remove")
			return 1
		}
		fmt.Fprintf(eOut, "%q not implemented\n", verb)
		return 1
	case "config":
		//FIXME: Need to output an example TOML configuration file
		fmt.Fprintf(eOut, "%q not implemented\n", verb)
		return 1
	case "start":
		if len(args) == 0 {
			fmt.Fprintf(eOut, "Missing TIML config filename")
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
