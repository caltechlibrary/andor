
# roles.toml

The "roles.toml" defines three relationships. 

1. Permissions associated with the role
2. The queue(s) subject to the role's permissions 
3. The queues which a role can assign objects into (e.g. workflow)

Permissions are create, read, update, delete (CRUD). A "queue"
is one or more objects with a `._Queue` value with the same name.
E.g. A "published" queue would be objects containing 
`._Queue` value equal to "published".

Both queues and assign to queues can be listed explicitly
or the `"*"` value can be used. The latter means all objects
regardless of the `._Queue` value.

```toml
    #
    # Example "roles.toml". Lines starting with "#" are comments.
    # This file setups up the roles used by AndOr.
    #
    [admin]
    role_name = "Administrator"
    queues = [ "*" ]
    create = true
    read = true
    update = true
    delete = true
    assign_to = [ "*" ]
    
    # A depository can create an object in the review queue and
    # nothing else. E.g. a one time blind deposit.
    [depositor]
    role_name = "Depositor"
    queues = [ "review" ]
    create = true
    read = false
    update = false
    delete = false
    assign_to = [ ]
    
    # A reviewer can pub a published item in review, 
    # In "published" and in embargoed.
    [reviewer]
    role_name = "Reviewer"
    queues = [ "review" ]
    create = false
    read = true
    update = true
    delete = true
    assign_to = [ "review", "published", "embargoed" ]
```
