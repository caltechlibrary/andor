+++
markup = "mmark"
+++

# Oral Histories As Proof of Concept

## Problem

[Oral Histories](http://oralhistories.library.caltech.edu) is an
EPrints repository with a little less than 200 EPrint objects in
it. It currently runs in a multi-host EPrints on some of our oldest
hardware[^1].  Oral Histories also needs to run with HTTPS support.
To achieve that in the current deployment would require restructuring
the SSL certs. One of the deployments on that machine is Thesis
redoing the certs would mean downtime and I'd like to avoid that for
Thesis. In addition to needing SSL support, we need to swap out the
operating system[^2] and upgrade EPrints[^3].

I've started to migrate Oral Histories three times. Each time
something important has come up and stopped my progress. Because of
operating system and dependent software version drift returning to the
project is essentially starting over.

## Challenges

Lacking HTTPS support means our this means our Oral History Project
will increasingly become invisible to the rest of the web as the
search engines (e.g. Google) starts to down rank non-HTTPS
sites. Users will get disturbing warnings when they visit our site
(assuming they can find it) as Firefox, Chrome, et el. push people to
only access SSL protected sites.

Setting up a working vanilla EPrints is non-trivial. Migrating
our customizations is complex and error prone[^4].  Each operating
system, software upgrade and EPrints upgrades risks breaking EPrints
particularly when we've customized the code.

Our initial goal has been to run Oral Histories on AWS EC2 instance
as if it was hosted locally. This is only a first step. It leaves
running EPrints the most expensive way you can in the cloud.
The next step would be to migrate EPrint's disk0[^5] tree to S3.
This is a non-trivial implementation change too. The final step
towards a cloud friendly EPrints deployment would be to migrate
from MySQL/MariaDB to AWS's RDMS[^6]. Each change is essentially its
own project. Testing means it each step can wind up as a multi-week
project.

We need to migrate to a current Ubuntu LTS release from a very old
RedHat/CentOS release. Our version of RedHat (Cent OS) isn't getting the
same level of security patches we need. The Internet is a hostile
environment so you really want to keep you operating system and any
dependent software (e.g. Perl, OpenSSL, MySQL) current. Each upgrade
runs the risk of breaking our EPrints deployment.

There are no adequate non-EPrints replacements for EPrints today.
We've been betting a likely candidates would emerge but it might be
prudent to hedge that bet as they haven't appeared since I was
hired by Caltech Library.

I believe migrating any EPrints instance is at minimum a 4 to 8 weeks of
full time continuous work for me. I think this estimate will hold for
the first repository through the last repository migrated. Ideally each
successive migration would be shorter as we automate the process.


## Opportunities

The Oral Histories repository is small, 200 objects. It has a small
number of users who curate the content (e.g. max number is likely
size of the Archives staff plus two). A few objects are added/updated
per week depending on funding and interviewable individuals.
The metadata and media files can be curated separately from the
public facing website (like we do now with the feeds or the
archives own website).

The curator needs to have a list of new objects needing work and also
be able to access published objects and assets should the need arise.
Ideally you wouldn't need to cut and paste the metadata between the
curation tool and ArchivesSpace. Integrating generally with ArchivesSpace
would be desirable.

It would be nice to have a newer designed website. It would be nice to
have a good full text search available for content on the site
(both in curation and for the public).

It is a straight forward project to take what we've already implemented
for feeds.library.caltech.edu and adapt to the needs of a much better
Oral Histories site for public viewing. This leaves us with allot of
flexibility in picking the minimum feature set needed to curate the
objects and media files.


## Curation workflow based on existing EPrints implementation

We have most of the software in parts for creating a light weight
EPrints like system. What is lacking is a web service that functions
like EPrint's REST API. We can use EPrint's existing Schema even if
we skip the RDMS backend.  We know how to harvest all the data in
EPrints and already have tooling in place to due so.

If we had a light weight EPrints like system migration would take
a few days to a week or so depending on the EPrints repository,
number of objects the customization to the curation form. This is
considerably less than migrating EPrints itself.  The primary
task would be the customization of the object edit and input forms
to support each specific implementation. Additionally this light
weight EPrints could be an __interim solution__ while we wait for
systems like
[Zenodo](https://zenodo.org/ "An Object Repository System used in conjunction with CERN") or
[Achepelligo](http://archipelago.nyc/ "Metro's demonstration of their Open Source Digital Object Repository")
to mature.


## What would a light weight EPrints look like?

An ability for staff to login and have access to a workflow
queues of objects. This includes the ability to create an object
and add any related media files (e.g. audio, video, pdfs).

Creating an EPrint object follows the following pattern, pick
a type, add general metadata (e.g. description/abstract, title).
Upload media files. Add any more internal data and then "deposit".
Once deposited the object is available for processing in EPrints.

EPrints uses a queue system for implementing a workflow. A user's
abilities are matched to the workflows they have access to. In effect
their set of workflows defines their permissions. This relationship
can be modeled as a User Object and Workflow Objects.

### Decomposing the EPrints Admin functionality for Oral Histories

If you are the "admin" user Oral Histories' EPrints curatorial
interface breaks down as follows--

Manage Deposits
: A queue of objects that have been deposited

Manage Records
: A page with links to manage EPrints and the EPrints record

Manage Shelves
: Doesn't work, this EPrints feature isn't used in Oral Histories

Profile
: This isn't used much and could be skipped, it shows how the "admin" user is setup

Saved Searches
: Not used

Review
: This is a queue of EPrint objects needing review

Admin
: This is the admin page for managing EPrints from the web browser

The "admin" user has the most capability. If you can capture the needed
functionality of the "admin" user you can capture all the functionality
available to all staff who use EPrints for Oral Histories.  Most of
the functionality above could be trimmed and still provide a workable
system. What you really need is "Manage Deposits", "Review" and the
"EPrints" record link for all objects from "Manage Records" page view.

This can be boiled down to the following functions

1. (Create) A screen to create an object, attach it's media files and put it into the "deposit" queue
2. (Read/Update) Move to "Review" queue from "deposit", update as needed
3. (Update) Move to "Published" queue so the object and media files are available to the public.
4. List objects by queue (sortable by date deposited, item type, last modified, Item Status, Title)
5. List All objects (with previous sorting options)
6. Search for objects
7. Support "deleting" objects, EPrints never deletes, it just hides the object in a "delete" queue.

The above suggests we need to know who are user is and what workflows
they are allowed to access. A workflow identifies a queue of objects, the
permissions to act on those objects (e.g. create, read, update) and the
next queue they will be passed to.

The workflow is "deposit", "review", "publish" or "delete". The user info
is a display name, a user id, what queue they create objects in
(e.g. "deposit") and what workflows they have access to (what queues they
can review and what queues they can place objects into).

The same form used to create an object can serve as the form to update an
object (just would have populated values). That form could be MUCH simpler
than EPrints while still capturing the same information. The form itself
could be implemented as a combination of static HTML and JavaScript where
the JavaScript talks to a web API to fill in the details.

An Empty "Oral Histories" deployment could be a few web pages that talk to
a web service in the same way that the Builder Widget can talk to feeds
and resent results.

### Harvesting is already in the works

Harvesting Oral Histories' metadata for feeds is already in development.
It use the EPrinttools suite to store the harvested content in a
dataset collection. The metadata harvested is stored in a
[dataset](https://github.com/caltechlibrary/dataset) collection.  
The harvester could easily include harvesting the underling
media files as well as the metadata.  **dataset** could replace the
storage framework currently implemented in EPrints.

### A Repository Engine already is implemented

**dataset** plus Lunr[^7] already provides the functionality
of a EPrints or Fedora based repository system. It only lacks a web
service interface. This is straight forward to implement in Go especially
if you limit the service to a few end points (e.g. list objects, list an
object's details, create/update an object). EPrints show us the
way in its REST API already. Implementing a web service like this
based on dataset is a straight forward process because we know how it
should work and have implemented things like this in the past[^8].

### EPrints shows us the way, we can continue to work the EPrints way

We don't need to develop scheme, we can use what EPrints XML provides.
EPrints shows the API we need because we already use EPrints REST API
to drive feeds.library.caltech.edu. EPrints also shows us the model
for users and workflows. Users are affiliated with workflows. Workflows
are a queue of objects (objects hold a single value for workflow state).
Works define what actions can be performed on an object (e.g. "published"
is read only, a curator can see an "edit" link which takes them to a
web form).

### We can avoid the things that repository systems (EPrints) complicated

+ EPrints is complicated because it allows complete customization. Each
customization though is technical debt because it complicates upgrades.
+ EPrints is complicated because it much define and manage workflows. In
a large library where administration is distributed across many staff
this make sense. We're a small library (a tiny archive) which could
make dues with a simple configuration file.
+ EPrints is complicated because it needs to support hosting MANY collections
independently behind the same web service. This is driven by history, this is
how we originally built content management systems because campuses had a
single web server. Today spinning up a web server is trivial, unless it's an EPrints
web server.

We can avoid the complexity of EPrints for most of our EPrints deployments. It is
likely we can even avoid complexity when supporting are larger repository if we
separate the curation/submission process from publicly viewable websites and switch
from SQL searches to full text searches based on something like Lunr, Solr or
Elastic Search.



[^1]: All the EPrints instances on the hardware are due to be migrated.

[^2]: Our EPrints deployments currently run on an old version of RedHat (Cent OS). Ubuntu 18.04 LTS is our current standard distribution.  Ubuntu LTS releases get regular security updates keep pace with the demands of a hostile internet. Our version of RedHat does not.

[^3]: Our EPrints deployment on coda2.cls.caltech.edu is version EPrints 3.3, EPrints us up to version 3.4. The newest version supports Ubuntu 18.04 LTS.

[^4]: EPrints is a highly customizable system. That customization can happen at any of three levels, you change the vanilla EPrints, you change an EPrints instance, you overwrite module behavior with a customized replacement.

[^5]: The disk0 folder is where both the EPrint XML is stored as well as any media (e.g. pdf, image file, audio file).

[^6]: AWS has a hosted RDBMS (relational database manage system) that appears to your application as if it is MySQL (what EPrints uses). There are edge cases particularly when your record data has problematic character encodings (i.e. not UTF-8) that make this swap a project in itself. The actual migration would be existing MySQL/MariaDB to a clean copy with all encodings fixed MariaDB then dump and restore into AWS's RDBMS. Starting from a clean slate with RDBMS on the other hand is most trivial (basic sort out the connection string and permissions).

[^7]: Lunr is a full text index specification implemented in Python and in JavaScript. For small indexes they can be completely run browser side, larger ones can be implemented server side with a simple Python web service. Stephen showed us this approach using feeds data.

[^8]: An earily version of dataset actually had a web service, we depreciated it because at the time we had no need of it.
