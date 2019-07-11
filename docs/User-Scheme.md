
# User Scheme

While **AndOr** defers for managin users via a web UI and from
storing user credentials it does still need to track some based
user information.

Below is an example JSON document describing the user Jane Doe.


'''json
    {
        "user_id": "jane.doe@example.edu",
        "display_name": "Jane Doe",
        "create_objects_as": "writer",
        "member_of": [ "publisher", "editor", "writer" ]
    }
'''

The JSON documents hold four pieces of information

user_id
: (string, required) in this case an email address is a unique string used to map a user to their assigned workflows

display_name
: (string, optional) a display name, a conveience field for us Humans when IDs like an ORCID are less obvious

create_objects_as
: (string, optional, defaults to null) if not null assigns the workflow value when an object is created (NOTE: workflow needs to allow creating objects)

member_of
: (list of strings, defaults to empty list) this lists the workflows available to this user. If workflow specified is "\*" it means the user is a member of all workflows, their for has all defined permissions


## How the user object is used

When Jane authenticates with the system she goes from being 
"anonymous" to "jane@example.edu" user.  This means she now has the
permissions associated with "publisher", "editor" and "writer" workflows.
If Jane creates a new object it will be created in the "writer" workflow.
Note if Jane doesn't have an have any workflows defined she would
have zero permissions to access any objects. Each workflow list in
"member\_of" establishes the permissions under which she can see objects.

## anonymous, the default user

There is always a "anonymous" user defined. It is used for anyone who
has not authenticated.  While the "anonymous" user is assumed to always
exist it is not provided with any workflows. If you want to allow
anonymous access to a collection (e.g. the general public can see
"published" objects) then you should create a JSON record for the
"anonymous" user and explicitly put them in a single workflow that
only has read permissions for "published" objects. See [Workflow Use Cases](Workflow-Use-Cases.html) for example.

## Picking IDs

You can avoid the toxic storage of secrets by using an external
authentication mechasism (e.g. OAuth 2) as well as limitting the
value of your Unique id.  If you create users using an email address
you do have some value but probably minimal value if that same
email address is publically known (e.g. published in the 
institute directory).  This would be similar if you were using
an ORCID as an identifier. ORCID is a nice choice because while
the number is listable at [orcid.org](https://orcid.org) the
ORCID owner controls directly how much information is exposed.
**AndOr** only needs to assert a users' claim they control
the ID it doesn't harvest any data from the ID provider.


