
# Role Use Cases

Below are use cases exploring how a user/role/queue model.

## Use case 1

We would like people (e.g. Jane) to curate the collection.
To curate the whole collection Jane needs to have the
following permissions-- create, read, update, delete,
change queues.  We can define a role that applies to
all objects in all queues and assign Jane to it.

```json
    {
        "role_name": "Curator",
        "role_id": "curator",
        "queues": [ "*" ],
        "create": true,
        "read" : true,
        "update" : true,
        "delete" : true,
        "assign_to": [ "*" ],
    }
```

Notice that we use "\*" twice. First we specify the 
queues value include "\*". This special, it means this
role applies to ALL objects holding any `._Queue`
value. The second place of "\*" is in the list of
`.assign_to`. Again this means a curator can assign objects
to any queue.

Jane needs to be a curator. We add "curator" to the
`.roles` list in her user object.
Jane's email address is "jane@example.edu" and what
we've used as a user id.

```json
    {
        "user_id": "jane@example.edu",
        "display_name": "Jane",
        "roles": [ "curator" ]
    }
```

When Jane authenticates with the system she goes from being
"anonymous" to "jane@example.edu" user.  This means she now has the
role of a curator. The role "curator" gives Jane
all rights of creation, read, update, delete and assignment on
any object in the collection. 


## Use case 2

We would like the public to be able to view the "published" contents
of the of our collection. We can do this by creating a role queue 
"public" and given read access in the "published" queue. We also
want to create a user object for "anonymous" (un-authenticated users)
and make "anonymous" a member of "public".

The role queue object would look like

```json
    {
        "role_name": "Public",
        "role_id": "public",
        "queue": [ "published" ],
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
        "roles": [ "public" ]
    }
```

Combining this with use case 1, we would have a single user
named Jane who can create, manange and  publish objects 
when she assignes the objects to the "published" queue.

## Use case 3

We want to allow anonymous users to "deposit" objects.  We can
create a role called "depositor" with create permissions
in a queue called "review".  If we give "anonymous" 
"deposit" membership anyone can deposit objects. The "review"
queue would function like an inbox.  With only create permission
they would not be able to see any objects except those they would
see if they have a role of "public". The review queue objects would
remain invisible to the anonymous user(s) until the objects were
moved into the "published" queue by a curator (e.g. Jane).

```json
    {
        "role_name": "Depositor",
        "role_id": "depositor",
        "queues": [ "review" ],
        "create": true,
        "read": false,
        "update": false, 
        "delete": false,
        "assign_to": [ ]
    }
```

Now if we update our "anonymous" user we can add the 
the "depositor" role queue. Any objects created by
"anonymous" would be created in "review". 

```json
    {
        "userid": "anonymous",
        "display_name": "Public",
        "roles": [ "public", "depositor" ]
    }
```

## Use case 4

Exploring the relationship between roles and queues.
We would like our objects to travel between the following 
states - review, published, embargoed.

We'd like objects to be created in review queue only.
Any authenticated user should be able to create objects. 
A reviewer's should be able to perform the editorial function
of moving objects to "published" and "embargoed" queues.
Our publisher should be able to move objects anywhere.

Here are some of our policies we want to enforce.

1. Allow any authenticated user to deposit 
    + i.e remove the "depositor" role from "anonymous" user
    + All users need to have the "depositor" role explicitly.
2. Allow reviewers to read and delete deposits but not create or update them. 
    + i.e. remove the "create" and "update" permissions from the "reviewer" role 
3. Allow reviewers to assign objects to "published" and "embargoed" queues
4. Allow curators need all the permissions of the review plus the "update" permission.
5. "Publisher" be able to create objects in "review" queue and read, update, delete objects in any queue. They should be able to move objects into any queue.
    + A "publisher" wouldn't be a single role but a composit of "depositor" and "currator".

Innez is the publisher, Jane is a curator, Millie is a reviewer and
Bea is a depositor.

Here's the steps to implement a solution started with previous use case.

1. Remove "depositor" from "anonymous" roles
2. Add/Update our users
    a. Assign Innez the roles of "depositor" and "currator"
    b. Assign Jane the roles of "depositor" and "currator"
    c. Assign Millie the roles of "depositor" and "reviewer"
    d. Assign Bea the role of "depositor"
3. Use our previous "depositor" role
4. Update our "reviewer" role with only read and delete permissions
5. Update our "curator" role to explicitly list queues and assignments, remove create permission


Step 1. our "anonymous" user now should look like

```json
    {
        "user_id": "anonymous",
        "display_name": "",
        "roles": [ "public" ]
    }
```

Let's create/update user objects for Innez, Jane, Millie and Bea.

```json
    {
        "user_id": "innez",
        "display_name": "Innez",
        "roles": [ "depositor", "currator" ]
    }
```

```json
    {
        "user_id": "jane",
        "display_name": "Jane",
        "roles": [ "depositor", "currator" ]
    }
```

```json
    {
        "user_id": "millie",
        "display_name": "Millie",
        "roles": [ "depositor", "reviewer" ]
    }
```

```json
    {
        "user_id": "bea",
        "display_name": "Bea",
        "roles": [ "depositor" ]
    }
```

Our curator role should look like

```json
    {
        "role_id": "curator",
        "role_name": "Curator",
        "queues": [ "review", "embargoed", "published" ],
        "create": false,
        "read": true,
        "update": true,
        "delete": true,
        "assign_to": [ "*" ]
    }
```

Our depositor role looks like

```json
    {
        "role_id": "depositor",
        "role_Name": "Depositor",
        "queues": [ "review" ],
        "create": true,
        "read": false,
        "update": false,
        "delete": false,
        "assign_to": [ ]
    }
```

And our reviewer would look like


```json
    {
        "role_name": "Reviewer",
        "role_id": "reviewer",
        "queues": [ "review" ],
        "create": false,
        "read": true,
        "update": false,
        "delete": true,
        "assign_to": [ "review", "published", "embargoed" ]
    }
```



