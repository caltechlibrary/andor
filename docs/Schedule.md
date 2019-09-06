+++
markup = "mmark"
+++

# Adjusted Development schedule

Due to actual time off the schedule was spread across more weeks.

| Week No. | Proposed start date | Activity |
|:--------- |:-------- | :-------- |
| 1 | 2019-07-22 | Model users and roles |
| 2 | 2019-07-29 | Write **AndOr** web service and testing |
| 3 | 2019-08-12 | Implement Basic UI in HTML and JavaScript plus more testing |
| 3 | 2019-08-05 | Haitus | Time off |
| 4 | 2019-08-26 | Implement demo UIs for demo collections |
| 4 | 2019-08-28 | Internal demo and discussion or **AndOr** | 
| 5 | 2019-09-02 | Revise, debug demo for DLD and others |
Table: Original schedule assumed was five weeks full time effort, actual was 5 weeks were stretch over two months


## People as Proof of concept

We are out growing our GSheet based People list.
**AndOr** is a proposal to explore an alternative curation 
tool.  The demo implements listing people records, creating,
reading, updating and "deleting" individual records using
a user, role, object state module to implement workflows.

Capabilities

1. Create, Update, Delete People records (rows in the GSheet today)
2. List People by object state
2. Support users authenticating via http BasicAUTH
3. Roles support "deposit", "review", "published", "embargoed" and "deleted" objects 

Currently we export the Google Spreadsheet for CaltechPEOPLE
and create a people dataset collection. This proof of concept shows
what it might be like to skip the spreadsheet and manage the
JSON objects directly via **And/Or**.

