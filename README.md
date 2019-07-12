
# AndOr

> <span class="red">An</span>other <span class="red">d</span>igital <span class="red">O</span>bject <span class="red">r</span>epository

This is a concept document for a very light weight digital object
repository. If prototyped successfully it could serve as an
interim repository for some of our EPrints repositories.

**AndOr** would be built from [dataset](https://caltechlibrary.github.io/dataset) collections, simple  configuration files, a minimal JSON
semi-RESTful web API plus statically hosted HTML and Javascript.
A running system would probably consist of only a few pieces of
software. The software stack could be as light as a web server
(e.g. NginX, Apache) providing integration options for authentication
(e.g. BasicAUTH, Shibboleth) plus the **AndOr** web service
written in Go. If the collection was larger we could create
a search interface via Python and Lunr, otherwise search could be
implemented browser side with Lunr.  This arrangement has the
advantage of limiting the code to be written to the
web service, some HTML pages and a bit of JavaScript.  This
particular architecture is much more aligned with cloud
hosting and keeping hosting cost to a minimum. It should work
on small to medium EC2 instances with an S3 bucket for larger
collections. A more elaborate but possibly more cost effective
implementation would be replacing the web server with Cloud Front,
the static file storage with S3 and running the API service
in AWS's container or via their Lambda service.


## Goals

+ Provide an __interim option__ for EPrints repositories
+ Thin stack (e.g. No RDMS requirement, only AndOr and a web server)
    + Be simpler than EPrints, Invenio, Drupal/Islandora
    + Be easier than migrating our EPrints repositories to the cloud
+ Use existing schema
+ Support versioned attached media files
+ Support diffs for JSON metadata changes
+ Support continuous migration
+ Support alternative front ends (e.g. Drupal)
+ Be extremely easy to migrate out of it (e.g. to Zenodo or Archipelago)
+ Be faster than EPrints particularly when under load


## Project Assumptions

+ [dataset](https://github.com/caltechlibrary/dataset) collections are sufficient to hold metadata and media
+ a small number of curatorial users
+ a larger number of readers or potential depositors
+ configurable workflows are a requirement
+ workflows describe capabilities (permissions)
+ workflows describe a work queue (object state)
+ use existing object scheme (e.g. EPrints Oral Histories)
+ authentication is external (e.g. Basic AUTH via web server)
+ search and query are handle independently of API
    + e.g. Lunr either browser or server side
+ UI can be implemented using HTML 5 elements and minimal JavaScript


## Limiting features and complexity

Some of the most complicated parts of a digital object repositories
are code managing customization, managing users, workflows,
permissions and enforcing storage scheme.  To simplify this **AndOr** is
using a workflow/work queue oriented permission scheme. The permission
scheme is configured/managed outside the web UI of **AndOr**. It's
just a mater of creating some JSON objects and storing them in a
dataset collection that **AndOr** reads.  Users are also define/managed
outside the web UI of **AndOr**.  User password management is deferred
to the authenticating service.  Managing users, workflows and
permissions reduces the end points needed in the API and radically
reduces the lines of code needed to produce a modern web UI.

Two API end points would be required `/COLLECTION_NAME/objects/OBJECT_ID`
to provide an Object's details and `/COLLECTION_NAME/objects/`. The
later may accept a filter by queue/workflow name. All other end
points are static resources (e.g. index.html, CSS, Lunr indexes).
Reducing service to two end points reflects our experience with
the EPrints REST API.  The API end points only need to support GET
and POST operations and can be restricted to expecting JSON
object responses which should simplify writing the web forms for
creating and updating objects.

Four pages would need to be designed and implemented in HTML, CSS and
JavaScript.

1. Display List records
2. Display Object details
3. Search UI
4. Login and landing page

Authenticated users return their authenticated user id
(e.g. ORCID if authenticating via orcid.org). The user id
will be used to map users to their workflows.

Unauthenticated users are treated as the "anonymous" user and
are restricted by workflows available for that user. This is how
you would control having a dark versus publicly visible repository.

A user's membership in workflows/queue defines their permissions.

**AndOr** is built on dataset so all objects must have unique ids.
Objects may include attached documents which will be versioned
automatically (a planned feature available in an upcoming
release of dataset).  We can also store diffs of the JSON
documents for metadata versioning.

Like EPrints **AndOr** does not directly support deleting objects.
Instead we can create the illusion of deleting objects by putting
objects into a "delete" queue.

## Additional ideas

+ [Workflow use cases](docs/Workflow-Use-Cases.html)
    + [Workflow Scheme](docs/Workflow-Scheme.html)
    + [AndOr User Scheme](docs/User-Scheme.html)
+ [Setting Up AndOr](docs/Setting-up-AndOr.html)
+ [Oral Histories Use Case](docs/Oral-Histories-Use-Case.html)
    + [Proposed Schedule for Proof of Concept](docs/Schedule.html)


