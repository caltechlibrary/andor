
# andor.toml

The "andor.toml" file is responsible for configuring the 
**AndOr** web service and specifiying the location of 
"users.toml", "roles.toml" and any dataset collections
being serviced by the service.

The "andor.toml" gets generated if none exist in when you
run the `AndOr init` command. Here's an example.

```toml
    #
    # Example "andor.toml" 
    #
    # Lines starting with "#" are comments.
    # This file configuration the AndOr web service.
    #
    roles_toml = "roles.toml"
    users_toml = "users.toml"
    repositories = [ "repository.ds" ]
    protocol = "http"
    port = "8246"
    host = "localhost"
```

