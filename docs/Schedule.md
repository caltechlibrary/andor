+++
markup = "mmark"
+++

# Oral Histories as Proof of concept

[Oral Histories](http://oralhistories.caltech.edu) is one of our
smaller EPrints installations. It will need migrating soon
because it is running on one of our older blades, lacks https
support and should be run on a more recent OS.  It has less
than 200 EPrint records. It is a good candidate for the proof of
concept **AndOr**.

| Week No. | Activity |
|:--------- |:-------- |
| 1 | Update dataset to have automatic attachment versioning |
| 1 | Add JSON document diff on update, versioning for metadata |
| 1 | Harvest Oral Histories to use for test data |
| 1 - 2 | Write **AndOr** web service and testing |
| 2 - 3 | Write UI in HTML and JavaScript plus more testing |
| 3 | Harvest Oral Histories for demo deployment |
| 3 | Lunr integration and NginX configuration for Authentication |
| 4 | Deploy a demo of a migrated Oral Histories on EC2 |
Table: Proposed schedule assumes full time effort

Migrating other EPrint repositories will involve validating
harvested content and updating the HTML/JavaScript to meet
the repositories' needs. This means the migration process
after the first one will likely decrease with each move instead
of remaining constant.

