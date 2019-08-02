+++
markup = "mmark"
+++

# Proposed Development schedule

| Week No. | Proposed start date | Activity |
|:--------- |:-------- | :-------- |
| 1 | 2019-07-22 | Model users and roles |
| 2 | 2019-07-29 | Write **AndOr** web service and testing |
| 3 | 2019-08-05 | Implement Basic UI in HTML and JavaScript plus more testing |
| 4 | 2019-08-12 | Implement demo UIs for demo collections |
| 5 | 2019-08-19 | Ready to demo **AndOr** |
Table: Proposed schedule assumes full time effort, dates subject to change

## People and Groups as Proof of concept

We are out growing our GSheet based People and Groups list.
**AndOr** is to provide an alternative curating interface.
The demo will show roles for managing two collections---
People and Groups. It should provide the following basic
capabilities

1. Create, Update, Delete People and Group records (rows in the GSheet today)
2. List People and Groups
2. Support users authenticating via a simple system via NginX or Apache
3. Have roles for "suggested", "published" and "manage" objects in each collection

## Oral Histories as Proof of concept

[Oral Histories](http://oralhistories.caltech.edu) is one of our
smaller EPrints installations. It will need migrating soon
because it is running on one of our older blades that lacks https
support. It is a good candidate for the proof of concept **AndOr**.
The demo should demonstrate the following

1. How to harvest the current EPrints based deployment
2. Setup **AndOr** to host the harvested content
3. Support user authenticating via a simple system
4. Implement roles currently available in EPrints

Migrating other EPrints repositories will involve validating
harvested content and updating the HTML/JavaScript to meet
the repositories' needs. It also should involve an assessment
of if we want any additional micro services to improve role
(e.g. integrate our DOI harvesting, external roles like ETD).
After our first migration the time to migrate should decrease 
with each move instead of remaining constant.

