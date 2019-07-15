+++
markup = "mmark"
+++


# AndOr

> <span class="red">An</span>other <span class="red">d</span>igital <span class="red">O</span>bject <span class="red">r</span>epository

This is a concept document for a very light weight digital object
repository. If prototyped successfully it could serve as an
interim repository for the EPrints repositories we plan to migrate.

**AndOr** would be built from [dataset](https://caltechlibrary.github.io/dataset)
collections, a semi-RESTful web API, and HTML with JavaScript.
A running system would probably consist of only two or three
pieces of software. The minimum would be a web server[^1] 
plus the **AndOr** service supporting interaction with the
dataset collection.  If the collection was larger you could create
a micro service providing search via Python and Lunr[^2].
This arrangement has the advantage of limiting the code to be written 
to the **AndOr** web service plus the HTML and JavaScript 
needed to create an acceptable UI[^3].

This particular architecture is much more aligned with 
cloud hosting and keeping hosting cost to a minimum. It should work
on small to medium EC2 instances with an S3 bucket for larger
collections. A more elaborate but possibly more cost effective
implementation would be replacing the web server with Cloud Front,
the static file storage with S3 and running the API service
in an AWS container or via the AWS Lambda service.


## Goals

+ Provide an __interim option__ for EPrints repositories requiring migration
+ Thin stack 
    + No RDMS requirement (only AndOr and a web server)
    + Be easier than migrating our EPrints
    + Be faster than EPrints under load
    + Be simpler than EPrints, Invenio, Drupal/Islandora
+ Use existing schema
+ Support versioned attached media files
+ Support continuous migration
+ Support alternative front ends (e.g. Drupal)


## Project Assumptions

+ [dataset](https://github.com/caltechlibrary/dataset) collections are sufficient to hold metadata and media
+ UI can be implemented using HTML 5 elements and minimal JavaScript
+ Small number of curatorial users, larger number of readers
+ Configurable workflows are a requirement
    + workflows describe capabilities (permissions)
    + workflows describe a work queue (object state)
+ Use existing object scheme (e.g. EPrints XML in Oral Histories)
+ Authentication is external (e.g. Basic AUTH via web server)
+ Search and query are handle independently of API
    + e.g. Lunr either browser or server side


## Limiting features and complexity

Some of the most complicated parts of digital object repositories
are managing customization, managing users, manage workflows,
manage permissions and enforcing storage scheme.  **AndOr**'s 
simplification involves either avoiding the requirements or relocating
them external to the system.  

Examples--

+ Authentication is handle externally. That means we don't need to worry about managing passwords, the volatile and sensitive data is outside of our system
+ **AndOr** itself is a simple web API that accepts URL requests 
and hands back JSON. The shape of the JSON is determined at time of
migrating into **AndOr**. There is no customization.  If you want to change your data shapes you write a script to do that or change your input form.
+ If you need additional end points beyond what **AndOr** provides (e.g. a search engine service) you supply those as micro services 

The web browser itself creates the illusion of a unified software system
or single process. A single application is not required to support desire
functionality.

Some features are unavoidable. The repository problem requires managing
users and workflows. It doesn't require users and workflows
be manage through the web. Setting up users and workflows can be 
managed through simpler command line tools in large part because 
you've off loaded identify management already. 

By focusing on a minimal feature set and leveraging technical
opportunities that already exist we can radically
reduce the lines of code written and maintained.

### Two end points for our API

Two web API end points would be required 
`/COLLECTION_NAME/objects/OBJECT_ID` to provide an Object's details 
and `/COLLECTION_NAME/objects/` to list objects. The later may 
accept a filter by queue/workflow name. All other end points are 
static resources (e.g. HTML files, CSS, JavaScript and 
Lunrjs indexes, a public faces website).  We can reducing our
requirements to two end  points because we've already discovered 
it was all we needed to integrate with the EPrints REST API.
Everything else can be synthesized from them.


### Building a UI

Five pages would need to be designed and implemented in HTML, CSS and
JavaScript for our proof of concept.

1. Login and landing page
2. Display List records (filterable by work queue)
3. Display Object details 
4. Create/edit Object details
5. Search UI

For public facing content (e.g. the things you want 
Google, Bing, et el. to find) can be deployed by a simple batch
process that updates a public website. It can be external to the
curatorial aspects of a digital object repository.

### user plus workflow, a simple model

An authenticated user exposes their user id. A user's id 
maps to membership in workflows. The workflow defines access.

Unauthenticated users are treated as the "anonymous" user and
are restricted by workflows available for the "anonymous" user. 
This is how you would control having a dark versus publicly 
visible repository.

Complicated use cases integrations like community deposit could 
then be deferred to a micro service(s) created for that purpose.
The micro services become simpler because of their narrow focus
and limited abilities.

### Under the hood

**AndOr** is built on [dataset](https://caltechlibrary.github.io/dataset).
Objects may include attached documents which can be versioned 
automatically. If metadata versioning becomes required dataset 
can be extended to store diffs as well as the JSON documents.

Like EPrints **AndOr** does not directly support deleting objects.
Instead it can create the illusion of deleting objects by putting
objects into a "delete" queue which is not visible to users.

## Additional ideas

+ [Workflow use cases](docs/Workflow-Use-Cases.html)
    + [Workflow Scheme](docs/Workflow-Scheme.html)
    + [AndOr User Scheme](docs/User-Scheme.html)
+ [Setting Up AndOr](docs/Setting-up-AndOr.html)
+ [Migrating an EPrints repository](docs/migrating-eprints.html)
+ [Oral Histories Use Case](docs/Oral-Histories-Use-Case.html)
    + [Proposed Schedule for Proof of Concept](docs/Schedule.html)



[^1]: NginX and Apache provide authentication mechanisms such as Basic AUTH, Shibboleth and OAuth 2.

[^2]: [Lunr](https://lunrjs.com) is a browser friendly indexing and search library that can now be supported server side too via Python.

[^3]: UI, user interface, the normal way a user interacts with a website
