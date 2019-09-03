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
GUI interface for curating objects.  A minimum running system 
would consist of two pieces of software.  The minimum is 
a web server[^1] plus the **And/Or** service supporting 
multi-user interaction with dataset collections.  Additional
functionality would be provided by other systems or services[^2].

**And/Or** is a extremely narrowly scoped web service. The focus 
is __ONLY__ on currating objects and related attachments. 

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
+ Support alternative front ends (e.g. Drupal)


## Project Assumptions

+ [dataset](https://github.com/caltechlibrary/dataset) collections are sufficient to hold metadata and media
+ UI can be implemented using HTML 5 elements, minimal JavaScript and CSS
+ Small number of curatorial [users](docs/User-Scheme.html)
+ Configurable [roles](docs/Roles-Scheme.html) are a requirement
    + roles describes permissions for objects in given states
    + roles describes states an object can be asigned to
+ Use existing object scheme (e.g. EPrints XML in Oral Histories)
+ Authentication is external (e.g. Basic Auth, JWT, Shibboleth, OAuth 2)
+ Other systems handle any additoinal requirements


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
+ Object sheme is determined at time of migrating into  the dataset collection. **And/Or** provides no customization.  If you want to change your data shapes you write a script to do that and you change your HTML form.
+ If you need additional end points beyond what **And/Or** provides (e.g. a search engine service, a electronic thesis workflow) you create those as micro services either behind the same webserver or else where.

The web browser creates the illusion of a unified software system
or single process. A single application is not required to support all
desire functionality (e.g. curration and public web consumption) because
**And/Or** uses a composible model available to all web applications.
Customization is deferred to other micro services and external 
systems (e.g. looking up a record on datacite.org or orcid.org)

Some features are unavoidable in curation tool. Repositories run
on the assumption of users and roles. Interestingly it 
doesn't require users and roles be manage through the web. 
Setting up users and roles can be managed through simple to implement
command line tools and configuration files.  This is reasonable in 
large part because **And/Or** off loads identify management 
and can be restarted quickly (i.e. configuration files are easily
parsed).

By focusing on a minimal feature set and leveraging technical
opportunities that already exist we can radically
reduce the lines of code written and maintained. 

## Under the hood

**And/Or**'s engine is [dataset](https://github.com/caltechlibrary/dataset).

> End points map directly to existing dataset operations

dataset operations supported in **And/Or** are "keys", "create", 
"read", "update", "delete", "attach", "attachments", "detach" 
and "prune". These map to URL paths each supporting a single 
HTTP Method.

+ `/COLLECTION_NAME/keys/` (GET) all object keys
+ `/COLLECTION_NAME/create/OBJECT_ID` (GET) to creates an Object, an OBJECT_ID must be unique to succeed
+ `/COLLECTION_NAME/read/OBJECT_IDS` (GET) returns on or more objects, if more than one is requested then an array of objects is returned.
+ `/COLLECTION_NAME/update/OBJECT_ID` (POST) to update an object
+ `/COLLECTION_NAME/delete/OBJECT_ID` (POST) to delete an object
+ `/COLLECTION_NAME/attach/OBJECT_ID/SEMVER` (POST) attach a document to object
+ `/COLLECTION_NAME/attachments/OBJECT_ID` (GET) list attachments to object
+ `/COLLECTION_NAME/detach/OBJECT_ID/SEMVER/ATTACHMENT_NAME` (GET) get an attachments to object
+ `/COLLECTION_NAME/prune/OBJECT_ID/SEMVER/ATTACHMENT_NAME` (POST) remove an attachment version

One additional end point is needed beyond dataset. We need to assign
object states to enable workflows.

+ `/COLLECTION_NAME/assign/OBJECT_ID/NEW_STATE` (POST) to assign an object to a new state

**And/Or** is a thin layer on top of existing dataset functionality.
E.g. dataset supplies attachment versioning, **And/Or** exposes that
in the attachment related end points. If dataset gained the ability
to version JSON documents (e.g. stored diffs of JSON documents[^4]),
that functionality could be included in **And/Or**.

### Web UI

Five pages would need to be designed per collection and 
implemented in HTML, CSS and JavaScript for our proof of concept.

1. Login and landing page 
2. Display List records (filterable by object state)
3. Display Object details 
4. Create/edit Object details
5. Page to display user roles

**And/Or** is NOT for public facing content 
(e.g. things Google, Bing, et el.  should find and index) 
Machanisms for public facing content should be deployed 
separately by process similar to how feeds.library.caltech.edu 
works. This keeps **And/Or** simple with fewer requirements.

### Examples of composibility

When listing a collection objects suggests the need for paging.
"keys" can be used client side to create a segmented key list and when
combined with "read" which can accept an id list, thus we get
paging for a collection's objects by getting the universe of keys,
and then requesting to "read" for the keys we want to display
in a page.

A search interface could be created as a microserve in the manner 
of Stevens' Lunr demo for Caltech People. If **And/Or** and the
search microserver are behind the same web server you could present
both services using a common URL namespace (Apache or NingX are
good candites from a front facing web server integrating **And/Or**
and your search system).

A deposit system could be created as a microservice (e.g. 
in Drupal) to accept metadata and documents before handing 
them of to **And/Or** via a service account.


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
create, read, update).

### The special case of deleting objects 

Like EPrints **And/Or** should not directly support deleting objects.
Instead the concept of deletion in **And/Or** is to assign the object's
`._State` value to "deleted". This makes deletion work like a Mac's 
trashcan and fully deleting objects would be accomplished by
a separte process performing emptying the trash[^5].


[^1]: NginX and Apache provide authentication mechanisms such as Basic AUTH, Shibboleth and OAuth 2.

[^2]: Public websites can be generated feeds.library.caltech.edu, a search interface can be implemented with [Lunr](https://lunrjs.com).

[^3]: UI, user interface, the normal way a user interacts with a website

[^4]: This could be done in the manner of EPrints which can show a diff of the EPrint XML document

[^5]: Empting the trash boils down to traversing all collecting the keys of objects that are in the `._State` == "deleted" and then removing the content from disc.
