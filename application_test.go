package andor

import (
	"bytes"
	"path"
	"strings"
	"testing"
)

func TestApplication(t *testing.T) {
	input := []byte{}
	output := []byte{}
	errorOut := []byte{}
	Stdin := bytes.NewReader(input)
	Stdout := bytes.NewBuffer(output)
	Stderr := bytes.NewBuffer(errorOut)

	andorTOML := path.Join("testdata", "andor.toml")
	appName := "ApplicationTest"
	args := []string{}

	// Testing init without parameters, then with
	args = append(args, "init")
	if r := Application(andorTOML, appName, args, Stdin, Stdout, Stderr); r != 1 {
		t.Errorf("expected 1, got %d for %s %s", r, appName, strings.Join(args, " "))
	}

	// Now envoke with repository names
	args = append(args, "test_repo.ds", "test_users.AndOr", "test_workflows.AndOr")
	if r := Application(andorTOML, appName, args, Stdin, Stdout, Stderr); r != 0 {
		t.Errorf("Expected application return 0, got %d for %s %s", r,
			appName, strings.Join(args, " "))
	}

	//FIXME: make sure we have our three collections initialized.
}
