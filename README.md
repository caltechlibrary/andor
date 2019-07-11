
# AndOr

> <span class="red">An</span>other <span class="red">d</span>igital <span class="red">O</span>bject <span class="red">r</span>epository

This is concept document for a very light weight digital object
repository. It would serve as an interim repository for some of our
simple EPrints deployments. **Andor** would be built from 
[dataset](https://caltechlibrary.github.io/dataset) 
collections, a simple JSON oriented web API and UI created via
HTML 5 and JavaScript held in a static htdocs directory.  It is 
an exercise in explorating a minimalist useful digital object
repository. 

## Goals

+ Provide an __interum option__ for simple EPrints repositories
+ Simpler than EPrints 
+ Simpler than Invenio
+ Simpler than Drupal
+ Thin stack (e.g. No RDMS requirement, only AndOr and a web server)
+ Agnostic about the shape of objects stored 
+ Support versionable attached media files
+ Support continuous migration
+ Use the web browser as a smart client
+ Support multiple front ends (e.g. static HTML or Drupal)


## Assumptions

+ [dataset](https://github.com/caltechlibrary/dataset) collections are sufficient to hold metadata and media
+ small number of curratorial users, larger number of readers and potential depositors
+ configurable workflows are a requirement
+ workflows describe capabilities (permissions) and are queue based
+ use existing object scheme (e.g. EPrints Oral Histories)
+ authentication is external (e.g. PAM, OAuth2, Secrets store or Basic Author)
+ search and discovery are handle independently (e.g. Lunr client or server side) can be used to query objects
+ UI can be implemented using HTML 5 elements and minimal JavaScript 
    + This would allow **AndOr** to sit behind Drupal easily

## Limitting features and complexity

One of the most complicated parts of a digital object repository
is enforcing and managing users, workflows, permissions and storage
scheme.  To simplify this **AndOr** is using a workflow/queue 
oriented permission scheme. The permission scheme is configured 
outside the UI of **AndOr**. It is defined/managed in configuration 
files.  Users are also defined/managed outside of **AndOr** 
in configuration files. Password management is deferred to the 
authenticating service. Managing users, workflows and permissions 
from configuration files reduces the end points needed in the 
API and radically reduces the lines of code needed to produce a 
modern web UI.

Two API end points would be required `/COLLECTION_NAME/objects/OBJECT_ID` 
to provide an Object's details and `/COLLECTION_NAME/objects/`. The
later may accept a filter by queue/workflow name. All other end 
points are static resources (e.g. index.html, CSS, lunr indexes). 
These end points were inspired by EPrints' REST API. The API also
only supports GET and POST of JSON encoded objects. This
simplifies the translation in the API and fully supports nested 
fields which will help keep web form construction simpler.

Initially the UI will consist of listing records, viewing record 
details, list of records by work queue (workflow state) and 
a screen to display/create/edit objects and a login page.

Authenticated users return their authenticated user id, e.g. 
(e.g.  ORCID using OAuth). The user id will be used to map users
to their workflows. 

Unauthenticated users are treated as the "anonymous" user and
are restricted by workflows available for that user. This is how
you would control having a dark versus publically visible repository.

A user's membership in workflows defines their permissions. 

**AndOr** is built on dataset so all objects must have unique ids. 
Objects may include attached documents which will be versioned 
automatically (a feature available in an upcoming release of
dataset). 

Objects flagged as deleted in the API will simply be pushed to the
workflow of "deleted". A separate offline process can be developed
for garbage collecting those records at a future date (inspired 
by EPrints).

## Additional ideas

+ [Workflow use cases](docs/Workflow-Use-Cases.html)
    + [Workflow Scheme](docs/Workflow-Scheme.html)
    + [AndOr User Scheme](docs/User-Scheme.html)
+ [Setting Up AndOr](docs/Setting-up-AndOr.html)



