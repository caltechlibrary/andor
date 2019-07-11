
# AndOr

> <span class="red">An</span>other <span class="red">d</span>igital <span class="red">O</span>bject <span class="red">r</span>epository

This is a concept document for a very light weight digital object
repository. If prototoyped it could serve as an interim repository 
for some of our simple EPrints deployments. **Andor** would be 
built from [dataset](https://caltechlibrary.github.io/dataset) 
collections, configuration files, a simple JSON oriented web API 
and UI created via HTML 5 and JavaScript served from a static 
htdocs directory.  It would consist of to pieces of software,
a web server (e.g. NginX, Apache) providing integration options
for authentication as well as search engine support and the
**AndOr** program which would be use to create a web service
to respond to the static content and UI.

## Goals

+ Provide an __interum option__ for simple EPrints repositories
+ Thin stack (e.g. No RDMS requirement, only AndOr and a web server)
    + Be simpler than EPrints, Invenio, Drupal/Islandora
    + Be easier than migrating our EPrints repositories to the cloud
+ Support existing schema
+ Support versioned attached media files
+ Support continuous migration
+ Support multiple front ends (e.g. static HTML or Drupal)


## Assumptions

+ [dataset](https://github.com/caltechlibrary/dataset) collections are sufficient to hold metadata and media
+ small number of curratorial users
+ larger number of readers or potential depositors
+ configurable workflows are a requirement
+ workflows describe capabilities (permissions)
+ workflows describe a work queue (object state)
+ use existing object scheme (e.g. EPrints Oral Histories)
+ authentication is external (e.g. OAuth or Basic Author)
+ search and query are handle independently (e.g. Lunr)
+ UI can be implemented using HTML 5 elements and minimal JavaScript 


## Limitting features and complexity

One of the most complicated parts of a digital object repository
is managing customization,  managing users, workflows, permissions 
and enforcing storage scheme.  To simplify this **AndOr** is 
using a workflow/work queue oriented permission scheme. The permission 
scheme is configured/managed outside the web UI of **AndOr**. 
Users are define/managed outside the web UI of **AndOr**.
Password management is deferred to the authenticating service. 
Managing users, workflows and permissions reduces the end points 
needed in the API and radically reduces the lines of code needed 
to produce a modern web UI.

Two API end points would be required `/COLLECTION_NAME/objects/OBJECT_ID` 
to provide an Object's details and `/COLLECTION_NAME/objects/`. The
later may accept a filter by queue/workflow name. All other end 
points are static resources (e.g. index.html, CSS, lunr indexes). 
The two end points reflects our experience with the EPrints REST API.
The API also only supports GET and POST and always recieves or responds
with JSON encoded objects. 

Initially the UI will consist of listing records, viewing record 
details, list of records by work queue (workflow state), 
a page to create/read/edit individual objects and a login page.

Authenticated users return their authenticated user id, e.g. 
(e.g.  ORCID using OAuth). The user id will be used to map users
to their workflows. 

Unauthenticated users are treated as the "anonymous" user and
are restricted by workflows available for that user. This is how
you would control having a dark versus publically visible repository.

A user's membership in workflows/queue defines their permissions. 

**AndOr** is built on dataset so all objects must have unique ids. 
Objects may include attached documents which will be versioned 
automatically (a feature available in an upcoming release of
dataset). 

**AndOr** does not directly support deleting objects. You can
create this ability through creating a "delete" workflow type.
Objects flagged as deleted in the API will would be changed to
the queue status of "deleted". A separate offline process could
be developed for garbage collecting those records at a future 
date (this is how EPrints works).

## Additional ideas

+ [Workflow use cases](docs/Workflow-Use-Cases.html)
    + [Workflow Scheme](docs/Workflow-Scheme.html)
    + [AndOr User Scheme](docs/User-Scheme.html)
+ [Setting Up AndOr](docs/Setting-up-AndOr.html)



