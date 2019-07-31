+++
markup = "mmark"
+++

# Role Scheme

A role encapsulates three things. Permissions,
object states[^1] and state assignments. A user
has one or more roles the roles give them permissions
to act on objecets in the states assocaited with the role.
It also defines what states objects can be assigned into.
This is what gives us workflows.  Objects state assignment
is independent of the permissions.  In this way a role
might define creation of an object but also the assignment
to a new state that is not covered by the role. E.g. A
writer might create an object in draft state, then when ready
assign it to "review" state where an editor continues working
on it. 

Each role specifies the operations that can be
performed on a list of states. These include basic
CRUD[^2] operations on all objects in the state.
It also specifies the assignment (handing off) of
objects to other states.

An addition state exists without definition. That queue
is "deleted". Like EPrints[^3] **And/Or** doesn't delete
objects from disc. They continue to exist in a trashbin
like state.  They could be truely deleted by writing
a garbage collection script using dataset.

Here's an example object  defining a "depositor" role
associated with a "review" state. A user working in 
the "depositor" role can create an objects but can 
not read them.  When an object created it s assigned 
the `._State` value "review". So our depository user can 
create objects in the "review" state but they have no 
other rights beyond that.

```json
    {
        "role_name": "Depositor",
        "role_id": "deposit",
        "states": [ "review" ],
        "create": true,
        "read": false,
        "update": false,
        "delete": false,
        "assign_to": [ ]
    }

```

Next is an example of a "reviewer" role. A
reviewer can read, update, and delete objects 
but not create them. They can also move the 
objects into review, published, and embargoed
states. Once an object is "published" they would
not be able to update or re-assign the object.

```json
    {
        "role_name": "Reviewer",
        "role_id": "reviewer"
        "states": [ "review", "embargoed" ],
        "create": false,
        "read" : true,
        "update": true,
        "delete": true,
        "assign_to": [ "review", "embargoed", "published" ]
    }
```

Finally our "publisher" role would have the ability
to perform all operations on all states.  Note how
we specifify "states" and "assign\_to" values using
the name "*". "*" is a wild card meaning all states.

```json
    {
        "role_name": "Publisher",
        "role_id": "publisher"
        "states": [ "*" ],
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

states
: (a list of string, required, may be empty list) the associated state where the CRUD permissions apply

create
: (bool, defaults to false if not defined) the ability to create an object in the state(s)

read
: (bool, defaults to false if not defined) the ability to read objects in the state(s)

update
: (bool, defaults to false if not defined) the ability to update (edit) objects in the state(s)

delete
: (bool, defaults to false if not defined) the ability to "delete" objects
(really move to a trashbin state) in the state(s)

assign\_to
: (list of string, may be empty) the list of states this role may assign objects into



[^1]: objects have a field called ".\_State" that holds the name of the role queue they are currently associated with

[^2]: CRUD, refer to "create", "read", "update", "delete" operations on a object

[^3]: EPrints is an excellent Open Source Repository system from South Hampton University

