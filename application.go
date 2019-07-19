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
		for _, cName := range args[1:] {
			_, err := dataset.InitCollection(cName)
			if err != nil {
				fmt.Fprintf(eOut, "%s\n", err)
				return 1
			}
		}
		//FIXME: if files don't exist write out example
		// workflows.toml and users.toml file so they can be
		// easily edited to setup access.
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
