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
It will provide a semi-RESTful JSON API and use 
HTML, CSS and JavaScript to provide a GUI interface for curating.
A minimum running system would consist of only two or three
pieces of software. The minimum would be a web server[^1] 
plus the **And/Or** service supporting multi-user interaction with 
dataset collections.  Depending on size you could create
a micro service providing search via Python and Lunr[^2] or generate
Lunrjs indexes to run browser side.  This arrangement has the 
advantage of limiting the code to be written to the **AndOr** 
web service plus the HTML, CSS and JavaScript needed to create 
an acceptable UI[^3].

This particular architecture aligns with small machine hosting
and cloud hosting keeping recurring costs to a minimum. 
In the cloud it should work on a small to medium EC2 instance.
A more elaborate version be using Cloud Front, S3, and running 
the API service in an AWS container or via the AWS Lambda. 
In house hosting could as light weight as Raspberry Pi 4 or
as elaborate as a server with attached NAS[^4].


## Goals

+ Provide a curatorial platform for metadata outside our existing repositories
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
+ UI can be implemented using HTML 5 elements, minimal JavaScript and CSS
+ Small number of curatorial users, larger number of readers
+ Configurable workflows are a requirement
    + workflows describe capabilities (permissions)
    + workflows describe a work queue (object state)
+ Use existing object scheme (e.g. EPrints XML in Oral Histories)
+ Authentication is external (e.g. Shibboleth, OAuth 2)
+ Search and query are handle independent of API
    + e.g. Lunr either browser or server side


## Limiting features and complexity

Some of the most complicated parts of digital object repositories
are managing customization, managing users, manage workflows,
manage permissions and enforcing storage scheme.  **AndOr**'s 
simplification involves either avoiding the requirements or relocating
them to an appropriate external system.  

Examples--

+ Authentication is handle externally. That means we don't need to worry about managing passwords, the volatile and sensitive data is outside of our system
+ **AndOr** itself is a simple web API that accepts URL requests 
and hands back JSON. The shape of the JSON is determined at time of
migrating into **AndOr**. There is no customization.  If you want to change your data shapes you write a script to do that or change your input form.
+ If you need additional end points beyond what **AndOr** provides (e.g. a search engine service) you supply those as micro services behind the same web server

The web browser itself creates the illusion of a unified software system
or single process. A single application is not required to support desire
functionality. Customization then can be deferred to other micro services
or even external systems (e.g. looking up something at datacite.org or
orcid.org).

Some features are unavoidable. The repository problem requires managing
users and workflows. It doesn't require users and workflows
be manage through the web. Setting up users and workflows can be 
managed through simpler command line tools and configuration files.
The is reasonable in large part because you've off loaded 
identify management already. 

By focusing on a minimal feature set and leveraging technical
opportunities that already exist we can radically
reduce the lines of code written and maintained. 

### Two end points for our API

Two web API end points would be required 

1. `/COLLECTION_NAME/objects/` to list objects keys
2.  `/COLLECTION_NAME/objects/OBJECT_ID` to provide an Object's details

Lists can be filtered by workflow's queue name
All other end points are static resources (e.g. HTML files, 
CSS, JavaScript and Lunrjs indexes, a public faces website).  
We can reducing our requirements to two end points because 
we've discovered that is all we needed based on our work
with [EPrints](https://www.eprints.org "A repository system developed at University of South Hampton").
Everything else can be synthesized from a simple key list and object access.


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
This also keeps **AndOr** simple with fewer requirements.


### user plus workflow, a simple model

An authenticated user exposes their user id to **AndOr**'s
web service. A user's id maps to membership in workflows. 
The workflow defines access to queues to list objects.

Unauthenticated users are treated as the "anonymous" user and
are restricted by workflows available for the "anonymous" user. 
This is how you would control having a dark versus publicly 
visible repository.

Complicated use cases like electronic thesis deposits
or faculty self deposit of articles could be implemented as 
separate specialized micro services that interacts with
external authentication then implements a specialized 
workflow interacting with the API via a service account
an appropriate workflow associated with it.


### Under the hood

**AndOr** is built on [dataset](https://caltechlibrary.github.io/dataset).
Objects may include attached documents which can be versioned 
automatically. If metadata versioning becomes required dataset 
can be extended to store diffs as well as the JSON documents.

Like EPrints **AndOr** does not directly support deleting objects.
Instead it can create the illusion of deleting objects by putting
objects into a "deleted" queue which you can exclude from your
workflows.


## Additional ideas

+ Use cases
    + [Users, Workflows and Queues](docs/Workflow-Use-Cases.html)
+ Concept proofs
    + [People and Groups](docs/people-groups.html)
    + [Migrating an EPrints Repository](docs/migrating-eprints.html) 
    + [Oral Histories](Oral-Histories-as-Proof-of-Concept.html)
+ Scheme walk through
    + [User Scheme](docs/User-Scheme.html)
    + [Workflow Scheme](docs/Workflow-Scheme.html)
    + [Queue Scheme](docs/Queue-Scheme.html)
    + [Object Scheme](docs/Object-Scheme.html)




[^1]: NginX and Apache provide authentication mechanisms such as Basic AUTH, Shibboleth and OAuth 2.

[^2]: [Lunr](https://lunrjs.com) is a browser friendly indexing and search library that can now be supported server side too via Python.

[^3]: UI, user interface, the normal way a user interacts with a website

[^4]: NAS, network attached storage similar to what are now common in research labs
