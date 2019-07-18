
# Queues Scheme

A Queue is a list of objects associated with a 
given state. Queues have names and a retrieved
by their names. Queues also are associated with
workflows what define the actions available in
the queue by users associated with those workflows.

When you "add" an object to a queue you are changing
the internal `._Queue` value in the object after 
verifying that a given user and the workflows allow
the addition.

The queues are created when **AndOr** is started
and it loads "workflows.toml". The queue lasts
as long as **AndOr** remains running. If an object
has a `._Queue` value that is no longer defined in
any workflow the object will be invisible to the
**AndOr** web API. If that happens you well need
to write a script or use the **dataset**[^1] command
to update the `._Queue` value to something 
specified in a workflow.


[^1]: dataset is a command line workflow for managing JSON objects in collections, see https://github.com/caltechlibrary/dataset


