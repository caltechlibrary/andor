
# Loading Workflow(s)

Workflows are a central organzation of **AndOr** for
access (permissioning) as well as organizating curitorial 
into work queues.  The `load-workflow` verb will 
create/update workflows available to **AndOr**. A
workflow is defined in [TOML]().  A toml file can
contain one workflow or many. Below are three workflows
"admin", "deposit", "published". The workflow
names correspond to work queues of "admin", "deposit",
"published".

```toml
    # This defines the workflow "admin"
    # It sets workflow id is "admin"
    [admin]
    # The display name is "Admin"
    workflow_name = "Admin"
    # Object level permisions in the "admin" queue
    object_permissions = [ "create", "read", "update" ]
    # The available objects can be assigned to
    assign_to = ["*"]
    # The workflows where our object permissions apply.
    queues = ["*"]

    # This defines our "deposit" workflow
    [deposit]
    # The display name is "Deposit"
    workflow_name = "Deposit"
    # Object level permissions in the "deposit" queue
    object_permissions = [ "create", "read", "update" ]
    # The available objects can be assigned to
    assign_to = [ "deposit", "admin" ]
    # The workflows where our object permissions apply.
    queues = [ "deposit" ]

    # This defines our "published" workflow, it's terminal
    # So both object permissions and assign_to are restrictive.
    [published]
    # The display name is "Published"
    workflow_name = "Published"
    # Object level permissions in the "published" queue
    object_permissions = [ "read" ]
    # The available objects can be assigned to
    assign_to = [ ]
    # The workflows where our object permissions apply.
    queues = [ "published" ]
```

The admin user would have access to everything.  Our "deposit" 
workflow, in this example only can create, read and update 
objects in the "deposit" state but they can assign
the object to the "admin" queue to decide if they should be
published or not. The users who are members of "publish"
(e.g. an anonymous user if you defined that) only can read
objects. They can't assign objects to a different workflow,
they can change or create objects.


## Loading our workflow file

If our workflow file was named "workflows.toml" we'd run
the following command to setup/update the workflows in **AndOr**

```bash
    bin/AndOr load-workflow workflows.toml
```

The file's content will add/update workflows in the 
[dataset](https://caltechlibrary.github.io/dataset) collection
named "workflows.AndOr". NOTE: Loading a file doesn't
remove previously defined workflows! It only adds/updates.

