
# Workflow Use Cases

Below are use cases exploring how a user/workflow/queue model
would allow or prevent actions on objects in a collection.

## Use case 1

We would like people (e.g. Jane) to curate the collection.
To curate the whole collection Jane needs to have the
following permissions-- create, read, update, delete,
change workflow queues.  We can define a workflow to work
apply to all objects and and make Jane a member of that workflow.

```json
    {
        "workflow_name": "Curators",
        "workflow_id": "curators",
        "queue": "*",
        "create": true,
        "read" : true,
        "update" : true,
        "delete" : true,
        "assign_to": [ "*" ],
    }
```

Notice that we use "\*" twice. First we specify the 
queue value as "\*". This special, it means this
workflow applies to ALL objects holding any `._Queue`
value. The second place of "\*" is in the list of
`.assign_to`. Again this means a curator can assign objects
in this workflow to ANY workflow.

Jane needs to be a curator. We add "curator" to the
`.member_of` list in her user object.
Jane's email address is "jane@example.edu" and what
we've used as a user id.

```json
    {
        "user_id": "jane@example.edu",
        "display_name": "Jane",
        "member_of": [ "curators" ]
    }
```

When Jane authenticates with the system she goes from being
"anonymous" to "jane@example.edu" user.  This means she now has the
permissions of a curator. The workflow "curator" gives Jane
all rights of creation, read, update, delete and assignment on
any object in the collection(s). This workflow is generally
too permissive outside private collections.


## Use case 2

We would like the public to be able to view the "published" contents
of the of our collection. We can do this by creating a workflow queue 
"public" and given read access in the "published" queue. We also
want to create a user object for "anonymous" (un-authenticated users)
and make "anonymous" a member of "public".

The workflow queue object would look like

```json
    {
        "workflow_name": "Public",
        "workflow_id": "public",
        "queue": "published",
        "create": false,
        "read": true,
        "update": false,
        "delete": false,
        "assign_to": [ ]
    }
```

The "anonymous" user object would look like

```json
    {
        "user_id": "anonymous",
        "display_name": "",
        "member_of": [ "public" ]
    }
```

Combining this with use case 1, we would have a single user
named Jane who can publish objects by assigning objects to
"public".

## Use case 3

We want to allow anonymous users to "deposit" objects.  We can
create a workflow queue called "deposit". It should only have create
permissions for the "review" queue.  If we give "anonymous" 
"deposit" membership they can create objects that will be added
to the "review" queue. They will not be in the "published" queue
to until someone intervines they will be invisible to the public.
Remember since anonymous is a member of "public" they will be 
able to still read all published records.

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

Now if we update our "anonymous" user we can add them to 
the "deposit" workflow queue and associate the created object with
a workflow queue called "deposit".

```json
    {
        "userid": "anonymous",
        "display_name": "Public",
        "member_of": [ "public", "deposit" ]
    }
```

Because the object created will have the workflow queue "review" and
our "curators" workflow queue has permissions to list all objects workflow 
queues we can treat the "deposit" workflow queue as an inbox to be 
processed.  If the curators approved the deposit they can change the 
objects' workflow queue to "published".

## Use case 4

Creating workflows with queues. We would like our objects to travel
between the following states - review, editorial, published, embargoed.
Everything starts in review, our reviewers can move the object onto
either editorial, published or embargoed states.

Here are some of our policies we want to enforce.

1. Allow any authenticated user to deposit 
    + remove the "deposit" workflow from "anonymous" user
2. Allow reviewers to read and delete deposits but not create or update them. 
3. Allow reviewers to assign objects to editorial, published and embargoed queues
4. Currators have full access to all objects in all queues

Jane is a curator. Millie is a reviewer. Anne is a user who we want
to be able to "deposit" objects but she has no other responsibilities.

Here's the steps to implement a solution

1. Remove "anonymous" membership in "depost"
2. Add user objects for Mille and Anne
3. Create a workflow called "depositor" with can create objects in the "review" queue
3. Create a workflow called "reviewer" that operates on the "review" queue and given that workflow read, delete permission for objects in the queue and is allowed to assign objects to editorial, published and embargoed queues
4. Create three more workflows to trigger the generation of our remaining

Step 1. our "anonymous" user now should look like

```json
    {
        "user_id": "anonymous",
        "display_name": "",
        "member_of": [ "public" ]
    }
```

Let's create user objects for Millie and Anne.

```json
    {
        "user_id": "millie",
        "display_name": "Millie",
        "member_of": [ "reviewer" ]
    }
```

```json
    {
        "user_id": "anne",
        "display_name": "Anne",
        "member_of": [ "depositor" ]
    }
```

Our depositor workflow looks like

```json
    {
        "workflow_id": "depositor",
        "workflow_Name": "Depositor",
        "queue": "review",
        "create": true,
        "read": false,
        "update": false,
        "delete": false,
        "asign_to": [ ]
    }
```

And our reviewer would look like


```json
    {
        "workflow_name": "Reviewer",
        "workflow_id": "reviewer",
        "queue": "review",
        "create": false,
        "read": true,
        "update": false,
        "delete": true,
        "assign_to": [ "review", "editorial", "published", "embargoed" ]
    }
```

If later we decide Millie should be able to create objects then
we can add her the deposit workflow queue. She would have the 
cummulative rights of both the "deposit" and "reviewer" workflows.


