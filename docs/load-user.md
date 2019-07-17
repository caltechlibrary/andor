
# Load User

**AndOr** uses a dataset collection called "user.AndOr" 
when operating. You can load or update this collection
using simple [TOML]() files. The file TOML file(s) can
contain one or more user definitions. Here is an example
of a single user TOML file.

```toml
    # User id
    ["jane.doe@example.edu"]
    # Display name
    display_name = "Jane Doe"
    # By default objects are create in this queue
    create_queue = "deposit"
    # Jane is a member of the "deposit" workflow/queue
    member_of = ["deposit"]
```



