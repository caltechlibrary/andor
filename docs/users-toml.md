
# users.toml

The "users.toml" file specifies users and roles who can access
the **AndOr** web service. Notice it DOES NOT specify authentication.
Authentication is handled at the web server level via an 
external mechanism like OAuth 2, Shibboleth or in the development
content BasicAUTH or Digest AUTH.

The "users.toml" file needs to have a set of entries for each 
user. In the example below the user "rsdoiel" starts off with
a `["rsdoiel"]` heading. The "member\_of" field is very important.
This field designates a list of workflows available to the user.
Workflows control access and capabilities in the **AndOr**
web service.

```toml
    #
    # Example "users.toml". Lines starting with "#" are comments.
    # This shows example users.
    #
    ["rsdoiel"]
    display_name = "R. S. Doiel""
    member_of = [ "admin" ]
```
