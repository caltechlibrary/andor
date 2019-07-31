
# roles.toml

The "roles.toml" defines three relationships. 

1. Permissions associated with the role
2. The state(s) subject to the role's permissions 
3. The states which a role can assign objects into (e.g. workflow)

Permissions are create, read, update, delete (CRUD). A "state"
is one or more objects with a `._State` value with the same name.
E.g. A "published" state would be objects containing 
`._State` value equal to "published".

Both states and assign to states can be listed explicitly
or the `"*"` value can be used. The latter means all objects
regardless of the `._State` value.

```toml
    #
    # Example "roles.toml". Lines starting with "#" are comments.
    # This file setups up the roles used by AndOr.
    #
    [admin]
    role_name = "Administrator"
    states = [ "*" ]
    create = true
    read = true
    update = true
    delete = true
    assign_to = [ "*" ]
    
    # A depository can create an object in the review state and
    # nothing else. E.g. a one time blind deposit.
    [depositor]
    role_name = "Depositor"
    states = [ "review" ]
    create = true
    read = false
    update = false
    delete = false
    assign_to = [ ]
    
    # A reviewer can pub a published item in review, 
    # In "published" and in embargoed.
    [reviewer]
    role_name = "Reviewer"
    states = [ "review" ]
    create = false
    read = true
    update = true
    delete = true
    assign_to = [ "review", "published", "embargoed" ]
```
