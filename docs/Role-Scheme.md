+++
markup = "mmark"
+++

# Role Scheme

A role encapsulates three things. Permissions,
queues[^1] and queue assignments (workflows). A user
has one or more roles the roles given the permissions
for objects in the applicable queues and allow the
objects to be reassigned to other queues (which are
independent of the permissions).  A set of roles 
describe how objects can move through the repository 
systems.

Each role specifies the operations that can be
performed on a list of queue. These include basic
CRUD[^2] operations on all objects in the queues.
It also specifies the assignment (handing off) of
objects to other queues.

An addition queue exists without definition. That queue
is "deleted". Like EPrints[^3] **And/Or** doesn't delete
objects from disc. They continue to exist in a trashbin
like state.  They can be truely deleted by writing
a trivial garbage collection script using dataset.

Here's an example object of defining a "deposit" role
associated with a "deposit" queue. A user working in 
the "deposit" role create objects but can not read them. 
When an object created it s assigned the `._Queue`
value "review". So our depository user can create objects 
in the "review" queue but they have no other rights 
beyond that.

```json
    {
        "role_name": "Depositor",
        "role_id": "deposit",
        "queues": [ "review" ],
        "create": true,
        "read": false,
        "update": false,
        "delete": false,
        "assign_to": [ ]
    }

```

Next is an example of a "review" role. A
reviewer can read, update, and delete objects 
but not create them. They can also move the 
objects into review, published, and embargoed
queues. Once an object is "published" they would
not be able to update or re-assign the object.

```json
    {
        "role_name": "Review",
        "role_id": "review"
        "queues": [ "review", "embargoed" ],
        "create": false,
        "read" : true,
        "update": true,
        "delete": true,
        "assign_to": [ "review", "embargoed", "published" ]
    }
```

Finally our "publisher" role would have the ability
to perform all operations on all queues.  Note how
we specifify "queues" and "assign\_to" values.

```json
    {
        "role_name": "Publisher",
        "role_id": "publisher"
        "queues": [ "*" ],
        "create": true,
        "read" : true,
        "update": true,
        "delete": true,
        "assign_to": [ "*" ]
    }
```

What if we want to pull an object back from published state?
The publisher has the right to assign the object to any state
such as "embargoed" and "review".

The fields

role\_name
: (string, optional) the human readable name for the role

role\_id
: (string, required, must be unique) the id for this role

queues
: (a list of string, required, may be empty list) the associated queues wher the CRUD permissions apply

create
: (bool, defaults to false if not defined) the ability to create an object in the queue(s)

read
: (bool, defaults to false if not defined) the ability to read objects in the queue(s)

update
: (bool, defaults to false if not defined) the ability to update (edit) objects in the queue(s)

delete
: (bool, defaults to false if not defined) the ability to "delete" objects (really move to a trashbin state) in the queue(s)

assign\_to
: (list of string, may be empty) the list of queues this role may assign objects into



[^1]: objects have a field called ".\_Queue" that holds the name of the role queue they are currently associated with

[^2]: CRUD, refer to "create", "read", "update", "delete" operations on a object

[^3]: EPrints is an excellent Open Source Repository system from South Hampton University

