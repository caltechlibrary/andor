+++
markup = "mmark"
+++

# Workflow Scheme

Workflows encapsulate the concept of a queue[^1]
and access crontol.  If a user if a member of a
workflow they gain the permissions of that workflow for objects
defined viewable in that workflow.  If a user
is assigned to a workflow that is not defined they gain no additional
permissions.  A user who is a member of a workflow acquires the 
permissions of that workflow for objects associated with that workflow(s).  

Here's an example of a workflow for someone allowed to currate any object
in a collection.

```json
    {
        "workflow_name": "Currator",
        "workflow_id": "currator"
        "object_permissions": [
            "create",
            "read",
            "update"
        ]
        "assign_to": [ "*" ],
        "view_object_ids": [ "*" ]
    }
```

workflow\_name
: (string, optional) the human readable name for the workflow

workflow\_id
: (string, required, must be unique) the id for this workflow

object\_permissions
: (array of string, default is empty list) this is where you give object level permissions for create, read and update. For the object level permissions objects must be viewable in the workflow. See "view\_objects" field.


assign\_to
: (array of string, default is empty list) this holds the ids of the workflows you are allow to assign an object to, if the "\*" string is included it means this workflow can assign an object to any other workflow

view\_objects
: (array of string, default is empty list) this holds the workflow ids of the objects you're allow to view, this is how we create a queue of objects. Ifthe workflow is listed as "\*" it means you can see any object in the collection.  Note if you can see the object then the object\_permissions apply. 


In our example the workflow name is "Currator" and the id is "currator".
Because "assign\_to" contains the "\*" string. A user who is associated
with the workflow "currator" (you can also think of this as a queue)
can move objects to any other workflow. A currator can also see all objects
because "view\_objects" also has the "\*" string.


Normally a workflow is defined more granularly. A typical system
would have a queue for currator, published and delete. If you
limit the permissions for each of these you can create an integrated
workflow.

If you allow deleting objects you can create a workflow that
has no object permissions in it. A "delete" workflow could look like

```json
    {
        "workflow_name": "Deleted objects",
        "workflow_id": "delete"
        "object_permissions": []
        "assign_to": [],
        "view_object_ids": []
    }
```

If a user has "delete" in their "member\_of" OR they have
"delete" in the "asign\_to" list of a workflow they are a
member of they could put the objects into the delete workflow.
If the a workflow they belong to can see all objects (e.g. our
currator example) the object would still be viewable but be in the workflow state (queue) of deleted. Otherwise the objects in "delete" workflow
state would not be visible.

Another common workflow would be "published". The published workflow
could be defined this way.


```json
    {
        "workflow_name": "Publishd Objects",
        "workflow_id": "published"
        "object_permissions": [ "read" ]
        "assign_to": [ ],
        "view_object_ids": [ "published" ]
    }
```

If a user's "member\_of" field contained "published" they would
then be allowed to see all objects held in the "published" workflow.
Notice that we also give "published" the read permission at the
object level. This is needed to display the objects details
(e.g. a title field) as well as the id. "view\_objects" only means
you can view the id, that it the id will be returned in a list of 
ids.

Likewise if you wanted to make your "published" items viewable by 
the public you create a record for the "anonymous" user and
add "published" to the anonymous users' "member\_of" field. That
would allow unauthenticated users to see published objects.

The "assign\_to" controls where you can move an object to and
"view\_objects" controls which objects you can see while
in this workflow. A object's workflow assignment also makes
a workflow behave as a queue.

Object permissions can be "create", "update" and "read".  **AndOr**
doesn't support delating objects.  You can create the 
illusion of deletion by be clever with your workflow definitions.

[^1]: objects have a field called "\_WorkQueue" that holds the name of the workflow queue they are currently associated with
