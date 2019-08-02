+++
markup = "mmark"
+++


# And/Or

> <span class="red">An</span>other <span class="red">d</span>igital / <span class="red">O</span>bject <span class="red">r</span>epository

This is a concept document for a very light weight digital object
repository implemented as "a multi-user version of 
[dataset](https://caltechlibrary.github.io/dataset) with a web 
based GUI". It targets the ability to curate 
metadata objects and attachments outside the scope of 
our existing repositories.  

**And/Or** is based on [dataset](https://caltechlibrary.github.io/dataset).
It is a JSON API plus HTML, CSS and JavaScript to provide a web 
GUI interface for curating objects.  A minimum running system 
would consist of only one or two pieces of software. The minimum 
would be a web server[^1] plus the **And/Or** service supporting 
multi-user interaction with dataset collections.  Additional
functionality or public website generation would be provided
by other services. E.g. a search service can be easily implemented
using Python and Lunr[^2].  A public website can be generated via 
a similar process we used to build feeds.library.caltech.edu.

**And/Or** is a extremely narrowly scoped web service. The focus 
is __ONLY__ on currating objects and related attachments. 

Limiting **And/Or**'s scope leads to a simpler system. Code 
needed is limited to **And/Or** web service plus the HTML, 
CSS and JavaScript needed for an acceptable UI[^3].

This particular architecture aligns with small machine hosting
and cloud hosting keeping recurring costs to a minimum. 
In the cloud it should work on a tiny or small EC2 instance.
In house hosting could as light weight as Raspberry Pi 4 or
as elaborate as a server with attached NAS[^4].


## Goals

+ Provide a curatorial platform for metadata outside our existing repositories
+ Provide an __interim curration option__ for EPrints repositories requiring migration
+ Thin stack 
    + No RDMS requirement (only And/Or and a web server)
    + Be easier than migrating our EPrints
    + Be faster than EPrints under load
    + Be simpler than EPrints, Invenio, Drupal/Islandora
+ Use existing schema 
+ Support role based workflows
+ Support versioned attached media files
+ Support continuous migration
+ Support alternative front ends (e.g. Drupal)
+ Play nicely with other web based services


## Project Assumptions

+ [dataset](https://github.com/caltechlibrary/dataset) collections are sufficient to hold metadata and media
+ UI can be implemented using HTML 5 elements, minimal JavaScript and CSS
+ Small number of curatorial [users](docs/User-Scheme.html)
+ Configurable [roles](docs/Roles-Scheme.html) are a requirement
    + roles describes permissions for objects in given states
    + roles describes states an object can be asigned to
+ Use existing object scheme (e.g. EPrints XML in Oral Histories)
+ Authentication is external (e.g. Basic Auth, JWT, Shibboleth, OAuth 2)
+ Search and query are handle independent of API
    + e.g. Solr, Elastric Search, Bleve, Lunr 


## Limiting features and complexity

Some of the most complicated parts of digital object repositories
are managing customization, managing users, manage roles,
manage permissions and enforcing storage scheme.  **And/Or**'s 
simplification involves either avoiding the requirement by relocating
it to an appropriate external system and by keeping the scope of
what **And/Or** does very narrow.  

Examples--

+ Authentication is handle externally. That means we don't need to create UI to manage passwords and user profiles. Volatile and sensitive data is outside of **And/Or** so it can't be stolen from **And/Or**.
+ **And/Or** itself is a simple web API that accepts URL requests 
and hands back JSON. It supports a small number of URL end points that support specific actions with one HTTP method (e.g. GET or POST).
+ Object sheme is determined at time of migrating into  the dataset collection. **And/Or** provides no customization.  If you want to change your data shapes you write a script to do that and you change your HTML form.
+ If you need additional end points beyond what **And/Or** provides (e.g. a search engine service, a electronic thesis workflow) you create those as micro services either behind the same webserver or else where.

The web browser creates the illusion of a unified software system
or single process. A single application is not required to support all
desire functionality (e.g. curration and public website generate) because
**And/Or** fits a composible model of all web applications that supply
an API.  Customization is deferred to other micro services and external 
systems (e.g. looking up something at datacite.org or orcid.org)

Some features are unavoidable in curation tool. Repositories run
on the assumption of users and roles. Interestingly it 
doesn't require users and roles be manage through the web. 
Setting up users and roles can be managed through simpler command 
line tools and configuration files.  This is reasonable in large part 
because **And/Or** off loads identify management and can be restarted
quickly.

By focusing on a minimal feature set and leveraging technical
opportunities that already exist we can radically
reduce the lines of code written and maintained. 


### End points map directly to existing dataset operations

dataset supports the a limited set of operations in **And/Or**.
The operations are "keys", "create", "read", "update", "delete",
"attach", "attachments", "detach" and "prune".
These map to URL paths supporting a single HTTP Method.

+ `/COLLECTION_NAME/keys/` (GET) all object keys
+ `/COLLECTION_NAME/keys/OBJECT_STATES` (GET) get a list of objects with the given states
+ `/COLLECTION_NAME/create/OBJECT_ID` (GET) to provide an Object's creates a new object, OBJECT_ID must be unique to succeed
+ `/COLLECTION_NAME/read/OBJECT_IDS` (GET) if single object_id return record otherwise a list of objects is returned
+ `/COLLECTION_NAME/update/OBJECT_ID` (POST) to update an object
+ `/COLLECTION_NAME/delete/OBJECT_ID` (POST) to delete an object
+ `/COLLECTION_NAME/attach/OBJECT_ID/SEMVER` (POST) attach a document to object
+ `/COLLECTION_NAME/attachments/OBJECT_ID` (GET) list attachments to object
+ `/COLLECTION_NAME/detach/OBJECT_ID/SEMVER/ATTACHMENT_NAME` (GET) get an attachments to object
+ `/COLLECTION_NAME/prune/OBJECT_ID/SEMVER/ATTACHMENT_NAME` (POST) remove an attachment version


One additional end point is needed beyond dataset. We need to assign
object states to enable workflows.

+ `/COLLECTION_NAME/assign/OBJECT_ID/OLD_STATE/NEW_STATE` (POST) to delete an object

All other end points are static resources (e.g. HTML files, 
CSS, JavaScript).  


### Examples of composibility

When listing a collection objects suggests the need for paging.
"keys" can be used client side to create a segmented key list and when
combined with "read" which can accept more than one id we now have
a means of paging through a collection's objects.

A collection could be publically viewable (e.g. Oral Histories)
by writing a script that reads the dataset collection **And/Or** is
referencing rendering web pages appropriately (e.g. like we do
with feeds).

A search interface can be created by indexing the dataset collection
and presenting a search UI. (E.g. Like Stephen's demo of Lunr 
providing a search Caltech People pages).

A deposit system could be created as a microservice in Drupal 
to accept metadata and documents before handing them of to 
**And/Or** via a service account.

### Building a UI

Five pages would need to be designed and implemented in HTML, CSS and
JavaScript for our proof of concept.

1. Login and landing page 
2. Display List records (filterable by object state)
3. Display Object details 
4. Create/edit Object details
5. Page to display user roles

For public facing content (e.g. things Google, Bing, et el. 
should find and index) can be deployed separately by 
process similar to how feeds.library.caltech.edu works.
This also keeps **And/Or** simple with fewer requirements.


### user/role/object state is a simple model

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


### Under the hood

**And/Or** is built on [dataset](https://caltechlibrary.github.io/dataset).
It is a thin layer on top of existing dataset functionality.
__dataset__ supplies attachment versioning, **And/Or** exposes that
in the attachment related end points. If dataset gained the ability
to version JSON documents (e.g. stored diffs of JSON documents[^5]),
that functionality becomes easily wrapped by **And/Or**.


Like EPrints **And/Or** should not directly support deleting objects.
Instead the concept of deletion in **And/Or** is to assign the object's
`._State` value to "deleted". This makes deletion work like a Mac's 
trashcan and fully deleting objects would be accomplished by
a separte process performing emptying the trash[^6].


[^1]: NginX and Apache provide authentication mechanisms such as Basic AUTH, Shibboleth and OAuth 2.

[^2]: [Lunr](https://lunrjs.com) is a browser friendly indexing and search library that can now be supported server side too via Python.

[^3]: UI, user interface, the normal way a user interacts with a website

[^4]: NAS, network attached storage similar to what are now common in research labs

[^5]: This could be done in the manner of EPrints which can show a diff of the EPrint XML document

[^6]: Empting the trash boils down to traversing all collecting the keys of objects that are in the `._State` == "deleted" and then removing the content from disc.
