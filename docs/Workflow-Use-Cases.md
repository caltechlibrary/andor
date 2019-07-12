
# Workflow Use Cases

Below are use cases exploring how a workflow queue permission
system could be used.

## Use case 1

We would like people (e.g. Jane) to curate the collection.
To curate the whole collection Jane needs to have the
following permissions-- create, read, update, delete,
change workflow queues, and to list all objects
in the collection. To do that we can create a workflow queue called 
"curators".

```json
    {
        "workflow_name": "Curators",
        "workflow_id": "curators",
        "object_permissions": [
            "create",
            "read",
            "update"
        ],
        "assign_to": [ "*" ],
        "view_object_ids": [ "*" ]
    }
```

Notice that we have lists for the "assign\_to" "view\_object\_ids" as
well as "object\_permissions".  If the list contains
a string with an asterisks any workflow will be matched.
The alternative would be to list specific workflow queues, owners.

Jane needs to be a curator. We need to create a user record for her
and to make her a member of the "curators" workflow queue.  Jane's
email address is "jane@example.edu" so we use that as her user id.
Jane's user object would look like.

```json
    {
        "user_id": "jane@example.edu",
        "display_name": "Jane",
        "create_objects_as": "curators",
        "member_of": [ "curators" ]
    }
```

When Jane authenticates with the system she goes from being
"anonymous" to "jane@example.edu" user.  This means she now has the
permissions of a curator. If she creates an object the object will
be assigned to the "curators" workflow queue. Since the curators
workflow queue has change workflow queue permissions for any workflow
queue of objects in the collection can take any action with any
object as needed.


## Use case 2

We would like the public to be able to view the "published" contents
of the of our collection. We can do this by creating a workflow queue called
"published". The published workflow queue will be given read access to the
objects. It'll be allowed to list objects in the "published" workflow queue.

The workflow queue object would look like

```json
    {
        "workflow_name": "Published",
        "workflow_id": "published",
        "collection_permissions": [
            "read"
        ],
        "assign_to": [ ],
        "view_object_ids": [ "published" ]
    }
```

Any member of "published" will be able to read and objects
associated with "published" workflow queue.

Now we need to explicitly associated "anonymous"
with the "published" workflow queue.

```json
    {
        "userid": "anonymous",
        "display_name": "Public",
        "create_objects_as": null,
        "member_of": [ "published" ]
    }
```

When anonymous goes to access an object and the workflow queue associated with
the object is "published" then anonymous will be able to read the object.
Likewise anonymous is not a member of any other workflow queues so they will
have no permissions to access them.

## Use case 3

We want to allow anonymous users to "deposit" objects.  We can
create a workflow queue called "deposit". It should only have create
permissions.  If we add "anonymous" to the "deposit" workflow queue
we could then allow the public to submit records but nothing else.
If anonymous is a member of "published" they will be able to still
read all published records.

```json
    {
        "workflow_name": "Depositor",
        "workflow_id": "deposit",
        "object_permissions": [ "create" ],
        "assign_to": [ ],
        "view_object_ids": [ ]
    }
```

Now if we update our "anonymous" user we can add them to 
the "deposit" workflow queue and associate the created object with
a workflow queue called "deposit".

```json
    {
        "userid": "anonymous",
        "display_name": "Public",
        "create_objects_as": "deposit",
        "member_of": [ "published", "deposit" ]
    }
```

Because the object created will have the workflow queue "deposit" and
our "curators" workflow queue has permissions to list all objects workflow queues
we can treat the "deposit" workflow queue a an inbox to be processed.
If the curators approved the deposit they can change the objects'
workflow queue to "published".

## Use case 4

Creating workflows with workflow queues. We would like our objects to travel
through the following states - deposit, review, then either be
flagged with publish, embargo, and needs curation.

Here are some of our policies we want to enforce.

1. Allow any authenticated user to deposit
2. Allow reviewers to choose between the following status -- deposit, review, published, embargo, rejected and needs further curation
3. If "needs curation" is chosen then the reviewer should no longer be able to see the object
4. The depositor should not see the object after it is "deposited"
5. Reviewers should not be able to change an object only change the workflow queue
6. Reviewers can't create new objects

Jane is a curator. Millie is a reviewer. Millie should not be able
to update the objects but she should be able to list objects that
have been deposited and change objects and workflow queue on an
object to As 

Jane is already in the curators workflow queue previously defined. We need
to define a reviewer workflow queue. Millie will need to be created and be
a member of "reviewer".

```json
    {
        "workflow_name": "Reviewer",
        "workflow_id": "review",
        "object_permissions": [ "read" ],
        "assign_to": [
            "deposit",
            "review",
            "published",
            "embargo",  
            "rejected",
            "currators"
            ],
        "view_object_ids": [
            "deposit",
            "review",
            "published",
            "embargo",
            "rejected"
        ]
    }
```

Millie's email is "mille@example.edu", her account would look like

```json
    {
        "display_name": "Millie",
        "userid": "millie@example.edu",
        "create_objects_as": null,
        "member_of": [ "reviewer" ]
    }
```

If later we decide Millie should be able to create objects then
we can add her the deposit workflow queue. She would not be able to 
edit her deposits but she could change the workflow queue value and list it.


## Use case 5

We want a four level review process. Writers can create 
objects and edit them until they pass them on for review.
The reviewer can either send them back to the writers or pass
them on to the editors for further editing.  Editors can
pass them back to the reviewers or writers or bump them up
to the publisher. Editors can make changes to the
objects. Publishers can do anything.

+ Jane is a publisher, jane@example.org
+ Millie is a editor, millie@example.org
+ Alfred is an reviewer, alfred@example.org
+ Olive is a writer, olive@example.org

In this use case Olive needs to be able to edit her
deposits but not someone else's. We can do this by
create an "olive" workflow queue which and having that as
the default workflow queue when she creates a new object.
We can also associate olive with the writer workflow queue
which can change permissions to reviewer.

Here is Olive's workflow queue

```json
    {
        "workflow_name": "Olive",
        "workflow_id": "olive",
        "object_permissions": [
            "create",
            "read",
            "update"
        ],
        "assign_to": [ ],
        "view_object_ids": [
            "olive"
        ]
    }
```

Here is the general writer's workflow queue object

```json
    {
        "workflow_name": "Writer",
        "workflow_id": "writer",
        "object_permissions": [ ],
        "assign_to": [
            "reviewer"
        ],
        "view_object_ids": [ ]
    }
```

Here is Olive's user object

```json
    {
        "display_name": "Olive",
        "userid": "olive@example.edu",
        "create_objects_as": "olive",
        "member_of": [ "olive", "writer" ]
    }
```

Olive's workflow is then to create, edit update or delete
any object with a workflow queue of "olive" (her objects) and because
she is a member of the writer's workflow queue she has permission to
change the ownership of her object to "reviewer". At this
point she will not be able to see or change the object. It is
now the reviewer's responsibility to do something with that object.

```json
    {
        "workflow_name": "Reviewer",
        "workflow_id": "reviewer",
        "object_permissions": [ "read" ],
        "assign_to": [
            "editor",
            "reviewer",
            "olive"
        ],
        "view_object_ids": [
            "reviewer",
            "olive"
        ]
    }
```

Our reviewer can either pass the object on to the editor,
leave it in the review queue or pass it back to Olive our
writer.

The editor is allowed to change the object and the editor
is allowed to pass the object back down the workflow. The
editor workflow queue would look like

```json
    {
        "workflow_name": "Editor",
        "workflow_id": "editor",
        "object_permissions": [
            "create",
            "read",
            "update"
        ],
        "assign_to": [  
            "editor",
            "rejected",
            "burried",
            "publisher",
            "reviewer",
            "writer",
            "olive"
        ],
        "view_object_ids": [
            "olive",
            "writer",
            "reviewer",
            "editor",
            "rejected",
            "buried"
        ]
    }
```

And finally the publisher has permissions on everything.

```json
    {
        "workflow_name": "Publisher",
        "workflow_id": "publisher",
        "object_permissions": [
            "create",
            "read",
            "update"
        ],
        "assign_to": [ "*" ],
        "view_object_ids": [ "*" ]
    }
```
