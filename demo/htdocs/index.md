+++
markup = "mmark"
+++


# Welcome to [And/Or](https://github.com/caltechlibrary/andor) Demos

This is a demo of And/Or a web base multi-user version of 
[dataset](https://github.com/caltechlibrary/dataset). It is a web
service designed to build curation tools for metadata that doesn't
currently fit in our existing repository and catalog systems.

The demos shows basic curation operations like listing keys,
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
defined a set of configuration files.

## The demos

[people](/people/)
: demonstrates a simple person object collection. People objects consist of a 
given name, family name and a series of unique ids (e.g. orcid, snac, wikidata, 
viaf). It is based on the Caltech Library People Identity project currently
maintained in a Google Sheet.



