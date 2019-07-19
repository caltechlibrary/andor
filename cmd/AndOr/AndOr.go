//
// AndOr command is built on the "andor" and "dataset" go packages.
// It implements a proof of concept light weight digital object
// repository system.
//
// @Author R. S. Doiel, <rsdoiel@library.caltech.edu>
//
package main

import (
	"flag"
	"fmt"
	"os"
	"path"

	// AndOr package
	"github.com/caltechlibrary/AndOr"
)

var (
	showVersion bool
	showHelp    bool

	helpMsg = `
USAGE 

	%s ACTION [PARAMETERS]

%s is used to manage another digital object
repository. It is a command line 
program. The general syntax is to provide an
action followed by any needed parameters.
Most actions will prompt for parameters so
you don't need to provide them. 

Actions:

  init
            will create a dataset collection.
            three are needed to host AndOr.
            Typically these are repository.ds,
            workflows.ds and users.ds.

  create-user    
            create a new user for AndOr

  create-workflow    
            will create a new workflow

  host
            will run the AndOr web service, if
            a configuration file is not found
            it will prompt for the values to
            create one.

%s %s
`
)

func usage(appName, version string, args []string, exitCode int) {
	out := os.Stdout
	if exitCode != 0 {
		out = os.Stderr
	}
	fmt.Fprintf(out, helpMsg, appName, appName, appName, version)
	os.Exit(exitCode)
}

func main() {
	flag.BoolVar(&showVersion, "v", false, "display version")
	flag.BoolVar(&showVersion, "version", false, "display version")
	flag.BoolVar(&showHelp, "h", false, "display help")
	flag.BoolVar(&showHelp, "help", false, "display help")
	flag.Parse()

	appName := path.Base(os.Args[0])
	args := flag.Args()

	if showVersion {
		fmt.Fprintf(os.Stdout, "%s %s", appName, andor.Version)
		os.Exit(0)
	}
	if showHelp {
		usage(appName, andor.Version, args, 0)
	}
	os.Exit(andor.Application(appName, "andor.toml", args, os.Stdin, os.Stdout, os.Stderr))
}
