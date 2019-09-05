+++
markup = "mmark"
+++


# About And/Or

> <span class="red">An</span>other <span class="red">d</span>igital / <span class="red">O</span>bject <span class="red">r</span>epository

This is a concept document for a very light weight digital object
repository implemented as "a multi-user version of 
[dataset](https://caltechlibrary.github.io/dataset) with a web 
based GUI". It targets the ability to curate 
metadata objects and attachments outside the scope of 
our existing repositories.  

**And/Or** is based on [dataset](https://caltechlibrary.github.io/dataset).
It is a web JSON API plus HTML, CSS and JavaScript providing a web 
GUI interface for curating objects.  It is intended
to run using a standard web server like Apache or NginX
and the **And/Or** service running via reverse proxy behind.
For the purposes of a proof of concept the minimum is 
a **And/Or** web server[^1] providing both static website
service plus supporting multi-user interaction with dataset 
collections.  The concept is additional functionality 
is provided by other microservices and systems[^2].

**And/Or** is a extremely narrowly scoped web service. The focus 
is __ONLY__ on currating JSON objects. 

Limiting **And/Or**'s scope leads to a simpler system. Code 
is limited to **And/Or** web service plus the HTML, 
CSS and JavaScript needed for an acceptable UI[^3].

This architecture aligns with small machine hosting
and cloud hosting. Both keeping recurring costs to a minimum. 
**And/Or** could be run on a tiny to small EC2 instance or
on hardware as small as a Rasbpberry Pi.


## Goals

+ Provide a curatorial platform for metadata outside our existing repositories
+ Thin stack 
    + No RDMS requirement (only And/Or and a web server)
    + Be easier to implement than migrating a repository
    + Be faster than EPrints for curating objects
    + Be simpler than EPrints, Invenio and Drupal/Islandora
+ Use existing schema 
+ Support role based workflows
+ Support continuous migration
+ Support alternative front ends (e.g. Drupal in front of And/Or)


## Project Assumptions

+ [dataset](https://github.com/caltechlibrary/dataset) collections are sufficient to hold metadata
+ UI can be implemented using HTML 5 elements, minimal JavaScript and CSS
+ Small number of curatorial [users](docs/User-Scheme.html)
+ Configurable [roles](docs/Roles-Scheme.html) are a requirement
    + roles describes permissions for objects in given states
    + roles describes states an object can be asigned to
+ Use existing object scheme (e.g. EPrints XML in Oral Histories)
+ Authentication is external (e.g. Basic Auth, JWT, Shibboleth, OAuth 2)
+ Other systems handle any additoinal service requirements (e.g. search, storage of uploaded objects)


## Limiting features and complexity

Some of the most complicated parts of digital object repositories
are managing customization, managing users, manage roles,
manage permissions, enforcing storage scheme and presenting
public can private views of respository content.  **And/Or**'s 
simplification involves avoiding functionality provided
by other systems and relocating requirements to an appropriate 
external system while only focusing on the narrow problem
of curating objects. 

Examples--

+ Authentication is handle externally. That means we don't need to create web UI to manage passwords and user profiles. Volatile and sensitive data is outside of **And/Or** so it can't be stolen from **And/Or**.
+ **And/Or** itself is a simple web API that accepts URL requests 
and hands back JSON. It supports a small number of URL end points that support specific actions with one HTTP method (e.g. GET or POST) per end point.
+ Object sheme is determined at time of migrating into the dataset collection. **And/Or** itself provides no mechanism customization beyond the input forms retrieving and submitting objects.  If you want to change your data shapes you write a script to do that and you change your HTML form.
+ If you need additional end points beyond what **And/Or** provides (e.g. a search engine service, a electronic thesis workflow) you create those as micro services either behind the same webserver or else where.

A web browser can create the illusion of a unified software system
or single process. A single application is not required to support all
desire functionality (e.g. curration and public web consumption) because
**And/Or** uses a composible model available to all web applications.
Customization is limited to static HTML, CSS and JavaScript or deferred 
to other micro services and external systems (e.g. looking up a record 
on datacite.org or orcid.org)

Some features are unavoidable in curation tool. Repositories run
on the assumption of users and roles. Interestingly it 
doesn't require users and roles be manage through the web. 
Setting up users and roles can be managed through simple to implement
command line tools and configuration files.  This is reasonable in 
large part because **And/Or** off loads identify management 
and can be restarted quickly (i.e. configuration files are easily
parsed and reloaded).

By focusing on a minimal feature set and leveraging technical
opportunities that already exist we can radically
reduce the lines of code written and maintained. 

## Under the hood

**And/Or**'s JSON document storage engine is [dataset](https://github.com/caltechlibrary/dataset).

> End points map directly to existing dataset operations

dataset operations supported in **And/Or** are "keys", "create", 
"read", "update", "delete".  These map to URL paths each supporting 
a single HTTP Method (either GET or POST).

+ `/COLLECTION_NAME/keys/` (GET) all object keys
+ `/COLLECTION_NAME/create/OBJECT_ID` (GET) to creates an Object, an OBJECT_ID must be unique to succeed
+ `/COLLECTION_NAME/read/OBJECT_IDS` (GET) returns on or more objects, if more than one is requested then an array of objects is returned.
+ `/COLLECTION_NAME/update/OBJECT_ID` (POST) to update an object
+ `/COLLECTION_NAME/delete/OBJECT_ID` (POST) to delete an object

**And/Or** is a thin layer on top of existing dataset functionality.
E.g. dataset supplies attachment versioning, **And/Or** exposes that
in the attachment related end points. If dataset gained the ability
to version JSON documents (e.g. stored diffs of JSON documents[^4]),
that functionality could be included in **And/Or**.

### Web UI

Four pages would need to be designed per collection and 
implemented in HTML, CSS and JavaScript for our proof of concept.

1. Login and landing page 
2. Display List records (filterable by object state)
3. An edit for that serves to display object details in an editable form as well as create, update and retrieve an object by id.
4. Page to display user roles

**And/Or** is NOT for public facing content 
(e.g. things Google, Bing, et el.  should find and index) 
Machanisms for public facing content should be deployed 
separately by processes similar to how feeds.library.caltech.edu 
works. This keeps **And/Or** simple with fewer requirements.

### Examples of composibility

When listing a large collection objects prudence 
suggests the need for paging. After retrieving all keys we can
implement paging by using the "read" method with a list of keys
we want to view.  This allows us to segment a large collection 
into manageable chunks.

A search interface could be created as a microserve in the manner 
of Stevens' Lunr demo for Caltech People. If **And/Or** and the
search microserver are behind the same web server you could present
both services using a common URL namespace (Apache or NingX are
good candites from a front facing web server integrating **And/Or**
and your search system).


### User/role/object state is a simple model

An authenticated user exposes their user id to 
**And/Or**'s web service. The web service can then
retrieve the available roles that scope the permissions
the user has to operate on objects in a given set of states.
The role can also be used to define which objects we show
the user.  This can be implemented with a small number
of functions such as `getUsername()`, `getUserRoles()`, 
`isAllowed()` and `canAssign()`.

Once authorization is calculated then approvided actions
can be handle with simple HTTP handlers that perform a simple
task mapping to an existing dataset function (e.g. keys, 
create, read, update, delete).

### A special case of deleting objects 

While **And/Or** service can delete objects it's more
prudent to take the EPrints approach and define "delete"
as a specific object state. This way you could treat
deleted objects as being in a trashcan and leave actual
deletion for a garbage collection routine.  This  approach would 
make deletion work like a Mac's trashcan and fully deleting 
objects would be accomplished by a separte process performing 
emptying the trash[^5].


[^1]: NginX and Apache could provide authentication mechanisms such as Basic AUTH, Shibboleth and OAuth 2 and pass them back to a real And/Or implementation.

[^2]: Public websites can be generated feeds.library.caltech.edu, a search interface can be implemented with [Lunr](https://lunrjs.com).

[^3]: UI, user interface, the normal way a user interacts with a website

[^4]: This could be done in the manner of EPrints which can show a diff of the EPrint XML document

[^5]: Empting the trash boils down to traversing all collecting the keys of objects that are in the `._State` == "deleted" and then removing the content from disc.
