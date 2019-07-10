
# AndOr

> Another digital Object repository

This is a proof of concept for very light weight digital object
repository. It is built on [dataset](https://caltechlibrary.github.io/dataset) 
collections, simple JSON oriented web API and UI created via
HTML 5 and JavaScript held in an htdocs directory.  It is an exercise 
exploring implementing how to implement a useful digital object
repository with the fewest features.

## Goals

+ Simpler the EPrints and Invenio
+ Simpler than Drupal
+ No RDMS
+ Agnostic about the shape of objects stored as long as it is valid JSON objects
+ Supports versionable attached media
+ Support continuous migration


## Assumptions

+ dataset collections are sufficient to hold metadata and media
+ small number of curratorial users larger number of readers and depositors
+ authentication is external (e.g. PAM, OAuth2)
+ workflows are a requirement
+ object scheme can be enforced before sending objects to the API for storage
+ An external mechanism (e.g. Solr, Elastic Search, Lunr) can be used to query objects
+ UI can be solely implemented as a web app using HTML 5 and JavaScript 
    + This means you could also sit Andor behind Drupal easily

## Limitting features and complexity

One of the most complicated parts of a digital object repository
is enforcing and managing users, workflows, permissions and storage
scheme.  For focus on using a workflow oriented permission scheme
but does not provide a web user interface to setup or maintain it.
Instead managing users and workflows is a matter of configuration.
Configuration is stored in the dataset collection's root folder
and can be either JSON or TOML documents.

Authenticated users return their authenticated user id, e.g. 
(e.g. email, ORCID) this is how me map a user into a workflow.

Unauthenticated users are treated as the "anonymous" user and
are restricted by workflows available for that user. 

A user's membership in workflows defines their permissions. 

Because **AndOr** is built on dataset all objects must have unique ids. 
Those ids are supplied (or calculated) client side.  Objects may include 
attached documents which can be versioned automatically. 

## Use cases

Next are some use cases exploring this interaction.

### Use case 1

We would like people (e.g. Jane) to currate the collection.
To currate the whole collection Jane needs to have the
following permissions-- create, read, update, delete,
change groups, and to list all objects
in the collection. To do that we can create a group called 
"currators".

```json
    {
        "group_name": "Currators",
        "group_id": "currators"
        "object_permissions": [
            "create",
            "read",
            "update",
            "delete"
        ]
        "change_groups_to": [ "*" ],
        "list_objects_in_group": [ "*" ]
    }
```

Notice that we have lists for the change group, change
owner and list objects in group permission. The list contains
an asterisks. An will match any group. The alternative would
be to list specific groups, owners.

Jane needs to be a currator. We need to create a user record for her
and to make her a member of the "currators" group.  Jane's email address 
is "jane@example.edu" so we use that as her user id. Jane's user 
object would look like.

'''json
    {
        "user_id": "jane@example.edu",
        "display_name": "Jane",
        "create_objects_as": "currators",
        "member_of": [ "currators" ]
    }
'''

When Jane authenticates with the system she goes from being 
"anonymous" to "jane@example.edu" user.  This means she now has the
permissions of a currator. If she creates an object the object will 
be assigned to the "currators" group. Since
the currators group has change group permissions for any group of objects
in the collection can take any action with any object as needed.


### Use case 2

We would like the public to be able to view the "published" contents 
of the of our collection. We can do this by creating a group called 
"published". The published group will be given read access to the
objects. It'll be allowed to list objects in the "published" group.

The group object would look like

```json
    {
        "group_name": "Published",
        "group_id": "published"
        "collection_permissions": [
            "read"
        ]
        "change_groups_to": [],
        "list_objects_in_group": [ "published ]
    }
```

Any member of "published" will be able to read and objects
associated with "published" group.

Now we need to explicitly associated "anonymous"
with the "published" group.

```json
    {
        "userid": "anonymous",
        "display_name": "Public",
        "create_objects_as": null,
        "member_of": [ "published" ]
    }
```

When anonymous goes to access an object and the group associated with
the object is "published" then anonymous will be able to read the object.
Likewise anonymous is not a member of any other groups so they will
have no permissions to access them.

### Use case 3

We want to allow anonymous users to "deposit" objects.  We can
create a group called "deposit". It should only have create
permissions.  If we add "anonymous" to the "deposit" group
we could then allow the public to submit records but nothing else.
If anonymous is a member of "published" they will be able to still
read all published records.

```json
    {
        "group_name": "Depositor",
        "group_id": "deposit",
        "object_permissions": [ "create" ]
        "change_groups_to": [],
        "list_objects_in_group": []
    }
```.

Now if we update our "anonymous" user we can add them to 
the "deposit" group and associate the created object with
a group called "deposit".

```json
    {
        "userid": "anonymous",
        "display_name": "Public",
        "create_objects_as": "deposit",
        "member_of": [ "published", "deposit" ]
    }
```

Because the object created will have the group "deposit" and
our "currators" group has permissions to list all objects groups
we can treat the "deposit" group a an inbox to be processed.
If the currators approved the deposit they can change the objects'
group to "published".

### Use case 4

Creating workflows with groups. We would like our objects to travel
through the following states - deposit, review, then either be
flagged with publish, embargo, and needs curration.

Here are some of our policies we want to enforce.

1. Allow any authenticated user to deposit
2. Allow reviewers to choose between the following status -- deposit, review, published, embargo, rejected and needs further curration
3. If "needs curration" is chosen then the reviewer should no longer be able to see the object
4. The depositor should not see the object after it is "deposited"
5. Reviewers should not be able to change an object only change the group
6. Reviewers can't create new objects

Jane is a currator. Millie is a reviewer. Millie should not be able
to update the objects but she should be able to list objects that
have been deposited and change objects and group on an
object to As 

Jane is already in the currators group previously defined. We need
to define a reviewer group. Millie will need to be created and be
a member of "reviewer".

```json
        "group_name": "Reviewer",
        "group_id": "review",
        "object_permissions": [ "read" ]
        "change_groups_to": [ 
            "deposit", 
            "review", 
            "published", 
            "embargo",  
            "rejected",
            "currators"
            ],
        "list_objects_in_group": [
            "deposit",
            "review",
            "published",
            "embargo",
            "rejected"
        ]
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

If later we deside Millie should be able to create objects then
we can add her the deposit group. She would not be able to 
edit her deposits but she could change the group value and list it.


### Use case 5

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
deposits but not someone elses. We can do this by
create an "olive" group which and having that as
the default group when she creates a new object.
We can also associate olive with the writer group
which can change permissions to reviewer. 

Here is Olive's group

```json
    {
        "group_name": "Olive",
        "group_id": "olive",
        "object_permissions": [ 
            "create", 
            "read", 
            "update", 
            "delete" 
        ],
        "change_groups_to": [ ],
        "list_objects_in_group": [
            "olive"
        ]
    }
```

Here is the general writer's group object

```json
    {
        "group_name": "Writer",
        "group_id": "writer",
        "object_permissions": [],
        "change_groups_to": [ 
            "reviewer"
        ],
        "list_objects_in_group": []
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
any object with a group of "olive" (her objects) and because
she is a member of the writer's group she has permission to 
change the ownership of her object to "reviewer". At this
point she will not be able to see or change the object. It is
now the reviewer's responsibility to do something with that object.

```json
    {
        "group_name": "Reviewer",
        "group_id": "reviewer",
        "object_permissions": [ read ],
        "change_groups_to": [ 
            "editor",
            "reviewer",
            "olive"
        ],
        "list_objects_in_group": [
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
editor group would look like

```json
    {
        "group_name": "Editor",
        "group_id": "editor",
        "object_permissions": [ 
            "create",
            "read", 
            "update", 
            "delete" 
        ],
        "change_groups_to": [  
            "editor",
            "rejected",
            "burried",
            "publisher",
            "reviewer", 
            "writer", 
            "olive",
        ],
        "list_objects_in_group": [
            "olive",
            "writer",
            "reviewer",
            "editor",
            "rejected",
            "burried",
        ]
    }
```

And finally the publisher has permissions on everything.

```json
    {
        "group_name": "Publisher",
        "group_id": "publisher",
        "object_permissions": [
            "create",
            "read",
            "update",
            "delete",
        ],
        "change_groups_to": [ "*" ],
        "list_objects_in_group": [ "*" ]
    }
```

