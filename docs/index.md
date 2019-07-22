+++
markup = "mmark"
+++


# Documentation

## Project

Below are development documents outlining some of the
concepts and rationale around **AndOr**.

+ Concepts and proofs
    + [People and Groups](people-groups.html) 
        + Example of custom metadata objects
    + [Oral Histories](Oral-Histories-as-Proof-of-Concept.html)
        + [Migrating an EPrints Repository](migrating-eprints.html) 
        + Integrating curating Oral Histories content
    + [Users, Workflows and Queues](Workflow-Use-Cases.html)
+ Scheme and scheme user cases
    + [User Scheme](User-Scheme.html)
    + [Workflow Scheme](Workflow-Scheme.html)
    + [Object Scheme](Object-Scheme.html)
+ [Schedule](Schedule.html)
+ [Reference Material](Reference.html)
    + NginX auth options
    + NginX reverse proxy to service

## (Example) Running **AndOr**

The following are actions available from
the command line usage of **AndOr**. While
it takes advantage of the rich library of
packages and tools developed at Caltech Library 
**AndOr** itself is a self contained command
line program.

+ [Setting up AndOr](Setting-Up-AndOr.html)

Basic **AndOr** actions available on the command line

+ [init](init.html) - initialize dataset collections for use by **AndOr** and create example "users.toml", "workflows.toml" and "andor.toml" if needed
+ [check](check.html) - validates "users.toml", "workflows.toml", and "andor.toml"
+ [start](start.html) - start **AndOr** web service using "users.toml", "workflows.toml" and "andor.toml" for configuration

## (Example) What is ...

+ [What's a TOML file?](toml-basics.html)
+ [andor.toml](andor-toml.html)
+ [users.toml](users-toml.html)
+ [workflows.toml](workflows-toml.html)

[^1]: GUI, Graphics user interface, in this case a web based user interface

