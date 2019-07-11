package andor

import (
	"fmt"
	"io"
)

// Application runs the command line interaction
// for AndOr. It returns an exit status (e.g. 0
// if everything is OK, non-zero for error).
func Application(appName string, args []string, in io.Reader, out io.Writer, eOut io.Writer) int {
	fmt.Fprintf(eOut, "%s not implemented.", appName)
	return 1
}
