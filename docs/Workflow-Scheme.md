+++
markup = "mmark"
+++

# Workflow Scheme

Workflows encapsulate the concept of a queue[^1]
and capabilities (e.g. access control).  A queue if it 
is defined in at least one workflow.  **AndOr** 
supports basic CRUD[^2] operations on any object associated 
with a workflow's queue. In addition to CRUD operations 
an additional capability is assignment. 
Assignment is how we pass objects between workflow 
states (e.g. "deposit", "review", "publish").

An addition queue exists without definition. That queue
is "deleted". Like EPrints[^3] **AndOr** doesn't delete
objects from disc. They continue to exist in a trashbin
like state.  It is trivial to write a garbage collected
that operates on objects having a `._Queue` value of 
"deleted". 

Here's an example object of defining a "deposit" workflow
associated with a "deposit" queue. A user working in 
the "deposit" workflow create objects but can not read them. 
When an object in a workflow is assigned the `._Queue`
value listed in the wrokflow. So our depository user
can create objects in the "deposit" queue but they 
have no other rights beyond that.

```json
    {
        "workflow_name": "Depositor",
        "workflow_id": "deposit",
        "queue": "review",
        "create": true,
        "read": false,
        "update": false,
        "delete": false,
        "assign_to": [ ]
    }

```

Someone who is a member of the "review" workflow and who should
be is allowed to read, update, and delete objects (but not create) 
and move them to a "published" queue would look like --

```json
    {
        "workflow_name": "Review",
        "workflow_id": "review"
        "queue": "review",
        "create": false,
        "read" : true,
        "update": true,
        "delete": true,
        "assign_to": [ "review", "published" ]
    }
```

Finally our "published" workflow would only define the
ability to read objects in the queue "published". 

```json
    {
        "workflow_name": "Published",
        "workflow_id": "published"
        "queue": "published",
        "create": false,
        "read" : true,
        "update": false,
        "delete": false,
        "assign_to": [ ]
    }
```

What if we want to pull an object back from published state?
We can create a workflow called "Unpublish" that is allowed to
operate on the "published" queue. It has the right to read
objects in "published" queue and to assign objects back to
the "review" queue.

```json
    {
        "workflow_name": "Unpublish",
        "workflow_id": "unpublish"
        "queue": "published",
        "create": false,
        "read" : true,
        "update": false,
        "delete": false,
        "assign_to": [ "published", "review" ]
    }
```

workflow\_name
: (string, optional) the human readable name for the workflow

workflow\_id
: (string, required, must be unique) the id for this workflow

queue
: (string, required, may be unique) the associated queue the workflow describes

create
: (bool, defaults to false if not defined) the ability to create an object in the queue

read
: (boo, defaults to false if not defined) the ability to read an object in the queue

update
: (bool, defaults to false if not defined) the ability to update (edit) an object in the queue

delete
: (bool, defaults to false if not defined) the ability to "delete" an object (really move to a trashbin state)

assign\_to
: (list of string, may be empty) the list of queues that a workflow may assign objects into



[^1]: objects have a field called "\_Queue" that holds the name of the workflow queue they are currently associated with

[^2]: CRUD, refer to "create", "read", "update", "delete" operations on a object

[^3]: EPrints is an excellent Open Source Repository system from South Hampton University

