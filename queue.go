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
)

// Queue describes a queue's state, the object ideas and
// and sorted id lists and workflows associated with the queue.
type Queue struct {
	// Name holds the name of the queue
	Name string `json:"name"`
	// Workflows thats operating on this queue
	Workflows []string `json:"workflows"`
	// A list of ids currently in the queue
	ObjectIDs []string `json:"obect_ids"`
}

// QueueMap holds a map of queues, it is responsible for
// tracking the state of all known objects in *AndOr*
type QueueMap map[string]*Queue

func NewQueueMap() *QueueMap {
	qMap := make(QueueMap)
	return qMap
}

func (qMap *QueueMap) AddQueue(q *Queue) error {
	if qMap == nil {
		qMap = NewQueueMap()
	}
	key := q.Name
	if key != "" {
		qMap[key] = q
		return nil
	}
	return fmt.Errorf("missing queue name")
}

// NewQueue creates a new queue and populates the known
// workflows.
func NewQueue(queueName string, workflows []string) *Queue {
	q := new(Queue)
	q.Name = queueName
	q.Workflows = workflows[:]
	return q
}

// objKey inspects a map[string]interface{} (an object)
// and returns the `._Key` value if it is set, otherwise an
// empty string.
func objKey(obj map[string]interface{}) string {
	if key, ok := obj["_Key"]; ok == true {
		switch key.(type) {
		case json.Number:
			return key.(json.Number).String()
		case float64:
			return fmt.Sprintf("%f", key.(float64))
		case int:
			return fmt.Sprintf("%d", key.(int))
		case string:
			return key.(string)
		}
	}
	return ""
}

// objQueue inspects a map[string]inteface{} (an object)
// and returns the `._Queue` value if it is set. An empty
// string is returned if it is not set.
func objQueue(obj map[string]interface{}) string {
	if queue, ok := obj["_Queue"]; ok == true {
		switch queue.(type) {
		case string:
			return queue.(string)
		}
	}
	return ""
}

// addObject takes a map[string]interface{} (an object)
// and adds its `._Key` value (if not empty) to the queue.
// Creates/Updates a `._Queue` to match queue name.
// AddObject enforces permissions and returns an error is
// assignment of object to queue is not allowed
func (q *Queue) AddObject(user *User, obj map[string]interface{}) (obj, error) {
	// First check to make sure user can add the object.
	addOK := false
	for _, workflow := range q.Workflows {
		if CanAssign(user, workflow, obj) {
			addOK = true
			break
		}
	}
	if addOK == false {
		return obj, fmt.Fprintf("assignment denied")
	}

	key := objKey(obj)
	if key != "" {
		q.ObjectIDs = append(q.ObjectIDs, key)
		obj["_Queue"] = q.Name
		return obj, nil
	}
	return fmt.Errorf("object missing queue")
}
