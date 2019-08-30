And/Or
=====================================================

> <span class="red">An</span>other <span class="red">d</span>igital / <span class="red">O</span>bject <span class="red">r</span>epository

[![License](https://img.shields.io/badge/License-BSD%203--Clause-blue.svg?style=flat-square)](https://choosealicense.com/licenses/bsd-3-clause)

<!-- [![Latest release](https://img.shields.io/badge/Latest_release-0.0.1-b44e88.svg?style=flat-square)](http://shields.io) -->



**And/Or** is a proof of concept for a simple object repository
based on Caltech Library's [dataset](https://caltechlibrary.github.io/dataset)
tool.  It provides a web-based multi-user verion of dataset suitable for
JSON object curation by a small group of users. It implements a role/object
state based permission scheme enabling support for simple workflows.



Table of contents
-----------------

* [Introduction](#introduction)
* [Installation](install.html)
* [Documentation](docs/)
* [Known issues and limitations](#known-issues-and-limitations)
* [Getting help](#getting-help)
* [Contributing](contributing.html)
* [License](#license)
* [Authors and history](#authors-and-history)
* [Acknowledgments](#authors-and-acknowledgments)


Introduction
------------

**And/Or** is a multi-user web version of [dataset](https://github.com/caltechlibrary/dataset). __dataset__ has proven to be a
useful tool for managing library metadata using a data science approach.
It is easy to import content into it and export from it. It lacks 
multi-user curation support or the convienence of having web browser based
edit forms. The **And/Or** prototype is an exploration
of using dataset as a storage engine and extending it to support multiple
users using a role/object state based permission scheme. It is a proof
of concept for a extremely light weight JSON object repository that supports
simple workflows.

**And/Or** provides a web friendly JSON API mapping __dataset__ functionality
to URLs. It supports serving static HTML, CSS and JavaScript for building
suitable user interfaces that humans night be inclined to use.  The JSON API 
supports creating, reading, update and deleting objects in a collection. The
API supports key list retrieval and limitted filtering.  Finally And/Or
provides BasicAUTH user authentication support for mapping users 
to role/object state permission scheme.  With these primitive facilities 
you can build a simple object repository systems using the HTML, CSS and 
JavaScript to create a web client interacting with the And/Or JSON API.


Installation
------------

See [INSTALL.md](install.html). This software is experimental
and pre-compiled binaries are NOT provided.  This software is written in 
[Go](https://golang.org) programming language and needs a Go compiler
to be compiled. 


```bash
    go get -u github.com/caltechlibrary/andor/...
```

A "Makefile" has been provided for your convience if you wish to
manually compile the program.

```bash
    git clone https://github.com/caltechlibrary/andor 
    cd andor
    go build -o bin/andor cmd/andor/andor.go
```

After compiling **And/Or** you can run a localhost
demo with the following command.

```
    ./bin/andor start demo-andor.toml
```

Point your web broser at http://localhost:8246 to try the demo.


Known issues and limitations
----------------------------

This is a proof-of-concept project. It SHOULD NOT be used
in a production setting.  It is ONLY suitable for demonstration
an approach to building light weight object repositories.

Getting help
------------

You can contact us via GitHub [issue tracker](https://github.com/caltechlibrary/andor/issues).

Contributing
------------

See [CONTRIBUTING.md](contributing.html)


License
-------

Software produced by the Caltech Library is Copyright (C) 2019, Caltech.  This software is freely distributed under a BSD/MIT type license.  Please see the [LICENSE](LICENSE) file for more information.


Authors and history
---------------------------

[R. S. Doiel](https://rsdoiel.github.io) is the culprit responsible for this proof of concept


Acknowledgments
---------------

This work was funded by the California Institute of Technology Library.

(If this work was also supported by other organizations, acknowledge them here.  In addition, if your work relies on software libraries, or was inspired by looking at other work, it is appropriate to acknowledge this intellectual debt too.)

<div align="center">
  <br>
  <a href="https://www.caltech.edu">
    <img width="100" height="100" src="assets/caltech-round.svg">
  </a>
</div>

