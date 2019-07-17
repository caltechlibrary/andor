+++
markup = "mmark"
+++


# Documentation

## Project

Below are development documents outlining some of the
concepts and rationale around **AndOr**.

+ Proof of concept, a GUI[^1] multi-user version of dataset
    + [People and Groups](people-groups.html) - Example of custom metadata objects
    + [Migrating an EPrints Repository](migrating-eprints.html) and [Oral Histories as Proof of Concept](Oral-Histories-as-Proof-of-Concept.html)
    + Proposed [Schedule](Schedule.html)
+ How **AndOr** will work
    + [Setting up AndOr](Setting-Up-AndOr.html)
+ Scheme and scheme user cases
    + [Workflow Use Cases](Workflow-Use-Cases.html)
    + AndOr Scheme
    + [User Scheme](User-Scheme.html)
    + [Workflow Scheme](Workflow-Scheme.html)
    + [Object Scheme](Object-Scheme.html)
+ [Reference Material](Reference.html)
    + NginX auth options
    + NginX reverse proxy to service

## (Example) Running **AndOr**

The following are actions available from
the command line usage of **AndOr**. While
it takes avantage of the rich library of
packages and tools developed at Caltech Library 
**AndOr** itself is a self contained command
line program.

Basic **AndOr** actions available on the command line

+ [init](init.html) - intialize dataset collections for use by **AndOr**
+ [load-workflow](load-workflow.html) - reads a [TOML]() file and adds/updates workflow(s)
+ [list-workflow](list-workflow.html) - Output workflow(s) in TOML
+ [remove-workflow](remove-workflow.html) - Removes a workflow
+ [load-user](load-user.html) - Output user(s) in TOML
+ [remove-user](remove-user.html) - Removes a user
+ [config](config.html) - configure **AndOr** web service
+ [start](start.html) - start **AndOr** web service


## (Example) How to ...

+ [Add a user](add-user.html)
+ [Update a user](update-user.html)
+ [Remove a user](remove-user.html)
+ [Add a workflow](add-user-workflow.html)
+ [Update a workflow](update-a-workflow.html)
+ [Remove a workflow](remove-a-workflow.html)
+ [Listing one or more users](listing-users.html)
+ [Listing one or more workflows](listing-workflows.html)
+ [Create a TOML file?](toml-basics.html)

[^1]: GUI, Graphics user interface, in this case a web based user interface

