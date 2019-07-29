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
	"bytes"
	"fmt"
	"io/ioutil"
	"path"
	"strings"
	"testing"
)

// TestLoadWorkflows ...
func TestLoadWorkflows(t *testing.T) {
	workflowFile := path.Join("testdata", "workflows.toml")
	w, q, err := LoadWorkflows(workflowFile)
	if err != nil {
		t.Errorf("Failed to load %q, %s", workflowFile, err)
	}
	if len(w) != 3 {
		t.Errorf("expected 3 workflows, got %d", len(w))
	}
	if len(q) != 3 {
		t.Errorf("expected 3 queues, got %d", len(q))
	}
	for _, wName := range []string{"draft", "review", "published"} {
		if workflow, ok := w[wName]; ok == false {
			t.Errorf("expected %q, not found -> %+v", wName, w)
			if strings.Compare(workflow.Queue, wName) != 0 {
				t.Errorf("expected %q, got %q for Queue", wName, workflow.Queue)
			}
		}
	}
}

// TestBytes for workflow structs
func TestBytes(t *testing.T) {
	workflowFile := path.Join("testdata", "workflows2.toml")
	workflowSrc, err := ioutil.ReadFile(workflowFile)
	if err != nil {
		t.Errorf("expected to read %s, %s", workflowFile, err)
		t.FailNow()
	}
	w, _, err := LoadWorkflows(workflowFile)
	if err != nil {
		t.Errorf("expected to load %q, got %s", workflowFile, err)
	}
	src := []byte{}
	for _, k := range []string{"draft", "review", "published"} {
		v, _ := w[k]
		src = append(src, []byte(fmt.Sprintf("[%s]\n", k))...)
		src = append(src, v.Bytes()...)
		src = append(src, []byte("\n")...)
	}
	if len(src) == 0 {
		t.Errorf("expected a []byte with data for %+v", w)
	}
	if bytes.Compare(workflowSrc, src) != 0 {
		t.Errorf("expected sources to match, got\n%s\n", src)
	}
}

// TestString for workflow structs
func TestString(t *testing.T) {
	workflowFile := path.Join("testdata", "workflows2.toml")
	workflowSrc, err := ioutil.ReadFile(workflowFile)
	if err != nil {
		t.Errorf("expected to read %s, %s", workflowFile, err)
		t.FailNow()
	}
	w, _, err := LoadWorkflows(workflowFile)
	if err != nil {
		t.Errorf("expected to load %q, got %s", workflowFile, err)
	}
	s := []string{}
	for _, k := range []string{"draft", "review", "published"} {
		v, _ := w[k]
		s = append(s, fmt.Sprintf("[%s]\n", k))
		s = append(s, v.String())
		s = append(s, "\n")
	}
	src := strings.Join(s, "")
	if len(src) == 0 {
		t.Errorf("expected a string with data for %+v", w)
	}
	if strings.Compare(string(workflowSrc), src) != 0 {
		t.Errorf("expected sources to match, got\n%s\n", src)
	}
}
