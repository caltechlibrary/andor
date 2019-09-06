+++
markup = "mmark"
+++


# Welcome to [And/Or](https://github.com/caltechlibrary/andor) Demos

This is a demo of And/Or. And/Or is a web based multi-user version of 
[dataset](https://github.com/caltechlibrary/dataset). It is a web
service designed to build curation tools for metadata that doesn't
currently fit in our existing repository and catalog systems.

The demo shows basic curation operations like listing keys,
and creating, reading, updating and deleting JSON documents 
stored in dataset collections.

And/Or supports a role/object state permission scheme. This
allows us to create simple workflows by assigning users to roles
which define a set of create, read, update, delete permissions for
a set of object states as well as a list of states the role can
assign objects into.

The user interface of And/Or accesseble dataset collections is
created via HTML, CSS and JavaScript stored statically on disc. 
And/Or itself only provides a static file web service and a JSON
API that maps supported dataset commands to URLs. The JSON
API is also responsible for enforcing the permission policies
defined configuration files.

The login method for the demo uses BasicAUTH. BasicAUTH doesn't
support logout.  If you use private browser tabs then you can close 
the tab, own a new private tab and try the next user account to see 
how roles affects what the user can do.  The demo uses a set of 
usernames - __esther__, __jane__ and __millie__.  Each 
username has the same password, **hello**.

## The demo

[people](/people/)
: demonstrates a simple person object collection. People objects consist of a 
given name, family name and a series of unique ids (e.g. orcid, snac, wikidata, 
viaf). It is based on the Caltech Library People Identity project currently
maintained in a Google Sheet.


