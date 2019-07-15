+++
markup = "mmark"
+++

# Oral Histories as Proof of concept

[Oral Histories](http://oralhistories.caltech.edu) is one of our
smaller EPrints installations. It will need migrating soon
because it is running on one of our older blades that lacks https
support. Additionally we should be run on a more recent OS.  
Oral Histories as less than 200 EPrint records. It is a good candidate 
for the proof of concept **AndOr**.

| Week No. | Activity |
|:--------- |:-------- |
| 1 | Write script to remap EPrint XML based document paths |
| 1 - 2 | Write **AndOr** web service and testing |
| 2 - 3 | Write UI in HTML and JavaScript plus more testing |
| 3 | Lunr integration and NginX configuration for Authentication |
| 4 | Harvest Oral Histories for demo deployment |
| 4 | Deploy a demo of a migrated Oral Histories on EC2 |
Table: Proposed schedule assumes full time effort

Migrating other EPrint repositories will involve validating
harvested content and updating the HTML/JavaScript to meet
the repositories' needs. It also should involve an accessment
of if we want any additional microservices to improce workflow
(e.g. integrate our DOI harvesting, external workflows like ETD).
After our first migration the time to migrate should decrease 
with each move instead of remaining constant.

