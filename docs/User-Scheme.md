
# User Scheme

While **AndOr** manages users outside the web UI it still 
needs to keep track of basic user information such as the
roles they have access to.

Below is an example JSON document describing the user Jane Doe.

```json
    {
        "user_id": "jane.doe@example.edu",
        "display_name": "Jane Doe",
        "roles": [ "publisher", "editor", "writer" ]
    }
```

The JSON documents hold four pieces of information

user\_id
: (string, required) in this case an email address is a unique string used to map a user to their assigned roles

display\_name
: (string, optional) a display name, a convenience field for us Humans when IDs like an ORCID are less obvious

roles
: (list of strings, defaults to empty list) this lists the roles available to this user. If role specified is "\*" it means the user is a member of all roles, this is useful for repository administrators

Here's an example of expressing that user in TOML.

```toml
    # User id
    ["jane.doe@example.edu"]
    # Display Name
    display_name = "Jane Doe"
    # The queues/roles this user can see.
    roles = [ "publisher", "editor", "writer" ]
```

## How the user object is used

When Jane authenticates with the system she goes from being
"anonymous" to "jane@example.edu" user.  This means she now has the
permissions associated with "publisher", "editor" and "writer" roles.
If the "writer" role allows object creation Jane can create
objects as a "writer". 
Note if Jane doesn't have an have any roles defined she would
have zero permissions to access any objects. Each role list in
"roles" establishes her capabilities to interact with 
objects in the collection(s) of the **AndOr** running.

## anonymous, the default user

There is always a "anonymous" user defined. It is used for anyone who
has not authenticated.  While the "anonymous" user is assumed to always
exist it is not provided with any roles. If you want to allow
anonymous access to a collection (e.g. the general public can see
"published" objects) then you should create a JSON record for the
"anonymous" user and explicitly put them in a single role that
only has read permissions for "published" objects. See [Role Use Cases](Role-Use-Cases.html) for example.

## Picking IDs

You can avoid the toxic storage of secrets by using an external
authentication mechanism (e.g. OAuth 2, Shibboleth) as well as 
limiting the value of your Unique id.  If you create users using 
an email address you do have some value but probably minimal value 
if that same email address is publicly known (e.g. published in 
the institute directory).  This would be similar if you were using
an ORCID as an identifier. ORCID is a nice choice because while
the number is listable at [orcid.org](https://orcid.org) the
ORCID owner controls directly how much information is exposed.
**AndOr** only needs to assert a users' claim they control
the ID it doesn't harvest any data from the ID provider aside
from the ID verified.


