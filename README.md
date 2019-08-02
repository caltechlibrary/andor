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

**And/Or** is a extremely narrowly scoped web service. The focus 
is __ONLY__ on currating objects and related attachments. This is done
to keep the service simple. Other services and systems can pickup
other responsiblies such as create public websites, supporting
complex workflows like thesis deposits. It embodies the Unix
philophy of single purpose tools doing spefic things while 
remaining composible for more complex demands.

This particular architecture aligns with small machine hosting
and cloud hosting keeping recurring costs to a minimum. 
In the cloud it should work on a small to medium EC2 instance.
A more elaborate version could be implemented using Cloud Front, 
S3, and running the API service in an AWS container.
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
    + roles describes capabilities (permissions)
    + roles describes object states for the capabilities
    + roles describes states an object can be put in
+ Use existing object scheme (e.g. EPrints XML in Oral Histories)
+ Authentication is external (e.g. Basic Auth, JWT, Shibboleth, OAuth 2)
+ Search and query are handle independent of API
    + e.g. Solr, Elastric Search, Bleve, Lunr 


## Limiting features and complexity

Some of the most complicated parts of digital object repositories
are managing customization, managing users, manage roles,
manage permissions and enforcing storage scheme.  **And/Or**'s 
simplification involves either avoiding the requirement by relocating
it to an appropriate external system and keeping the scope of
what **And/Or** does very narrow.  

Examples--

+ Authentication is handle externally. That means we don't need to create UI to manage passwords. Volatile and sensitive data is outside of **And/Or** 
+ **And/Or** itself is a simple web API that accepts URL requests 
and hands back JSON. The shape of the JSON is determined at time of
migrating into **And/Or**. There is no customization.  If you want to change your data shapes you write a script to do that or change your input form.
+ If you need additional end points beyond what **And/Or** provides (e.g. a search engine service) you those as micro services behind the same web server or externally

The web browser itself creates the illusion of a unified software system
or single process. A single application is not required to support desire
functionality (e.g. curration and public website generate). Customization 
then can be deferred to other micro services or even external systems (e.g. looking up something at datacite.org or orcid.org, rendering the public
website).

Some features are unavoidable. The repository problem requires managing
users and roles. It doesn't require users and roles
be manage through the web. Setting up users and roles can be 
managed through simpler command line tools and configuration files.
This is reasonable in large part because we've off loaded 
identify management already and restarting the **And/Or** service is fast. 

By focusing on a minimal feature set and leveraging technical
opportunities that already exist we can radically
reduce the lines of code written and maintained. 

### End points for our API map directly to dataset operations

dataset has the following operations available in **And/Or**--- 
keys, create, read, update, delete, attach, attachments, detach and prune.
They map to URL paths as follows.

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
CSS, JavaScript and Lunrjs indexes, a public facing website).  


### Notes on using the API

"keys" can be used client side to create a segmented key list and when
combined with "read" (which can accept more than one id) it can
but used for paging.

If a collection should be publically viewable (e.g. Oral Histories)
that website would be generated from the content currated in **And/Or**.
The generation could happen on a regular schedule like we do with 
feeds.library.caltech.edu or we could implement a object status
service that lists changes in object states (e.g. an array of
object ids, timestamps, current state and last operation). This could
be done via a streaming log or as a web service itself. The stream
of object state change changes would then be used to trigger a 
refresh in the public site.


### Building a UI

Five pages would need to be designed and implemented in HTML, CSS and
JavaScript for our proof of concept.

1. Login and landing page 
2. Display List records (filterable by work state)
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
The role defines access to states, states are list objects
with a matching value of `._State`.

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
objects into a "deleted" state which you can exclude from your
roles or garbage collection through a separate process.




[^1]: NginX and Apache provide authentication mechanisms such as Basic AUTH, Shibboleth and OAuth 2.

[^2]: [Lunr](https://lunrjs.com) is a browser friendly indexing and search library that can now be supported server side too via Python.

[^3]: UI, user interface, the normal way a user interacts with a website

[^4]: NAS, network attached storage similar to what are now common in research labs
