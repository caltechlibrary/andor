
# workflows.toml

The "workflows.toml" defines two relation ships. First is the
capabilities of users assigned to the workflow. The second
is the __queue__ of objects the workflow can workflow. A "queue"
represents an objects' state. Below is a set of example workflows
define queues of "review" and "publish" as well as the workflows
of "admin", "depositor", "reviewer", and "published".

```toml
    #
    # Example "workflows.toml". Lines starting with "#" are comments.
    # This file setups up the workflows used by AndOr.
    #
    [admin]
    workflow_name = "Administrator"
    queue = "*"
    create = true
    read = true
    update = true
    delete = true
    assign_to = [ "*" ]
    
    [depositor]
    workflow_name = "Depositor"
    queue = "review"
    create = true
    read = false
    update = false
    delete = false
    assign_to = [ ]
    
    [reviewer]
    workflow_name = "Reviewer"
    queue = "review"
    create = false
    read = true
    update = true
    delete = true
    assign_to = [ "published" ]
    
    [published]
    workflow_name = "Published"
    queue = "published"
    create = false
    read = true
    update = false
    delete = false
    assign_to = [ ]
```
