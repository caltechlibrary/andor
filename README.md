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
would consist of only two or three pieces of software. The minimum 
would be a web server[^1] plus the **And/Or** service supporting 
multi-user interaction with dataset collections.  Depending on size 
you could create a micro service providing search via Python and 
Lunr[^2] or generate Lunrjs indexes to run browser side.  This 
arrangement has the advantage of limiting development of new
code to be written to the **And/Or** web service plus the HTML, 
CSS and JavaScript needed for an acceptable UI[^3].

This particular architecture aligns with small machine hosting
and cloud hosting keeping recurring costs to a minimum. 
In the cloud it should work on a small to medium EC2 instance.
A more elaborate version could be implemented using Cloud Front, 
S3, and running the API service in an AWS container.
In house hosting could as light weight as Raspberry Pi 4 or
as elaborate as a server with attached NAS[^4].


## Goals

+ Provide a curatorial platform for metadata outside our existing repositories
+ Provide an __interim option__ for EPrints repositories requiring migration
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


## Project Assumptions

+ [dataset](https://github.com/caltechlibrary/dataset) collections are sufficient to hold metadata and media
+ UI can be implemented using HTML 5 elements, minimal JavaScript and CSS
+ Small number of curatorial users, larger number of readers
+ Configurable roles are a requirement
    + roles describe capabilities (permissions)
    + roles describe one or more work queues 
    + queues are an object state
+ Use existing object scheme (e.g. EPrints XML in Oral Histories)
+ Authentication is external (e.g. Basic Auth, JWT, Shibboleth, OAuth 2)
+ Search and query are handle independent of API
    + e.g. Lunr either browser or server side


## Limiting features and complexity

Some of the most complicated parts of digital object repositories
are managing customization, managing users, manage roles,
manage permissions and enforcing storage scheme.  **And/Or**'s 
simplification involves either avoiding the requirements or relocating
them to an appropriate external system.  

Examples--

+ Authentication is handle externally. That means we don't need to create UI to manage passwords, the volatile and sensitive data is outside of **And/Or** 
+ **And/Or** itself is a simple web API that accepts URL requests 
and hands back JSON. The shape of the JSON is determined at time of
migrating into **And/Or**. There is no customization.  If you want to change your data shapes you write a script to do that or change your input form.
+ If you need additional end points beyond what **And/Or** provides (e.g. a search engine service) you supply those as micro services behind the same web server

The web browser itself creates the illusion of a unified software system
or single process. A single application is not required to support desire
functionality. Customization then can be deferred to other micro services
or even external systems (e.g. looking up something at datacite.org or
orcid.org).

Some features are unavoidable. The repository problem requires managing
users and roles. It doesn't require users and roles
be manage through the web. Setting up users and roles can be 
managed through simpler command line tools and configuration files.
The is reasonable in large part because you've off loaded 
identify management already. 

By focusing on a minimal feature set and leveraging technical
opportunities that already exist we can radically
reduce the lines of code written and maintained. 

### End points for our API map to dataset operations

+ `/COLLECTION_NAME/keys/` (GET) to list objects keys
+ `/COLLECTION_NAME/create/OBJECT_ID` (GET) to provide an Object's creates a new object, OBJECT_ID must be unique to succeed
+ `/COLLECTION_NAME/read/OBJECT_IDS` (GET) if single object_id return record otherwise a list of objects is returned
+ `/COLLECTION_NAME/update/OBJECT_ID` (POST) to update an object
+ `/COLLECTION_NAME/delete/OBJECT_ID` (POST) to delete an object

"keys" can be filtered by role's queue name. Paging can be implemented
client side by segmenting the key list returned.
All other end points are static resources (e.g. HTML files, 
CSS, JavaScript and Lunrjs indexes, a public faces website).  
We can reducing our requirements to two end points because 
we've discovered that is all we needed based on our work
with [EPrints](https://www.eprints.org "A repository system developed at University of South Hampton"). Everything else can be synthesized 
from a simple key list and object access.


### Building a UI

Five pages would need to be designed and implemented in HTML, CSS and
JavaScript for our proof of concept.

1. Login and landing page 
2. Display List records (filterable by work queue)
3. Display Object details 
4. Create/edit Object details
5. Search UI

For public facing content (e.g. things Google, Bing, et el. 
should find and index) can be deployed separately by 
process similar to how feeds.library.caltech.edu works.
This also keeps **And/Or** simple with fewer requirements.


### user plus role, a simple model

An authenticated user exposes their user id to 
**And/Or**'s web service (e.g. via a JSON Web Token or 
Basic Auth header).  A user's id maps to membership in roles. 
The role defines access to queues, queues are list objects
with a matching value of `._Queue`.

Unauthenticated users are treated as the "anonymous" user and
are restricted by roles available for the "anonymous" user. 
This is how you could implement both dark and public
repositories.

Complicated use cases like electronic thesis deposits
or faculty self deposit of articles could be implemented as 
separate specialized micro services that interact with
**And/Or** via proxy service accounts.


### Under the hood

**And/Or** is built on [dataset](https://caltechlibrary.github.io/dataset).
Objects may include attached documents which can be versioned 
automatically. If metadata versioning becomes required dataset 
can be extended to store diffs as well as the JSON documents.

Like EPrints **And/Or** does not directly support deleting objects.
Instead it can create the illusion of deleting objects by putting
objects into a "deleted" queue which you can exclude from your
roles or garbage collection through a separate process.


## Additional ideas

+ Use cases
    + [Users, Roles and Queues](docs/Role-Use-Cases.html)
+ Concept proofs
    + [People and Groups](docs/people-groups.html)
    + [Migrating an EPrints Repository](docs/migrating-eprints.html) 
    + [Oral Histories](Oral-Histories-as-Proof-of-Concept.html)
+ Scheme walk through
    + [User Scheme](docs/User-Scheme.html)
    + [Role Scheme](docs/Role-Scheme.html)
    + [Queue Scheme](docs/Queue-Scheme.html)
    + [Object Scheme](docs/Object-Scheme.html)




[^1]: NginX and Apache provide authentication mechanisms such as Basic AUTH, Shibboleth and OAuth 2.

[^2]: [Lunr](https://lunrjs.com) is a browser friendly indexing and search library that can now be supported server side too via Python.

[^3]: UI, user interface, the normal way a user interacts with a website

[^4]: NAS, network attached storage similar to what are now common in research labs
