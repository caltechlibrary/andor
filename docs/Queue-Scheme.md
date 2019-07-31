
# Queues

A Queue is a list of objects associated with a 
given state. Queues have names and are retrieved
by their names. Queues also are associated with
roles that define the actions available in
the queue by users associated with those roles.

When you "add" an object to a queue you are changing
the internal `._State` value in the object after 
verifying that a given user and the roles allow
the addition.

The queues are created when **AndOr** is started
and it loads "roles.toml". The queue lasts
as long as **AndOr** remains running. If an object
has a `._State` value that is no longer defined in
any role the object will be invisible to the
**AndOr** web API. If that happens you well need
to write a script or use the **dataset**[^1] command
to update the `._State` value to something 
specified in a role.


[^1]: dataset is a command line role for managing JSON objects in collections, see https://github.com/caltechlibrary/dataset


