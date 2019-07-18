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
	"encoding/json"
	"fmt"
)

// ObjectKey inspects a map[string]interface{} (an object)
// and returns the `._Key` value if it is set. An empty
// string is return if it is not set.
func ObjectKey(obj map[string]interface{}) string {
	if queue, ok := obj["_Key"]; ok == true {
		switch queue.(type) {
		case json.Number:
			j := queue.(json.Number)
			return j.String()
		case int:
			i := queue.(int)
			return fmt.Sprintf("%d", i)
		case int64:
			i := queue.(int64)
			return fmt.Sprintf("%d", i)
		case float64:
			f := queue.(float64)
			return fmt.Sprintf("%f", f)
		case string:
			return queue.(string)
		}
	}
	return ""
}

// ObjectQueue inspects a map[string]inteface{} (an object)
// and returns the `._Queue` value if it is set. An empty
// string is returned if it is not set.
func ObjectQueue(obj map[string]interface{}) string {
	if queue, ok := obj["_Queue"]; ok == true {
		switch queue.(type) {
		case string:
			return queue.(string)
		}
	}
	return ""
}
