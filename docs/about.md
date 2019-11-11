+++
markup = "mmark"
+++


# About And/Or

> <span class="red">An</span>other <span class="red">d</span>igital / <span class="red">O</span>bject <span class="red">r</span>epository

This is a concept document for a very light weight digital object
repository implemented as "a multi-user version of 
[dataset](https://caltechlibrary.github.io/dataset) with a web 
based GUI". It targets the ability to curate metadata objects 
outside the scope of our existing repository systems.  

**And/Or** is based on [dataset](https://caltechlibrary.github.io/dataset)
and `libdataset`.  __libdataset__ is a C Shared library made accessible 
via Python 3.7's ctypes package. The primary differences between 
**And/Or** built-on  __libdataset__ and __dataset__ tool is that 
the library implements locking via a semifore for disc writes on object 
creation and updates.  It functions like a fedora-lite for a Python 
programming implementing web services. 

When a Go shared library runs under Python it runs in its child process.
The process is managed from the Python program. In this way you can 
efficiently act on a __dataset__ collection and with the use of 
a semiphore dataset can support asynchonous operations on the collection 
without needing to constantly open and close the dataset collection. 
The Python code focuses on providing URL end points, access control
and presentation of static HTML, CSS and JavaScript while you the 
__libdataset__ library handles the ansynchronous update of the 
collection(s) it is managing. 

**And/Or** is a extremely narrowly scoped. The focus is __ONLY__ on 
currating JSON objects in an asynchronous manner. 

Limiting **And/Or**'s scope leads to a simpler system. Code 
is limited to **And/Or** web service plus the HTML, 
CSS and JavaScript needed for an acceptable UI[^2].

This architecture aligns with small machine hosting and cloud 
hosting. Both keeping recurring costs to a minimum.  **And/Or** could 
be run on a tiny to small EC2 instance or on hardware as small as 
a Rasbpberry Pi.


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
+ Support alternative front ends


## Project Assumptions

+ [dataset](https://github.com/caltechlibrary/dataset) collections are sufficient to hold metadata
+ UI can be implemented using HTML 5 elements, minimal JavaScript and CSS
+ Web UI can easily be built using standard Python packages (e.g. Flask)
+ Small number of curatorial [users](docs/User-Scheme.html)
+ Use existing object scheme (e.g. EPrints XML in Oral Histories)
+ Authentication is external (e.g. Basic Auth, JWT, Shibboleth, OAuth 2)
    + integrated via Python Web Service
+ Other systems handle any additoinal service requirements (e.g. search)


## Limiting features and complexity

Some of the most complicated parts of digital object repositories
are managing customization, managing users, manage roles,
manage permissions, enforcing storage scheme and presenting
public and private views of respository content.  **And/Or**'s 
simplification involves focusing only on the storage and retrieve
of JSON objects. While a dataset collection can easily be used
to store information about users, roles, etc. it doesn't need to 
provide that support directly or even provide a web service for those
aspects of operation. **And/Or**'s assume is that other systems, 
e.g. cli written in Python. 

By focusing on a minimal feature set and leveraging technical 
opportunities that already exist we can radically reduce the lines 
of code written and maintained for a simple object repository. 


## Under the hood

**And/Or**'s JSON document storage engine is [dataset](https://github.com/caltechlibrary/dataset).

> End points map directly to existing dataset operations

dataset operations supported in **And/Or** are "keys", "create", 
"read", "update", "delete", "frame-create", "frame-read", 
"frame-update", "frame-delete", "frame-keys", "frame-refresh" and 
"frame-reframe".  These map to URL paths each supporting a single 
HTTP Method (either GET or POST).

+ `/COLLECTION_NAME/keys/` (GET) all object keys
+ `/COLLECTION_NAME/create/OBJECT_ID` (GET) to creates an Object, an OBJECT_ID must be unique to succeed
+ `/COLLECTION_NAME/read/OBJECT_IDS` (GET) returns on or more objects, if more than one is requested then an array of objects is returned.
+ `/COLLECTION_NAME/update/OBJECT_ID` (POST) to update an object
+ `/COLLECTION_NAME/delete/OBJECT_ID` (POST) to delete an object

**And/Or** is a thin layer on top of existing dataset functionality.
The idea is that as dataset matures and gains the abilities useful in a 
multi-user context they can be exposed in __libdataset__ and then made
accessible in **And/Or**. **And/Or**'s role is to map dataset features 
to an appropriate URL end point.  


### Web UI

Four pages static web pages need to be designed per collection and 
implemented in HTML, CSS and JavaScript for our proof of concept.

1. Login and landing page 
2. Display List records (filterable by object state)
3. An edit page that supports CRUD operations
4. Page to display the logged in user roles


### Limited functionality is intentional

**And/Or** is NOT for public facing content system
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
emptying the trash[^3].


[^1]: NginX and Apache could provide authentication mechanisms such as Basic AUTH, Shibboleth and OAuth 2 and pass them back to a real And/Or implementation.

[^2]: UI, user interface, the normal way a user interacts with a website

[^3]: Empting the trash boils down to traversing all collecting the keys of objects that are in the `._State` == "deleted" and then removing the content from disc.
